package domain

import (
	"codeview/internal/dto"
	"codeview/internal/entity"
	"codeview/utils/pagination"
	"context"
)

type ProblemRepository interface {
	CreateProblem(ctx context.Context, body *entity.Problem) (*entity.Problem, error)
	GetProblems(ctx context.Context, p *pagination.Pagination) ([]entity.Problem, error)
	GetProblemById(ctx context.Context, id uint) (*entity.Problem, error)
	GetProblemsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Problem, error)
	UpdateProblemById(ctx context.Context, id uint, body *entity.Problem) (*entity.Problem, error)
	DeleteProblemById(ctx context.Context, id uint) error

	CreateTag(ctx context.Context, body *entity.Tag) (*entity.Tag, error)
	GetTags(ctx context.Context, p *pagination.Pagination) ([]entity.Tag, error)
	GetTagById(ctx context.Context, id uint) (*entity.Tag, error)
	GetTagsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Tag, error)
	UpdateTagById(ctx context.Context, id uint, body *entity.Tag) (*entity.Tag, error)
	DeleteTagById(ctx context.Context, id uint) error

	CreateSource(ctx context.Context, body *entity.Source) (*entity.Source, error)
	GetSources(ctx context.Context, p *pagination.Pagination) ([]entity.Source, error)
	GetSourceById(ctx context.Context, id uint) (*entity.Source, error)
	GetSourcesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Source, error)
	UpdateSourceById(ctx context.Context, id uint, body *entity.Source) (*entity.Source, error)
	DeleteSourceById(ctx context.Context, id uint) error

	CreateDifficulty(ctx context.Context, body *entity.Difficulty) (*entity.Difficulty, error)
	GetDifficulties(ctx context.Context, p *pagination.Pagination) ([]entity.Difficulty, error)
	GetDifficultyById(ctx context.Context, id uint) (*entity.Difficulty, error)
	GetDifficultiesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Difficulty, error)
	UpdateDifficultyById(ctx context.Context, id uint, body *entity.Difficulty) (*entity.Difficulty, error)
	DeleteDifficultyById(ctx context.Context, id uint) error

	CreateSolution(ctx context.Context, body *entity.Solution) (*entity.Solution, error)
	GetSolutions(ctx context.Context, p *pagination.Pagination) ([]entity.Solution, error)
	GetSolutionById(ctx context.Context, id uint) (*entity.Solution, error)
	GetSolutionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Solution, error)
	UpdateSolutionById(ctx context.Context, id uint, body *entity.Solution) (*entity.Solution, error)
	DeleteSolutionById(ctx context.Context, id uint) error

	CreateQuestion(ctx context.Context, body *entity.Question) (*entity.Question, error)
	GetQuestions(ctx context.Context, p *pagination.Pagination) ([]entity.Question, error)
	GetQuestionById(ctx context.Context, id uint) (*entity.Question, error)
	GetQuestionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Question, error)
	UpdateQuestionById(ctx context.Context, id uint, body *entity.Question) (*entity.Question, error)
	DeleteQuestionById(ctx context.Context, id uint) error
}

type ProblemService interface {
	CreateProblem(ctx context.Context, req *dto.ProblemCreate) (*dto.ProblemResponse, error)
	GetProblems(ctx context.Context, p *pagination.Pagination) ([]dto.ProblemResponse, error)
	GetProblemById(ctx context.Context, id uint) (*dto.ProblemResponse, error)
	GetProblemsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]dto.ProblemResponse, error)
	UpdateProblemById(ctx context.Context, id uint, body *dto.ProblemUpdate) (*dto.ProblemResponse, error)
	DeleteProblemById(ctx context.Context, id uint) error

	CreateTag(ctx context.Context, req *dto.TagCreate) (*dto.TagResponse, error)
	GetTags(ctx context.Context, p *pagination.Pagination) ([]dto.TagResponse, error)
	GetTagById(ctx context.Context, id uint) (*dto.TagResponse, error)
	GetTagsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]dto.TagResponse, error)
	UpdateTagById(ctx context.Context, id uint, body *dto.TagUpdate) (*dto.TagResponse, error)
	DeleteTagById(ctx context.Context, id uint) error

	CreateSource(ctx context.Context, req *dto.SourceCreate) (*dto.SourceResponse, error)
	GetSources(ctx context.Context, p *pagination.Pagination) ([]dto.SourceResponse, error)
	GetSourceById(ctx context.Context, id uint) (*dto.SourceResponse, error)
	GetSourcesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]dto.SourceResponse, error)
	UpdateSourceById(ctx context.Context, id uint, body *dto.SourceUpdate) (*dto.SourceResponse, error)
	DeleteSourceById(ctx context.Context, id uint) error

	CreateDifficulty(ctx context.Context, req *dto.DifficultyCreate) (*dto.DifficultyResponse, error)
	GetDifficulties(ctx context.Context, p *pagination.Pagination) ([]dto.DifficultyResponse, error)
	GetDifficultyById(ctx context.Context, id uint) (*dto.DifficultyResponse, error)
	GetDifficultiesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]dto.DifficultyResponse, error)
	UpdateDifficultyById(ctx context.Context, id uint, body *dto.DifficultyUpdate) (*dto.DifficultyResponse, error)
	DeleteDifficultyById(ctx context.Context, id uint) error

	CreateSolution(ctx context.Context, req *dto.SolutionCreate) (*dto.SolutionResponse, error)
	GetSolutions(ctx context.Context, p *pagination.Pagination) ([]dto.SolutionResponse, error)
	GetSolutionById(ctx context.Context, id uint) (*dto.SolutionResponse, error)
	GetSolutionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]dto.SolutionResponse, error)
	UpdateSolutionById(ctx context.Context, id uint, body *dto.SolutionUpdate) (*dto.SolutionResponse, error)
	DeleteSolutionById(ctx context.Context, id uint) error

	CreateQuestion(ctx context.Context, req *dto.QuestionCreate) (*dto.QuestionResponse, error)
	GetQuestions(ctx context.Context, p *pagination.Pagination) ([]dto.QuestionResponse, error)
	GetQuestionById(ctx context.Context, id uint) (*dto.QuestionResponse, error)
	GetQuestionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]dto.QuestionResponse, error)
	UpdateQuestionById(ctx context.Context, id uint, body *dto.QuestionUpdate) (*dto.QuestionResponse, error)
	DeleteQuestionById(ctx context.Context, id uint) error
}
