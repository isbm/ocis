// Code generated by mockery v2.40.2. DO NOT EDIT.

package mocks

import (
	context "context"

	godata "github.com/CiscoM31/godata"

	libregraph "github.com/owncloud/libre-graph-api-go"

	mock "github.com/stretchr/testify/mock"

	url "net/url"
)

// Backend is an autogenerated mock type for the Backend type
type Backend struct {
	mock.Mock
}

type Backend_Expecter struct {
	mock *mock.Mock
}

func (_m *Backend) EXPECT() *Backend_Expecter {
	return &Backend_Expecter{mock: &_m.Mock}
}

// AddMembersToGroup provides a mock function with given fields: ctx, groupID, memberID
func (_m *Backend) AddMembersToGroup(ctx context.Context, groupID string, memberID []string) error {
	ret := _m.Called(ctx, groupID, memberID)

	if len(ret) == 0 {
		panic("no return value specified for AddMembersToGroup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) error); ok {
		r0 = rf(ctx, groupID, memberID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_AddMembersToGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddMembersToGroup'
type Backend_AddMembersToGroup_Call struct {
	*mock.Call
}

// AddMembersToGroup is a helper method to define mock.On call
//   - ctx context.Context
//   - groupID string
//   - memberID []string
func (_e *Backend_Expecter) AddMembersToGroup(ctx interface{}, groupID interface{}, memberID interface{}) *Backend_AddMembersToGroup_Call {
	return &Backend_AddMembersToGroup_Call{Call: _e.mock.On("AddMembersToGroup", ctx, groupID, memberID)}
}

func (_c *Backend_AddMembersToGroup_Call) Run(run func(ctx context.Context, groupID string, memberID []string)) *Backend_AddMembersToGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]string))
	})
	return _c
}

func (_c *Backend_AddMembersToGroup_Call) Return(_a0 error) *Backend_AddMembersToGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_AddMembersToGroup_Call) RunAndReturn(run func(context.Context, string, []string) error) *Backend_AddMembersToGroup_Call {
	_c.Call.Return(run)
	return _c
}

// CreateGroup provides a mock function with given fields: ctx, group
func (_m *Backend) CreateGroup(ctx context.Context, group libregraph.Group) (*libregraph.Group, error) {
	ret := _m.Called(ctx, group)

	if len(ret) == 0 {
		panic("no return value specified for CreateGroup")
	}

	var r0 *libregraph.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, libregraph.Group) (*libregraph.Group, error)); ok {
		return rf(ctx, group)
	}
	if rf, ok := ret.Get(0).(func(context.Context, libregraph.Group) *libregraph.Group); ok {
		r0 = rf(ctx, group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*libregraph.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, libregraph.Group) error); ok {
		r1 = rf(ctx, group)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Backend_CreateGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateGroup'
type Backend_CreateGroup_Call struct {
	*mock.Call
}

// CreateGroup is a helper method to define mock.On call
//   - ctx context.Context
//   - group libregraph.Group
func (_e *Backend_Expecter) CreateGroup(ctx interface{}, group interface{}) *Backend_CreateGroup_Call {
	return &Backend_CreateGroup_Call{Call: _e.mock.On("CreateGroup", ctx, group)}
}

func (_c *Backend_CreateGroup_Call) Run(run func(ctx context.Context, group libregraph.Group)) *Backend_CreateGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(libregraph.Group))
	})
	return _c
}

func (_c *Backend_CreateGroup_Call) Return(_a0 *libregraph.Group, _a1 error) *Backend_CreateGroup_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Backend_CreateGroup_Call) RunAndReturn(run func(context.Context, libregraph.Group) (*libregraph.Group, error)) *Backend_CreateGroup_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *Backend) CreateUser(ctx context.Context, user libregraph.User) (*libregraph.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *libregraph.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, libregraph.User) (*libregraph.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, libregraph.User) *libregraph.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*libregraph.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, libregraph.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Backend_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type Backend_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user libregraph.User
func (_e *Backend_Expecter) CreateUser(ctx interface{}, user interface{}) *Backend_CreateUser_Call {
	return &Backend_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, user)}
}

func (_c *Backend_CreateUser_Call) Run(run func(ctx context.Context, user libregraph.User)) *Backend_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(libregraph.User))
	})
	return _c
}

func (_c *Backend_CreateUser_Call) Return(_a0 *libregraph.User, _a1 error) *Backend_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Backend_CreateUser_Call) RunAndReturn(run func(context.Context, libregraph.User) (*libregraph.User, error)) *Backend_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteGroup provides a mock function with given fields: ctx, id
func (_m *Backend) DeleteGroup(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteGroup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_DeleteGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteGroup'
type Backend_DeleteGroup_Call struct {
	*mock.Call
}

// DeleteGroup is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *Backend_Expecter) DeleteGroup(ctx interface{}, id interface{}) *Backend_DeleteGroup_Call {
	return &Backend_DeleteGroup_Call{Call: _e.mock.On("DeleteGroup", ctx, id)}
}

func (_c *Backend_DeleteGroup_Call) Run(run func(ctx context.Context, id string)) *Backend_DeleteGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Backend_DeleteGroup_Call) Return(_a0 error) *Backend_DeleteGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_DeleteGroup_Call) RunAndReturn(run func(context.Context, string) error) *Backend_DeleteGroup_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUser provides a mock function with given fields: ctx, nameOrID
func (_m *Backend) DeleteUser(ctx context.Context, nameOrID string) error {
	ret := _m.Called(ctx, nameOrID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, nameOrID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type Backend_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - ctx context.Context
//   - nameOrID string
func (_e *Backend_Expecter) DeleteUser(ctx interface{}, nameOrID interface{}) *Backend_DeleteUser_Call {
	return &Backend_DeleteUser_Call{Call: _e.mock.On("DeleteUser", ctx, nameOrID)}
}

func (_c *Backend_DeleteUser_Call) Run(run func(ctx context.Context, nameOrID string)) *Backend_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Backend_DeleteUser_Call) Return(_a0 error) *Backend_DeleteUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_DeleteUser_Call) RunAndReturn(run func(context.Context, string) error) *Backend_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetGroup provides a mock function with given fields: ctx, nameOrID, queryParam
func (_m *Backend) GetGroup(ctx context.Context, nameOrID string, queryParam url.Values) (*libregraph.Group, error) {
	ret := _m.Called(ctx, nameOrID, queryParam)

	if len(ret) == 0 {
		panic("no return value specified for GetGroup")
	}

	var r0 *libregraph.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, url.Values) (*libregraph.Group, error)); ok {
		return rf(ctx, nameOrID, queryParam)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, url.Values) *libregraph.Group); ok {
		r0 = rf(ctx, nameOrID, queryParam)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*libregraph.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, url.Values) error); ok {
		r1 = rf(ctx, nameOrID, queryParam)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Backend_GetGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGroup'
type Backend_GetGroup_Call struct {
	*mock.Call
}

// GetGroup is a helper method to define mock.On call
//   - ctx context.Context
//   - nameOrID string
//   - queryParam url.Values
func (_e *Backend_Expecter) GetGroup(ctx interface{}, nameOrID interface{}, queryParam interface{}) *Backend_GetGroup_Call {
	return &Backend_GetGroup_Call{Call: _e.mock.On("GetGroup", ctx, nameOrID, queryParam)}
}

func (_c *Backend_GetGroup_Call) Run(run func(ctx context.Context, nameOrID string, queryParam url.Values)) *Backend_GetGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(url.Values))
	})
	return _c
}

func (_c *Backend_GetGroup_Call) Return(_a0 *libregraph.Group, _a1 error) *Backend_GetGroup_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Backend_GetGroup_Call) RunAndReturn(run func(context.Context, string, url.Values) (*libregraph.Group, error)) *Backend_GetGroup_Call {
	_c.Call.Return(run)
	return _c
}

// GetGroupMembers provides a mock function with given fields: ctx, id, oreq
func (_m *Backend) GetGroupMembers(ctx context.Context, id string, oreq *godata.GoDataRequest) ([]*libregraph.User, error) {
	ret := _m.Called(ctx, id, oreq)

	if len(ret) == 0 {
		panic("no return value specified for GetGroupMembers")
	}

	var r0 []*libregraph.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *godata.GoDataRequest) ([]*libregraph.User, error)); ok {
		return rf(ctx, id, oreq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *godata.GoDataRequest) []*libregraph.User); ok {
		r0 = rf(ctx, id, oreq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*libregraph.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *godata.GoDataRequest) error); ok {
		r1 = rf(ctx, id, oreq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Backend_GetGroupMembers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGroupMembers'
type Backend_GetGroupMembers_Call struct {
	*mock.Call
}

// GetGroupMembers is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - oreq *godata.GoDataRequest
func (_e *Backend_Expecter) GetGroupMembers(ctx interface{}, id interface{}, oreq interface{}) *Backend_GetGroupMembers_Call {
	return &Backend_GetGroupMembers_Call{Call: _e.mock.On("GetGroupMembers", ctx, id, oreq)}
}

func (_c *Backend_GetGroupMembers_Call) Run(run func(ctx context.Context, id string, oreq *godata.GoDataRequest)) *Backend_GetGroupMembers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*godata.GoDataRequest))
	})
	return _c
}

func (_c *Backend_GetGroupMembers_Call) Return(_a0 []*libregraph.User, _a1 error) *Backend_GetGroupMembers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Backend_GetGroupMembers_Call) RunAndReturn(run func(context.Context, string, *godata.GoDataRequest) ([]*libregraph.User, error)) *Backend_GetGroupMembers_Call {
	_c.Call.Return(run)
	return _c
}

// GetGroups provides a mock function with given fields: ctx, oreq
func (_m *Backend) GetGroups(ctx context.Context, oreq *godata.GoDataRequest) ([]*libregraph.Group, error) {
	ret := _m.Called(ctx, oreq)

	if len(ret) == 0 {
		panic("no return value specified for GetGroups")
	}

	var r0 []*libregraph.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *godata.GoDataRequest) ([]*libregraph.Group, error)); ok {
		return rf(ctx, oreq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *godata.GoDataRequest) []*libregraph.Group); ok {
		r0 = rf(ctx, oreq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*libregraph.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *godata.GoDataRequest) error); ok {
		r1 = rf(ctx, oreq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Backend_GetGroups_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGroups'
type Backend_GetGroups_Call struct {
	*mock.Call
}

// GetGroups is a helper method to define mock.On call
//   - ctx context.Context
//   - oreq *godata.GoDataRequest
func (_e *Backend_Expecter) GetGroups(ctx interface{}, oreq interface{}) *Backend_GetGroups_Call {
	return &Backend_GetGroups_Call{Call: _e.mock.On("GetGroups", ctx, oreq)}
}

func (_c *Backend_GetGroups_Call) Run(run func(ctx context.Context, oreq *godata.GoDataRequest)) *Backend_GetGroups_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*godata.GoDataRequest))
	})
	return _c
}

func (_c *Backend_GetGroups_Call) Return(_a0 []*libregraph.Group, _a1 error) *Backend_GetGroups_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Backend_GetGroups_Call) RunAndReturn(run func(context.Context, *godata.GoDataRequest) ([]*libregraph.Group, error)) *Backend_GetGroups_Call {
	_c.Call.Return(run)
	return _c
}

// GetUser provides a mock function with given fields: ctx, nameOrID, oreq
func (_m *Backend) GetUser(ctx context.Context, nameOrID string, oreq *godata.GoDataRequest) (*libregraph.User, error) {
	ret := _m.Called(ctx, nameOrID, oreq)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 *libregraph.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *godata.GoDataRequest) (*libregraph.User, error)); ok {
		return rf(ctx, nameOrID, oreq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *godata.GoDataRequest) *libregraph.User); ok {
		r0 = rf(ctx, nameOrID, oreq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*libregraph.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *godata.GoDataRequest) error); ok {
		r1 = rf(ctx, nameOrID, oreq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Backend_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type Backend_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - ctx context.Context
//   - nameOrID string
//   - oreq *godata.GoDataRequest
func (_e *Backend_Expecter) GetUser(ctx interface{}, nameOrID interface{}, oreq interface{}) *Backend_GetUser_Call {
	return &Backend_GetUser_Call{Call: _e.mock.On("GetUser", ctx, nameOrID, oreq)}
}

func (_c *Backend_GetUser_Call) Run(run func(ctx context.Context, nameOrID string, oreq *godata.GoDataRequest)) *Backend_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*godata.GoDataRequest))
	})
	return _c
}

func (_c *Backend_GetUser_Call) Return(_a0 *libregraph.User, _a1 error) *Backend_GetUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Backend_GetUser_Call) RunAndReturn(run func(context.Context, string, *godata.GoDataRequest) (*libregraph.User, error)) *Backend_GetUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUsers provides a mock function with given fields: ctx, oreq
func (_m *Backend) GetUsers(ctx context.Context, oreq *godata.GoDataRequest) ([]*libregraph.User, error) {
	ret := _m.Called(ctx, oreq)

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 []*libregraph.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *godata.GoDataRequest) ([]*libregraph.User, error)); ok {
		return rf(ctx, oreq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *godata.GoDataRequest) []*libregraph.User); ok {
		r0 = rf(ctx, oreq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*libregraph.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *godata.GoDataRequest) error); ok {
		r1 = rf(ctx, oreq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Backend_GetUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUsers'
type Backend_GetUsers_Call struct {
	*mock.Call
}

// GetUsers is a helper method to define mock.On call
//   - ctx context.Context
//   - oreq *godata.GoDataRequest
func (_e *Backend_Expecter) GetUsers(ctx interface{}, oreq interface{}) *Backend_GetUsers_Call {
	return &Backend_GetUsers_Call{Call: _e.mock.On("GetUsers", ctx, oreq)}
}

func (_c *Backend_GetUsers_Call) Run(run func(ctx context.Context, oreq *godata.GoDataRequest)) *Backend_GetUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*godata.GoDataRequest))
	})
	return _c
}

func (_c *Backend_GetUsers_Call) Return(_a0 []*libregraph.User, _a1 error) *Backend_GetUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Backend_GetUsers_Call) RunAndReturn(run func(context.Context, *godata.GoDataRequest) ([]*libregraph.User, error)) *Backend_GetUsers_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveMemberFromGroup provides a mock function with given fields: ctx, groupID, memberID
func (_m *Backend) RemoveMemberFromGroup(ctx context.Context, groupID string, memberID string) error {
	ret := _m.Called(ctx, groupID, memberID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveMemberFromGroup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, groupID, memberID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_RemoveMemberFromGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveMemberFromGroup'
type Backend_RemoveMemberFromGroup_Call struct {
	*mock.Call
}

// RemoveMemberFromGroup is a helper method to define mock.On call
//   - ctx context.Context
//   - groupID string
//   - memberID string
func (_e *Backend_Expecter) RemoveMemberFromGroup(ctx interface{}, groupID interface{}, memberID interface{}) *Backend_RemoveMemberFromGroup_Call {
	return &Backend_RemoveMemberFromGroup_Call{Call: _e.mock.On("RemoveMemberFromGroup", ctx, groupID, memberID)}
}

func (_c *Backend_RemoveMemberFromGroup_Call) Run(run func(ctx context.Context, groupID string, memberID string)) *Backend_RemoveMemberFromGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Backend_RemoveMemberFromGroup_Call) Return(_a0 error) *Backend_RemoveMemberFromGroup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_RemoveMemberFromGroup_Call) RunAndReturn(run func(context.Context, string, string) error) *Backend_RemoveMemberFromGroup_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateGroupName provides a mock function with given fields: ctx, groupID, groupName
func (_m *Backend) UpdateGroupName(ctx context.Context, groupID string, groupName string) error {
	ret := _m.Called(ctx, groupID, groupName)

	if len(ret) == 0 {
		panic("no return value specified for UpdateGroupName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, groupID, groupName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_UpdateGroupName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateGroupName'
type Backend_UpdateGroupName_Call struct {
	*mock.Call
}

// UpdateGroupName is a helper method to define mock.On call
//   - ctx context.Context
//   - groupID string
//   - groupName string
func (_e *Backend_Expecter) UpdateGroupName(ctx interface{}, groupID interface{}, groupName interface{}) *Backend_UpdateGroupName_Call {
	return &Backend_UpdateGroupName_Call{Call: _e.mock.On("UpdateGroupName", ctx, groupID, groupName)}
}

func (_c *Backend_UpdateGroupName_Call) Run(run func(ctx context.Context, groupID string, groupName string)) *Backend_UpdateGroupName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Backend_UpdateGroupName_Call) Return(_a0 error) *Backend_UpdateGroupName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_UpdateGroupName_Call) RunAndReturn(run func(context.Context, string, string) error) *Backend_UpdateGroupName_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: ctx, nameOrID, user
func (_m *Backend) UpdateUser(ctx context.Context, nameOrID string, user libregraph.User) (*libregraph.User, error) {
	ret := _m.Called(ctx, nameOrID, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 *libregraph.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, libregraph.User) (*libregraph.User, error)); ok {
		return rf(ctx, nameOrID, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, libregraph.User) *libregraph.User); ok {
		r0 = rf(ctx, nameOrID, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*libregraph.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, libregraph.User) error); ok {
		r1 = rf(ctx, nameOrID, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Backend_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type Backend_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - nameOrID string
//   - user libregraph.User
func (_e *Backend_Expecter) UpdateUser(ctx interface{}, nameOrID interface{}, user interface{}) *Backend_UpdateUser_Call {
	return &Backend_UpdateUser_Call{Call: _e.mock.On("UpdateUser", ctx, nameOrID, user)}
}

func (_c *Backend_UpdateUser_Call) Run(run func(ctx context.Context, nameOrID string, user libregraph.User)) *Backend_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(libregraph.User))
	})
	return _c
}

func (_c *Backend_UpdateUser_Call) Return(_a0 *libregraph.User, _a1 error) *Backend_UpdateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Backend_UpdateUser_Call) RunAndReturn(run func(context.Context, string, libregraph.User) (*libregraph.User, error)) *Backend_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewBackend creates a new instance of Backend. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBackend(t interface {
	mock.TestingT
	Cleanup(func())
}) *Backend {
	mock := &Backend{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
