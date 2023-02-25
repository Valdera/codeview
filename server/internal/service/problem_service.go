package service

import (
	"codeview/internal/dto/request"
	"codeview/internal/dto/response"

	"codeview/utils/pagination"
	"context"
)

type ProblemService interface {
	CreateProblem(ctx context.Context, req *request.ProblemCreate) (*response.Problem, error)
	GetProblems(ctx context.Context, p *pagination.Pagination) ([]response.Problem, error)
	GetProblemById(ctx context.Context, id uint) (*response.Problem, error)
	GetProblemsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Problem, error)
	UpdateProblemById(ctx context.Context, id uint, body *request.ProblemUpdate) (*response.Problem, error)
	DeleteProblemById(ctx context.Context, id uint) error

	CreateTag(ctx context.Context, req *request.TagCreate) (*response.Tag, error)
	GetTags(ctx context.Context, p *pagination.Pagination) ([]response.Tag, error)
	GetTagById(ctx context.Context, id uint) (*response.Tag, error)
	GetTagsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Tag, error)
	UpdateTagById(ctx context.Context, id uint, body *request.TagUpdate) (*response.Tag, error)
	DeleteTagById(ctx context.Context, id uint) error

	CreateSource(ctx context.Context, req *request.SourceCreate) (*response.Source, error)
	GetSources(ctx context.Context, p *pagination.Pagination) ([]response.Source, error)
	GetSourceById(ctx context.Context, id uint) (*response.Source, error)
	GetSourcesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Source, error)
	UpdateSourceById(ctx context.Context, id uint, body *request.SourceUpdate) (*response.Source, error)
	DeleteSourceById(ctx context.Context, id uint) error

	CreateDifficulty(ctx context.Context, req *request.DifficultyCreate) (*response.Difficulty, error)
	GetDifficulties(ctx context.Context, p *pagination.Pagination) ([]response.Difficulty, error)
	GetDifficultyById(ctx context.Context, id uint) (*response.Difficulty, error)
	GetDifficultiesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Difficulty, error)
	UpdateDifficultyById(ctx context.Context, id uint, body *request.DifficultyUpdate) (*response.Difficulty, error)
	DeleteDifficultyById(ctx context.Context, id uint) error

	CreateSolution(ctx context.Context, req *request.SolutionCreate) (*response.Solution, error)
	GetSolutions(ctx context.Context, p *pagination.Pagination) ([]response.Solution, error)
	GetSolutionById(ctx context.Context, id uint) (*response.Solution, error)
	GetSolutionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Solution, error)
	UpdateSolutionById(ctx context.Context, id uint, body *request.SolutionUpdate) (*response.Solution, error)
	DeleteSolutionById(ctx context.Context, id uint) error

	CreateQuestion(ctx context.Context, req *request.QuestionCreate) (*response.Question, error)
	GetQuestions(ctx context.Context, p *pagination.Pagination) ([]response.Question, error)
	GetQuestionById(ctx context.Context, id uint) (*response.Question, error)
	GetQuestionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Question, error)
	UpdateQuestionById(ctx context.Context, id uint, body *request.QuestionUpdate) (*response.Question, error)
	DeleteQuestionById(ctx context.Context, id uint) error
}
