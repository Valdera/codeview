package service

import (
	"codeview/internal/dto"
	"codeview/internal/entity"
	problemMapper "codeview/internal/mapper/problem"
	"codeview/utils/pagination"
	"context"
	"log"
)

func (s *problemService) CreateSource(ctx context.Context, body *dto.SourceCreate) (*dto.SourceResponse, error) {
	source, err := s.problemRepo.CreateSource(ctx, &entity.Source{
		Label: body.Label,
		Color: body.Color,
	})
	if err != nil {
		log.Printf("[ERROR] Problem Service - CreateSource : %v\n", err)
		return nil, err
	}

	response, err := problemMapper.SourceToDTO(source)
	if err != nil {
		log.Printf("[ERROR] Problem Service - CreateSource : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *problemService) GetSources(ctx context.Context, p *pagination.Pagination) ([]dto.SourceResponse, error) {
	sources, err := s.problemRepo.GetSources(ctx, p)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetSources : %v\n", err)
		return nil, err
	}

	responses := make([]dto.SourceResponse, len(sources))
	for i, source := range sources {
		response, err := problemMapper.SourceToDTO(&source)
		if err != nil {
			log.Printf("[ERROR] Problem Service - GetSources : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *problemService) GetSourceById(ctx context.Context, id uint) (*dto.SourceResponse, error) {
	source, err := s.problemRepo.GetSourceById(ctx, id)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetSourceById : %v\n", err)
		return nil, err
	}

	response, err := problemMapper.SourceToDTO(source)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetSourceById : %v\n", err)
		return nil, err
	}

	return response, nil
}

func (s *problemService) GetSourcesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]dto.SourceResponse, error) {
	sources, err := s.problemRepo.GetSourcesByIds(ctx, ids, p)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetSourcesByIds : %v\n", err)
		return nil, err
	}

	responses := make([]dto.SourceResponse, len(sources))
	for i, source := range sources {
		response, err := problemMapper.SourceToDTO(&source)
		if err != nil {
			log.Printf("[ERROR] Problem Service - GetSourcesByIds : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *problemService) UpdateSourceById(ctx context.Context, id uint, body *dto.SourceUpdate) (*dto.SourceResponse, error) {
	source, err := s.problemRepo.UpdateSourceById(ctx, id, &entity.Source{
		Label: body.Label,
		Color: body.Color,
	})
	if err != nil {
		log.Printf("[ERROR] Problem Service - UpdateSourceById : %v\n", err)
		return nil, err
	}

	response, err := problemMapper.SourceToDTO(source)
	if err != nil {
		log.Printf("[ERROR] Problem Service - UpdateSourceById : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *problemService) DeleteSourceById(ctx context.Context, id uint) error {
	if err := s.problemRepo.DeleteSourceById(ctx, id); err != nil {
		log.Printf("[ERROR] Problem Service - DeleteSourceById : %v\n", err)
		return err
	}

	return nil
}
