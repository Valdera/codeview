// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import entity "codeview/internal/entity"
import mock "github.com/stretchr/testify/mock"
import pagination "codeview/internal/util/pagination"

// QuestionRepository is an autogenerated mock type for the QuestionRepository type
type QuestionRepository struct {
	mock.Mock
}

// CreateQuestion provides a mock function with given fields: ctx, body
func (_m *QuestionRepository) CreateQuestion(ctx context.Context, body *entity.Question) (*entity.Question, error) {
	ret := _m.Called(ctx, body)

	var r0 *entity.Question
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Question) *entity.Question); ok {
		r0 = rf(ctx, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Question)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Question) error); ok {
		r1 = rf(ctx, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteQuestionById provides a mock function with given fields: ctx, id
func (_m *QuestionRepository) DeleteQuestionById(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetQuestionById provides a mock function with given fields: ctx, id
func (_m *QuestionRepository) GetQuestionById(ctx context.Context, id uint) (*entity.Question, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Question
	if rf, ok := ret.Get(0).(func(context.Context, uint) *entity.Question); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Question)
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

// GetQuestions provides a mock function with given fields: ctx, p
func (_m *QuestionRepository) GetQuestions(ctx context.Context, p *pagination.Pagination) ([]entity.Question, error) {
	ret := _m.Called(ctx, p)

	var r0 []entity.Question
	if rf, ok := ret.Get(0).(func(context.Context, *pagination.Pagination) []entity.Question); ok {
		r0 = rf(ctx, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Question)
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

// GetQuestionsByIds provides a mock function with given fields: ctx, ids, p
func (_m *QuestionRepository) GetQuestionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Question, error) {
	ret := _m.Called(ctx, ids, p)

	var r0 []entity.Question
	if rf, ok := ret.Get(0).(func(context.Context, []uint, *pagination.Pagination) []entity.Question); ok {
		r0 = rf(ctx, ids, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Question)
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

// UpdateQuestionById provides a mock function with given fields: ctx, id, body
func (_m *QuestionRepository) UpdateQuestionById(ctx context.Context, id uint, body *entity.Question) (*entity.Question, error) {
	ret := _m.Called(ctx, id, body)

	var r0 *entity.Question
	if rf, ok := ret.Get(0).(func(context.Context, uint, *entity.Question) *entity.Question); ok {
		r0 = rf(ctx, id, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Question)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, *entity.Question) error); ok {
		r1 = rf(ctx, id, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
