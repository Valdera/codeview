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

type solutionService struct {
	cfg          config.AppConfig
	solutionRepo repository.SolutionRepository
}

func NewSolutionService(cfg config.AppConfig, solutionRepo repository.SolutionRepository) service.SolutionService {
	return &solutionService{
		cfg,
		solutionRepo,
	}
}

func (s *solutionService) CreateSolution(ctx context.Context, body *request.SolutionCreate) (*response.Solution, error) {
	solution, err := s.solutionRepo.CreateSolution(ctx, &entity.Solution{
		Content: body.Content})
	if err != nil {
		log.Printf("[ERROR] Solution Service - CreateSolution : %v\n", err)
		return nil, err
	}

	response, err := mapper.SolutionToDTO(solution)
	if err != nil {
		log.Printf("[ERROR] Solution Service - CreateSolution : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *solutionService) GetSolutions(ctx context.Context, p *pagination.Pagination) ([]response.Solution, error) {
	solutions, err := s.solutionRepo.GetSolutions(ctx, p)
	if err != nil {
		log.Printf("[ERROR] Solution Service - GetSolutions : %v\n", err)
		return nil, err
	}

	responses := make([]response.Solution, len(solutions))
	for i, solution := range solutions {
		response, err := mapper.SolutionToDTO(&solution)
		if err != nil {
			log.Printf("[ERROR] Solution Service - GetSolutions : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *solutionService) GetSolutionById(ctx context.Context, id uint) (*response.Solution, error) {
	solution, err := s.solutionRepo.GetSolutionById(ctx, id)
	if err != nil {
		log.Printf("[ERROR] Solution Service - GetSolutionById : %v\n", err)
		return nil, err
	}

	response, err := mapper.SolutionToDTO(solution)
	if err != nil {
		log.Printf("[ERROR] Solution Service - GetSolutionById : %v\n", err)
		return nil, err
	}

	return response, nil
}

func (s *solutionService) GetSolutionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Solution, error) {
	solutions, err := s.solutionRepo.GetSolutionsByIds(ctx, ids, p)
	if err != nil {
		log.Printf("[ERROR] Solution Service - GetSolutionsByIds : %v\n", err)
		return nil, err
	}

	responses := make([]response.Solution, len(solutions))
	for i, solution := range solutions {
		response, err := mapper.SolutionToDTO(&solution)
		if err != nil {
			log.Printf("[ERROR] Solution Service - GetSolutionsByIds : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *solutionService) UpdateSolutionById(ctx context.Context, id uint, body *request.SolutionUpdate) (*response.Solution, error) {
	solution, err := s.solutionRepo.UpdateSolutionById(ctx, id, &entity.Solution{
		Content: body.Content,
	})
	if err != nil {
		log.Printf("[ERROR] Solution Service - UpdateSolutionById : %v\n", err)
		return nil, err
	}

	response, err := mapper.SolutionToDTO(solution)
	if err != nil {
		log.Printf("[ERROR] Solution Service - UpdateSolutionById : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *solutionService) DeleteSolutionById(ctx context.Context, id uint) error {
	if err := s.solutionRepo.DeleteSolutionById(ctx, id); err != nil {
		log.Printf("[ERROR] Solution Service - DeleteSolutionById : %v\n", err)
		return err
	}

	return nil
}
