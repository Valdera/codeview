package problemservice

import (
	"codeview/internal/dto/request"
	"codeview/internal/dto/response"
	"codeview/internal/entity"
	"codeview/internal/mapper"
	"codeview/internal/util/pagination"
	"context"
	"log"
)

func (s *problemService) CreateSolution(ctx context.Context, body *request.SolutionCreate) (*response.Solution, error) {
	solution, err := s.problemRepo.CreateSolution(ctx, &entity.Solution{
		Content: body.Content})
	if err != nil {
		log.Printf("[ERROR] Problem Service - CreateSolution : %v\n", err)
		return nil, err
	}

	response, err := mapper.SolutionToDTO(solution)
	if err != nil {
		log.Printf("[ERROR] Problem Service - CreateSolution : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *problemService) GetSolutions(ctx context.Context, p *pagination.Pagination) ([]response.Solution, error) {
	solutions, err := s.problemRepo.GetSolutions(ctx, p)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetSolutions : %v\n", err)
		return nil, err
	}

	responses := make([]response.Solution, len(solutions))
	for i, solution := range solutions {
		response, err := mapper.SolutionToDTO(&solution)
		if err != nil {
			log.Printf("[ERROR] Problem Service - GetSolutions : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *problemService) GetSolutionById(ctx context.Context, id uint) (*response.Solution, error) {
	solution, err := s.problemRepo.GetSolutionById(ctx, id)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetSolutionById : %v\n", err)
		return nil, err
	}

	response, err := mapper.SolutionToDTO(solution)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetSolutionById : %v\n", err)
		return nil, err
	}

	return response, nil
}

func (s *problemService) GetSolutionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Solution, error) {
	solutions, err := s.problemRepo.GetSolutionsByIds(ctx, ids, p)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetSolutionsByIds : %v\n", err)
		return nil, err
	}

	responses := make([]response.Solution, len(solutions))
	for i, solution := range solutions {
		response, err := mapper.SolutionToDTO(&solution)
		if err != nil {
			log.Printf("[ERROR] Problem Service - GetSolutionsByIds : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *problemService) UpdateSolutionById(ctx context.Context, id uint, body *request.SolutionUpdate) (*response.Solution, error) {
	solution, err := s.problemRepo.UpdateSolutionById(ctx, id, &entity.Solution{
		Content: body.Content,
	})
	if err != nil {
		log.Printf("[ERROR] Problem Service - UpdateSolutionById : %v\n", err)
		return nil, err
	}

	response, err := mapper.SolutionToDTO(solution)
	if err != nil {
		log.Printf("[ERROR] Problem Service - UpdateSolutionById : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *problemService) DeleteSolutionById(ctx context.Context, id uint) error {
	if err := s.problemRepo.DeleteSolutionById(ctx, id); err != nil {
		log.Printf("[ERROR] Problem Service - DeleteSolutionById : %v\n", err)
		return err
	}

	return nil
}
