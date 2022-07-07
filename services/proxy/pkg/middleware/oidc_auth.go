package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	gOidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/owncloud/ocis/v2/ocis-pkg/log"
	"github.com/owncloud/ocis/v2/ocis-pkg/oidc"
	"github.com/owncloud/ocis/v2/ocis-pkg/sync"
	"golang.org/x/oauth2"
)

// OIDCProvider used to mock the oidc provider during tests
type OIDCProvider interface {
	UserInfo(ctx context.Context, ts oauth2.TokenSource) (*gOidc.UserInfo, error)
}

// OIDCAuth provides a middleware to check access secured by a static token.
func OIDCAuth(optionSetters ...Option) func(next http.Handler) http.Handler {
	options := newOptions(optionSetters...)
	tokenCache := sync.NewCache(options.UserinfoCacheSize)

	h := oidcAuth{
		logger:        options.Logger,
		providerFunc:  options.OIDCProviderFunc,
		httpClient:    options.HTTPClient,
		oidcIss:       options.OIDCIss,
		tokenCache:    &tokenCache,
		tokenCacheTTL: options.UserinfoCacheTTL,
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// there is no bearer token on the request,
			if !h.shouldServe(req) {
				// oidc supported but token not present, add header and handover to the next middleware.
				userAgentAuthenticateLockIn(w, req, options.CredentialsByUserAgent, "bearer")
				next.ServeHTTP(w, req)
				return
			}

			if h.getProvider() == nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			token := strings.TrimPrefix(req.Header.Get("Authorization"), "Bearer ")

			claims, status := h.getClaims(token, req)
			if status != 0 {
				w.WriteHeader(status)
				return
			}

			// inject claims to the request context for the account_resolver middleware.
			next.ServeHTTP(w, req.WithContext(oidc.NewContext(req.Context(), claims)))
		})
	}
}

type oidcAuth struct {
	logger        log.Logger
	provider      OIDCProvider
	providerFunc  func() (OIDCProvider, error)
	httpClient    *http.Client
	oidcIss       string
	tokenCache    *sync.Cache
	tokenCacheTTL time.Duration
}

func (m oidcAuth) getClaims(token string, req *http.Request) (claims map[string]interface{}, status int) {
	hit := m.tokenCache.Load(token)
	if hit == nil {
		oauth2Token := &oauth2.Token{
			AccessToken: token,
		}

		userInfo, err := m.getProvider().UserInfo(
			context.WithValue(req.Context(), oauth2.HTTPClient, m.httpClient),
			oauth2.StaticTokenSource(oauth2Token),
		)
		if err != nil {
			m.logger.Error().Err(err).Msg("Failed to get userinfo")
			status = http.StatusUnauthorized
			return
		}

		if err := userInfo.Claims(&claims); err != nil {
			m.logger.Error().Err(err).Interface("userinfo", userInfo).Msg("failed to unmarshal userinfo claims")
			status = http.StatusInternalServerError
			return
		}

		expiration := m.extractExpiration(token)
		m.tokenCache.Store(token, claims, expiration)

		m.logger.Debug().Interface("claims", claims).Interface("userInfo", userInfo).Time("expiration", expiration.UTC()).Msg("unmarshalled and cached userinfo")
		return
	}

	var ok bool
	if claims, ok = hit.V.(map[string]interface{}); !ok {
		status = http.StatusInternalServerError
		return
	}
	m.logger.Debug().Interface("claims", claims).Msg("cache hit for userinfo")
	return
}

// extractExpiration currently just returns a hardcoded default for now. It was
// supposed to parse and extract the expiration time from the provided
// access_token.
// As the access_token is defined as an opaque string. Validating and parsing it
// can be tricky:
// 1. Try to treat it as a JWT:
//    - Verifying the validity of the token requires downloading the propoer public
//      key from the IDP (uri in "jwks_uri" in ".well-known/openid-configuration"
// 2. Verify and extract it via the introspection endpoint of the IDP (RFC7662) for
//    IDPs that provide that feature
// 3. Other IDP implementation specific methods.
// 4. Fallback to default value
func (m oidcAuth) extractExpiration(token string) time.Time {
	defaultExpiration := time.Now().Add(m.tokenCacheTTL)
	return defaultExpiration
}

func (m oidcAuth) shouldServe(req *http.Request) bool {
	header := req.Header.Get("Authorization")

	if m.oidcIss == "" {
		return false
	}

	// todo: looks dirty, check later
	// TODO: make a PR to coreos/go-oidc for exposing userinfo endpoint on provider, see https://github.com/coreos/go-oidc/issues/248
	for _, ignoringPath := range []string{"/konnect/v1/userinfo", "/status.php"} {
		if req.URL.Path == ignoringPath {
			return false
		}
	}

	return strings.HasPrefix(header, "Bearer ")
}

func (m *oidcAuth) getProvider() OIDCProvider {
	if m.provider == nil {
		// Lazily initialize a provider

		// provider needs to be cached as when it is created
		// it will fetch the keys from the issuer using the .well-known
		// endpoint
		provider, err := m.providerFunc()
		if err != nil {
			m.logger.Error().Err(err).Msg("could not initialize oidcAuth provider")
			return nil
		}

		m.provider = provider
	}
	return m.provider
}
