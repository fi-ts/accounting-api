// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/fi-ts/accounting-go/pkg/apis/v1"
)

// InfoServiceClient is an autogenerated mock type for the InfoServiceClient type
type InfoServiceClient struct {
	mock.Mock
}

// Projects provides a mock function with given fields: ctx, in, opts
func (_m *InfoServiceClient) Projects(ctx context.Context, in *v1.ProjectInfoRequest, opts ...grpc.CallOption) (*v1.ProjectInfoResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.ProjectInfoResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ProjectInfoRequest, ...grpc.CallOption) (*v1.ProjectInfoResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ProjectInfoRequest, ...grpc.CallOption) *v1.ProjectInfoResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ProjectInfoResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ProjectInfoRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tenants provides a mock function with given fields: ctx, in, opts
func (_m *InfoServiceClient) Tenants(ctx context.Context, in *v1.TenantInfoRequest, opts ...grpc.CallOption) (*v1.TenantInfoResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.TenantInfoResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.TenantInfoRequest, ...grpc.CallOption) (*v1.TenantInfoResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.TenantInfoRequest, ...grpc.CallOption) *v1.TenantInfoResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.TenantInfoResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.TenantInfoRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewInfoServiceClient creates a new instance of InfoServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInfoServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *InfoServiceClient {
	mock := &InfoServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
