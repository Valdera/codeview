// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import entity "codeview/internal/entity"
import mock "github.com/stretchr/testify/mock"
import pagination "codeview/internal/util/pagination"

// SourceRepository is an autogenerated mock type for the SourceRepository type
type SourceRepository struct {
	mock.Mock
}

// CreateSource provides a mock function with given fields: ctx, body
func (_m *SourceRepository) CreateSource(ctx context.Context, body *entity.Source) (*entity.Source, error) {
	ret := _m.Called(ctx, body)

	var r0 *entity.Source
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Source) *entity.Source); ok {
		r0 = rf(ctx, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Source)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Source) error); ok {
		r1 = rf(ctx, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSourceById provides a mock function with given fields: ctx, id
func (_m *SourceRepository) DeleteSourceById(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSourceById provides a mock function with given fields: ctx, id
func (_m *SourceRepository) GetSourceById(ctx context.Context, id uint) (*entity.Source, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Source
	if rf, ok := ret.Get(0).(func(context.Context, uint) *entity.Source); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Source)
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

// GetSources provides a mock function with given fields: ctx, p
func (_m *SourceRepository) GetSources(ctx context.Context, p *pagination.Pagination) ([]entity.Source, error) {
	ret := _m.Called(ctx, p)

	var r0 []entity.Source
	if rf, ok := ret.Get(0).(func(context.Context, *pagination.Pagination) []entity.Source); ok {
		r0 = rf(ctx, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Source)
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

// GetSourcesByIds provides a mock function with given fields: ctx, ids, p
func (_m *SourceRepository) GetSourcesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Source, error) {
	ret := _m.Called(ctx, ids, p)

	var r0 []entity.Source
	if rf, ok := ret.Get(0).(func(context.Context, []uint, *pagination.Pagination) []entity.Source); ok {
		r0 = rf(ctx, ids, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Source)
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

// UpdateSourceById provides a mock function with given fields: ctx, id, body
func (_m *SourceRepository) UpdateSourceById(ctx context.Context, id uint, body *entity.Source) (*entity.Source, error) {
	ret := _m.Called(ctx, id, body)

	var r0 *entity.Source
	if rf, ok := ret.Get(0).(func(context.Context, uint, *entity.Source) *entity.Source); ok {
		r0 = rf(ctx, id, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Source)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, *entity.Source) error); ok {
		r1 = rf(ctx, id, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
