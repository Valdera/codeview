package service

import (
	"codeview/internal/dto"
	"codeview/internal/entity"
	problemMapper "codeview/internal/mapper/problem"
	"codeview/utils/pagination"
	"context"
	"log"
)

func (s *problemService) CreateDifficulty(ctx context.Context, body *dto.DifficultyCreate) (*dto.DifficultyResponse, error) {
	difficulty, err := s.problemRepo.CreateDifficulty(ctx, &entity.Difficulty{
		Label: body.Label,
		Color: body.Color,
	})
	if err != nil {
		log.Printf("[ERROR] Problem Service - CreateDifficulty : %v\n", err)
		return nil, err
	}

	response, err := problemMapper.DifficultyToDTO(difficulty)
	if err != nil {
		log.Printf("[ERROR] Problem Service - CreateDifficulty : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *problemService) GetDifficulties(ctx context.Context, p *pagination.Pagination) ([]dto.DifficultyResponse, error) {
	difficulties, err := s.problemRepo.GetDifficulties(ctx, p)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetDifficulties : %v\n", err)
		return nil, err
	}

	responses := make([]dto.DifficultyResponse, len(difficulties))
	for i, difficulty := range difficulties {
		response, err := problemMapper.DifficultyToDTO(&difficulty)
		if err != nil {
			log.Printf("[ERROR] Problem Service - GetDifficulties : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *problemService) GetDifficultyById(ctx context.Context, id uint) (*dto.DifficultyResponse, error) {
	difficulty, err := s.problemRepo.GetDifficultyById(ctx, id)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetDifficultyById : %v\n", err)
		return nil, err
	}

	response, err := problemMapper.DifficultyToDTO(difficulty)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetDifficultyById : %v\n", err)
		return nil, err
	}

	return response, nil
}

func (s *problemService) GetDifficultiesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]dto.DifficultyResponse, error) {
	difficulties, err := s.problemRepo.GetDifficultiesByIds(ctx, ids, p)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetDifficultiesByIds : %v\n", err)
		return nil, err
	}

	responses := make([]dto.DifficultyResponse, len(difficulties))
	for i, difficulty := range difficulties {
		response, err := problemMapper.DifficultyToDTO(&difficulty)
		if err != nil {
			log.Printf("[ERROR] Problem Service - GetDifficultiesByIds : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *problemService) UpdateDifficultyById(ctx context.Context, id uint, body *dto.DifficultyUpdate) (*dto.DifficultyResponse, error) {
	difficulty, err := s.problemRepo.UpdateDifficultyById(ctx, id, &entity.Difficulty{
		Label: body.Label,
		Color: body.Color,
	})
	if err != nil {
		log.Printf("[ERROR] Problem Service - UpdateDifficultyById : %v\n", err)
		return nil, err
	}

	response, err := problemMapper.DifficultyToDTO(difficulty)
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
