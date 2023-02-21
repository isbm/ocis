package config

import (
	"context"

	"github.com/owncloud/ocis/v2/ocis-pkg/shared"
	settingsmsg "github.com/owncloud/ocis/v2/protogen/gen/ocis/messages/settings/v0"
)

// Config combines all available configuration parts.
type Config struct {
	Commons *shared.Commons `yaml:"-"` // don't use this directly as configuration for a service

	Service Service `yaml:"-"`

	Tracing *Tracing `yaml:"tracing"`
	Log     *Log     `yaml:"log"`
	Debug   Debug    `yaml:"debug"`

	HTTP HTTP       `yaml:"http"`
	GRPC GRPCConfig `yaml:"grpc"`

	GRPCClientTLS *shared.GRPCClientTLS `yaml:"grpc_client_tls"`

	StoreType   string                `yaml:"store_type" env:"SETTINGS_STORE_TYPE" desc:"Store type configures the persistency driver. Supported values are \"metadata\" and \"filesystem\"."`
	DataPath    string                `yaml:"data_path" env:"SETTINGS_DATA_PATH" desc:"The directory where the filesystem storage will store ocis settings. If not definied, the root directory derives from $OCIS_BASE_DATA_PATH:/settings."`
	Metadata    Metadata              `yaml:"metadata_config"`
	BundlesPath string                `yaml:"bundles_path" env:"SETTINGS_BUNDLES_PATH" desc:"The path to a JSON file with a list of bundles. If not definied, the default bundles will be loaded."`
	Bundles     []*settingsmsg.Bundle `yaml:"-"`

	AdminUserID string `yaml:"admin_user_id" env:"OCIS_ADMIN_USER_ID;SETTINGS_ADMIN_USER_ID" desc:"ID of the user that should receive admin privileges. Consider that the UUID can be encoded in some LDAP deployment configurations like in .ldif files. These need to be decoded beforehand."`

	TokenManager *TokenManager `yaml:"token_manager"`

	SetupDefaultAssignments bool `yaml:"set_default_assignments" env:"SETTINGS_SETUP_DEFAULT_ASSIGNMENTS;ACCOUNTS_DEMO_USERS_AND_GROUPS" desc:"The default role assignments the demo users should be setup."`

	Context context.Context `yaml:"-"`
}

// Metadata configures the metadata store to use
type Metadata struct {
	GatewayAddress string `yaml:"gateway_addr" env:"STORAGE_GATEWAY_GRPC_ADDR" desc:"GRPC address of the STORAGE-SYSTEM service."`
	StorageAddress string `yaml:"storage_addr" env:"STORAGE_GRPC_ADDR" desc:"GRPC address of the STORAGE-SYSTEM service."`

	SystemUserID     string `yaml:"system_user_id" env:"OCIS_SYSTEM_USER_ID;SETTINGS_SYSTEM_USER_ID" desc:"ID of the oCIS STORAGE-SYSTEM system user. Admins need to set the ID for the STORAGE-SYSTEM system user in this config option which is then used to reference the user. Any reasonable long string is possible, preferably this would be an UUIDv4 format."`
	SystemUserIDP    string `yaml:"system_user_idp" env:"OCIS_SYSTEM_USER_IDP;SETTINGS_SYSTEM_USER_IDP" desc:"IDP of the oCIS STORAGE-SYSTEM system user."`
	SystemUserAPIKey string `yaml:"system_user_api_key" env:"OCIS_SYSTEM_USER_API_KEY" desc:"API key for the STORAGE-SYSTEM system user."`
}
