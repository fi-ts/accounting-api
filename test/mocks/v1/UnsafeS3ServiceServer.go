// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeS3ServiceServer is an autogenerated mock type for the UnsafeS3ServiceServer type
type UnsafeS3ServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedS3ServiceServer provides a mock function with given fields:
func (_m *UnsafeS3ServiceServer) mustEmbedUnimplementedS3ServiceServer() {
	_m.Called()
}

// NewUnsafeS3ServiceServer creates a new instance of UnsafeS3ServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafeS3ServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafeS3ServiceServer {
	mock := &UnsafeS3ServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
