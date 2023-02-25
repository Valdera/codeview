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

func (s *problemService) CreateDifficulty(ctx context.Context, body *request.DifficultyCreate) (*response.Difficulty, error) {
	difficulty, err := s.problemRepo.CreateDifficulty(ctx, &entity.Difficulty{
		Label: body.Label,
		Color: body.Color,
	})
	if err != nil {
		log.Printf("[ERROR] Problem Service - CreateDifficulty : %v\n", err)
		return nil, err
	}

	response, err := mapper.DifficultyToDTO(difficulty)
	if err != nil {
		log.Printf("[ERROR] Problem Service - CreateDifficulty : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *problemService) GetDifficulties(ctx context.Context, p *pagination.Pagination) ([]response.Difficulty, error) {
	difficulties, err := s.problemRepo.GetDifficulties(ctx, p)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetDifficulties : %v\n", err)
		return nil, err
	}

	responses := make([]response.Difficulty, len(difficulties))
	for i, difficulty := range difficulties {
		response, err := mapper.DifficultyToDTO(&difficulty)
		if err != nil {
			log.Printf("[ERROR] Problem Service - GetDifficulties : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *problemService) GetDifficultyById(ctx context.Context, id uint) (*response.Difficulty, error) {
	difficulty, err := s.problemRepo.GetDifficultyById(ctx, id)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetDifficultyById : %v\n", err)
		return nil, err
	}

	response, err := mapper.DifficultyToDTO(difficulty)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetDifficultyById : %v\n", err)
		return nil, err
	}

	return response, nil
}

func (s *problemService) GetDifficultiesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Difficulty, error) {
	difficulties, err := s.problemRepo.GetDifficultiesByIds(ctx, ids, p)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetDifficultiesByIds : %v\n", err)
		return nil, err
	}

	responses := make([]response.Difficulty, len(difficulties))
	for i, difficulty := range difficulties {
		response, err := mapper.DifficultyToDTO(&difficulty)
		if err != nil {
			log.Printf("[ERROR] Problem Service - GetDifficultiesByIds : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *problemService) UpdateDifficultyById(ctx context.Context, id uint, body *request.DifficultyUpdate) (*response.Difficulty, error) {
	difficulty, err := s.problemRepo.UpdateDifficultyById(ctx, id, &entity.Difficulty{
		Label: body.Label,
		Color: body.Color,
	})
	if err != nil {
		log.Printf("[ERROR] Problem Service - UpdateDifficultyById : %v\n", err)
		return nil, err
	}

	response, err := mapper.DifficultyToDTO(difficulty)
	if err != nil {
		log.Printf("[ERROR] Problem Service - UpdateDifficultyById : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *problemService) DeleteDifficultyById(ctx context.Context, id uint) error {
	if err := s.problemRepo.DeleteDifficultyById(ctx, id); err != nil {
		log.Printf("[ERROR] Problem Service - DeleteDifficultyById : %v\n", err)
		return err
	}

	return nil
}
