// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import pagination "codeview/internal/util/pagination"
import request "codeview/internal/dto/request"
import response "codeview/internal/dto/response"

// TagService is an autogenerated mock type for the TagService type
type TagService struct {
	mock.Mock
}

// CreateTag provides a mock function with given fields: ctx, req
func (_m *TagService) CreateTag(ctx context.Context, req *request.TagCreate) (*response.Tag, error) {
	ret := _m.Called(ctx, req)

	var r0 *response.Tag
	if rf, ok := ret.Get(0).(func(context.Context, *request.TagCreate) *response.Tag); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *request.TagCreate) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTagById provides a mock function with given fields: ctx, id
func (_m *TagService) DeleteTagById(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTagById provides a mock function with given fields: ctx, id
func (_m *TagService) GetTagById(ctx context.Context, id uint) (*response.Tag, error) {
	ret := _m.Called(ctx, id)

	var r0 *response.Tag
	if rf, ok := ret.Get(0).(func(context.Context, uint) *response.Tag); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTags provides a mock function with given fields: ctx, p
func (_m *TagService) GetTags(ctx context.Context, p *pagination.Pagination) ([]response.Tag, error) {
	ret := _m.Called(ctx, p)

	var r0 []response.Tag
	if rf, ok := ret.Get(0).(func(context.Context, *pagination.Pagination) []response.Tag); ok {
		r0 = rf(ctx, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]response.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pagination.Pagination) error); ok {
		r1 = rf(ctx, p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTagsByIds provides a mock function with given fields: ctx, ids, p
func (_m *TagService) GetTagsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Tag, error) {
	ret := _m.Called(ctx, ids, p)

	var r0 []response.Tag
	if rf, ok := ret.Get(0).(func(context.Context, []uint, *pagination.Pagination) []response.Tag); ok {
		r0 = rf(ctx, ids, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]response.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []uint, *pagination.Pagination) error); ok {
		r1 = rf(ctx, ids, p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTagById provides a mock function with given fields: ctx, id, body
func (_m *TagService) UpdateTagById(ctx context.Context, id uint, body *request.TagUpdate) (*response.Tag, error) {
	ret := _m.Called(ctx, id, body)

	var r0 *response.Tag
	if rf, ok := ret.Get(0).(func(context.Context, uint, *request.TagUpdate) *response.Tag); ok {
		r0 = rf(ctx, id, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, *request.TagUpdate) error); ok {
		r1 = rf(ctx, id, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
