// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import request "codeview/internal/dto/request"

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, req
func (_m *AuthService) Login(ctx context.Context, req *request.Login) (string, error) {
	ret := _m.Called(ctx, req)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *request.Login) string); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *request.Login) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, req
func (_m *AuthService) Register(ctx context.Context, req *request.Register) (string, error) {
	ret := _m.Called(ctx, req)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *request.Register) string); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *request.Register) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
