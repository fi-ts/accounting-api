// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"

	v1 "github.com/fi-ts/accounting-go/pkg/apis/v1"
	mock "github.com/stretchr/testify/mock"
)

// S3ServiceServer is an autogenerated mock type for the S3ServiceServer type
type S3ServiceServer struct {
	mock.Mock
}

// BucketStats provides a mock function with given fields: _a0, _a1
func (_m *S3ServiceServer) BucketStats(_a0 context.Context, _a1 *v1.S3BucketReport) (*v1.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *v1.S3BucketReport) *v1.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.S3BucketReport) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Usage provides a mock function with given fields: _a0, _a1
func (_m *S3ServiceServer) Usage(_a0 context.Context, _a1 *v1.S3UsageRequest) (*v1.S3UsageResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.S3UsageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1.S3UsageRequest) *v1.S3UsageResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.S3UsageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.S3UsageRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewS3ServiceServerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewS3ServiceServer creates a new instance of S3ServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewS3ServiceServer(t NewS3ServiceServerT) *S3ServiceServer {
	mock := &S3ServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
