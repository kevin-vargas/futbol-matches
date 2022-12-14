// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	model "backend/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Auth is an autogenerated mock type for the Auth type
type Auth struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, username, password
func (_m *Auth) Login(ctx context.Context, username string, password string) (string, error) {
	ret := _m.Called(ctx, username, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, username, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignUp provides a mock function with given fields: ctx, u
func (_m *Auth) SignUp(ctx context.Context, u *model.User) (bool, error) {
	ret := _m.Called(ctx, u)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) bool); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.User) error); ok {
		r1 = rf(ctx, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
