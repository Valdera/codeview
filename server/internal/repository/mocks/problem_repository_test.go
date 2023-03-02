// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import entity "codeview/internal/entity"
import mock "github.com/stretchr/testify/mock"
import pagination "codeview/internal/util/pagination"

// ProblemRepository is an autogenerated mock type for the ProblemRepository type
type ProblemRepository struct {
	mock.Mock
}

// CreateProblem provides a mock function with given fields: ctx, body
func (_m *ProblemRepository) CreateProblem(ctx context.Context, body *entity.Problem) (*entity.Problem, error) {
	ret := _m.Called(ctx, body)

	var r0 *entity.Problem
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Problem) *entity.Problem); ok {
		r0 = rf(ctx, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Problem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Problem) error); ok {
		r1 = rf(ctx, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProblemById provides a mock function with given fields: ctx, id
func (_m *ProblemRepository) DeleteProblemById(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetProblemById provides a mock function with given fields: ctx, id
func (_m *ProblemRepository) GetProblemById(ctx context.Context, id uint) (*entity.Problem, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Problem
	if rf, ok := ret.Get(0).(func(context.Context, uint) *entity.Problem); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Problem)
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

// GetProblems provides a mock function with given fields: ctx, p
func (_m *ProblemRepository) GetProblems(ctx context.Context, p *pagination.Pagination) ([]entity.Problem, error) {
	ret := _m.Called(ctx, p)

	var r0 []entity.Problem
	if rf, ok := ret.Get(0).(func(context.Context, *pagination.Pagination) []entity.Problem); ok {
		r0 = rf(ctx, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Problem)
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

// GetProblemsByIds provides a mock function with given fields: ctx, ids, p
func (_m *ProblemRepository) GetProblemsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Problem, error) {
	ret := _m.Called(ctx, ids, p)

	var r0 []entity.Problem
	if rf, ok := ret.Get(0).(func(context.Context, []uint, *pagination.Pagination) []entity.Problem); ok {
		r0 = rf(ctx, ids, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Problem)
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

// UpdateProblemById provides a mock function with given fields: ctx, id, body
func (_m *ProblemRepository) UpdateProblemById(ctx context.Context, id uint, body *entity.Problem) (*entity.Problem, error) {
	ret := _m.Called(ctx, id, body)

	var r0 *entity.Problem
	if rf, ok := ret.Get(0).(func(context.Context, uint, *entity.Problem) *entity.Problem); ok {
		r0 = rf(ctx, id, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Problem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, *entity.Problem) error); ok {
		r1 = rf(ctx, id, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
