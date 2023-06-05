package service

import (
	"codeview/config"
	"codeview/internal/constant"
	"codeview/internal/dto/request"
	"codeview/internal/dto/response"
	"codeview/internal/entity"
	"codeview/internal/mapper"
	"codeview/internal/repository"
	"codeview/internal/service"
	"codeview/internal/util/pagination"
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type problemService struct {
	config       config.AppConfig
	db           *gorm.DB
	problemRepo  repository.ProblemRepository
	questionRepo repository.QuestionRepository
	solutionRepo repository.SolutionRepository
}

func NewProblemService(
	config config.AppConfig,
	db *gorm.DB,
	problemRepo repository.ProblemRepository,
	questionRepo repository.QuestionRepository,
	solutionRepo repository.SolutionRepository) service.ProblemService {
	return &problemService{
		config,
		db,
		problemRepo,
		questionRepo,
		solutionRepo,
	}
}

func (s *problemService) CreateProblem(ctx context.Context, body *request.ProblemCreate) (*response.Problem, error) {
	var response *response.Problem

	if err := s.db.Transaction(func(trx *gorm.DB) error {
		trx = s.db.Begin()
		defer func() {
			if r := recover(); r != nil {
				trx.Rollback()
			}
		}()

		problem, err := s.problemRepo.WithTrx(trx).CreateProblem(ctx, &entity.Problem{
			Title:        body.Title,
			DifficultyID: uuid.Must(uuid.Parse(body.DifficultyID)),
			Rating:       &body.Rating,
			Emoji:        body.Emoji,
		})
		if err != nil {
			trx.Rollback()
			log.Printf("[ERROR] Problem Service - CreateProblem : %v\n", err)
			return err
		}

		question, err := s.questionRepo.WithTrx(trx).CreateQuestion(ctx, &entity.Question{
			ProblemID: problem.ID,
			Content:   constant.QUESTION_DEFAULT,
		})
		if err != nil {
			trx.Rollback()
			log.Printf("[ERROR] Problem Service - CreateProblem : %v\n", err)
			return err
		}

		problem.Questions = append(problem.Questions, *question)

		response, err = mapper.ProblemToDTO(problem)
		if err != nil {
			log.Printf("[ERROR] Problem Service - CreateProblem : %v\n", err)
			return err
		}

		if err := trx.Commit().Error; err != nil {
			trx.Rollback()
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *problemService) GetProblems(ctx context.Context, p *pagination.Pagination) ([]response.Problem, error) {
	problems, err := s.problemRepo.GetProblems(ctx, p)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetProblems : %v\n", err)
		return nil, err
	}

	responses := make([]response.Problem, len(problems))
	for i, problem := range problems {
		response, err := mapper.ProblemToDTO(&problem)
		if err != nil {
			log.Printf("[ERROR] Problem Service - GetProblems : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *problemService) GetProblemById(ctx context.Context, id uuid.UUID) (*response.Problem, error) {
	problem, err := s.problemRepo.GetProblemById(ctx, id)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetProblemById : %v\n", err)
		return nil, err
	}

	response, err := mapper.ProblemToDTO(problem)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetProblemById : %v\n", err)
		return nil, err
	}

	return response, nil
}

func (s *problemService) GetProblemsByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]response.Problem, error) {
	problems, err := s.problemRepo.GetProblemsByIds(ctx, ids, p)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetProblemsByIds : %v\n", err)
		return nil, err
	}

	responses := make([]response.Problem, len(problems))
	for i, problem := range problems {
		response, err := mapper.ProblemToDTO(&problem)
		if err != nil {
			log.Printf("[ERROR] Problem Service - GetProblemsByIds : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *problemService) UpdateProblemById(ctx context.Context, id uuid.UUID, body *request.ProblemUpdate) (*response.Problem, error) {
	problem, err := s.problemRepo.UpdateProblemById(ctx, id, &entity.Problem{
		Title:        body.Title,
		DifficultyID: uuid.Must(uuid.Parse(body.DifficultyID)),
		Rating:       &body.Rating,
	})
	if err != nil {
		log.Printf("[ERROR] Problem Service - UpdateProblemById : %v\n", err)
		return nil, err
	}

	response, err := mapper.ProblemToDTO(problem)
	if err != nil {
		log.Printf("[ERROR] Problem Service - UpdateProblemById : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *problemService) DeleteProblemById(ctx context.Context, id uuid.UUID) error {
	if err := s.problemRepo.DeleteProblemById(ctx, id); err != nil {
		log.Printf("[ERROR] Problem Service - DeleteProblemById : %v\n", err)
		return err
	}

	return nil
}
