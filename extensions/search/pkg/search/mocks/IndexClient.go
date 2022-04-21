// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	providerv1beta1 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	mock "github.com/stretchr/testify/mock"

	v0 "github.com/owncloud/ocis/protogen/gen/ocis/services/search/v0"
)

// IndexClient is an autogenerated mock type for the IndexClient type
type IndexClient struct {
	mock.Mock
}

// Add provides a mock function with given fields: ref, ri
func (_m *IndexClient) Add(ref *providerv1beta1.Reference, ri *providerv1beta1.ResourceInfo) error {
	ret := _m.Called(ref, ri)

	var r0 error
	if rf, ok := ret.Get(0).(func(*providerv1beta1.Reference, *providerv1beta1.ResourceInfo) error); ok {
		r0 = rf(ref, ri)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Search provides a mock function with given fields: ctx, req
func (_m *IndexClient) Search(ctx context.Context, req *v0.SearchIndexRequest) (*v0.SearchIndexResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *v0.SearchIndexResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v0.SearchIndexRequest) *v0.SearchIndexResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0.SearchIndexResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v0.SearchIndexRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
