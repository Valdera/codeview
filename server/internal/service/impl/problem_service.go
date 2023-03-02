package service

import (
	"codeview/config"
	"codeview/internal/dto/request"
	"codeview/internal/dto/response"
	"codeview/internal/entity"
	"codeview/internal/mapper"
	"codeview/internal/repository"
	"codeview/internal/service"
	"codeview/internal/util/pagination"
	"context"
	"log"
)

type problemService struct {
	config      config.AppConfig
	problemRepo repository.ProblemRepository
}

func NewProblemService(config config.AppConfig, problemRepo repository.ProblemRepository) service.ProblemService {
	return &problemService{
		config,
		problemRepo,
	}
}

func (s *problemService) CreateProblem(ctx context.Context, body *request.ProblemCreate) (*response.Problem, error) {
	problem, err := s.problemRepo.CreateProblem(ctx, &entity.Problem{
		Title:        body.Title,
		DifficultyID: uint(body.DifficultyID),
		Rating:       &body.Rating,
	})
	if err != nil {
		log.Printf("[ERROR] Problem Service - CreateProblem : %v\n", err)
		return nil, err
	}

	response, err := mapper.ProblemToDTO(problem)
	if err != nil {
		log.Printf("[ERROR] Problem Service - CreateProblem : %v\n", err)
		return nil, err
	}

	return response, err
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

func (s *problemService) GetProblemById(ctx context.Context, id uint) (*response.Problem, error) {
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

func (s *problemService) GetProblemsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Problem, error) {
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

func (s *problemService) UpdateProblemById(ctx context.Context, id uint, body *request.ProblemUpdate) (*response.Problem, error) {
	problem, err := s.problemRepo.UpdateProblemById(ctx, id, &entity.Problem{
		Title:        body.Title,
		DifficultyID: uint(body.Difficulty),
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

func (s *problemService) DeleteProblemById(ctx context.Context, id uint) error {
	if err := s.problemRepo.DeleteProblemById(ctx, id); err != nil {
		log.Printf("[ERROR] Problem Service - DeleteProblemById : %v\n", err)
		return err
	}

	return nil
}
