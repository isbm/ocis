// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	gocloak "github.com/Nerzal/gocloak/v13"
	keycloak "github.com/owncloud/ocis/v2/ocis-pkg/keycloak"

	libregraph "github.com/owncloud/libre-graph-api-go"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, realm, user, userActions
func (_m *Client) CreateUser(ctx context.Context, realm string, user *libregraph.User, userActions []keycloak.UserAction) (string, error) {
	ret := _m.Called(ctx, realm, user, userActions)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *libregraph.User, []keycloak.UserAction) (string, error)); ok {
		return rf(ctx, realm, user, userActions)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *libregraph.User, []keycloak.UserAction) string); ok {
		r0 = rf(ctx, realm, user, userActions)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *libregraph.User, []keycloak.UserAction) error); ok {
		r1 = rf(ctx, realm, user, userActions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPIIReport provides a mock function with given fields: ctx, realm, username
func (_m *Client) GetPIIReport(ctx context.Context, realm string, username string) (*keycloak.PIIReport, error) {
	ret := _m.Called(ctx, realm, username)

	var r0 *keycloak.PIIReport
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*keycloak.PIIReport, error)); ok {
		return rf(ctx, realm, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *keycloak.PIIReport); ok {
		r0 = rf(ctx, realm, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keycloak.PIIReport)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, realm, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByParams provides a mock function with given fields: ctx, realm, params
func (_m *Client) GetUserByParams(ctx context.Context, realm string, params gocloak.GetUsersParams) (*libregraph.User, error) {
	ret := _m.Called(ctx, realm, params)

	var r0 *libregraph.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, gocloak.GetUsersParams) (*libregraph.User, error)); ok {
		return rf(ctx, realm, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, gocloak.GetUsersParams) *libregraph.User); ok {
		r0 = rf(ctx, realm, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*libregraph.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, gocloak.GetUsersParams) error); ok {
		r1 = rf(ctx, realm, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsername provides a mock function with given fields: ctx, realm, username
func (_m *Client) GetUserByUsername(ctx context.Context, realm string, username string) (*libregraph.User, error) {
	ret := _m.Called(ctx, realm, username)

	var r0 *libregraph.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*libregraph.User, error)); ok {
		return rf(ctx, realm, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *libregraph.User); ok {
		r0 = rf(ctx, realm, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*libregraph.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, realm, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendActionsMail provides a mock function with given fields: ctx, realm, userID, userActions
func (_m *Client) SendActionsMail(ctx context.Context, realm string, userID string, userActions []keycloak.UserAction) error {
	ret := _m.Called(ctx, realm, userID, userActions)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []keycloak.UserAction) error); ok {
		r0 = rf(ctx, realm, userID, userActions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
