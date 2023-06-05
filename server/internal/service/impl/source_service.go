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

	"github.com/google/uuid"
)

type sourceService struct {
	cfg        config.AppConfig
	sourceRepo repository.SourceRepository
}

func NewSourceService(cfg config.AppConfig, sourceRepo repository.SourceRepository) service.SourceService {
	return &sourceService{
		cfg,
		sourceRepo,
	}
}

func (s *sourceService) CreateSource(ctx context.Context, body *request.SourceCreate) (*response.Source, error) {
	source, err := s.sourceRepo.CreateSource(ctx, &entity.Source{
		Label: body.Label,
		Color: body.Color,
	})
	if err != nil {
		log.Printf("[ERROR] Source Service - CreateSource : %v\n", err)
		return nil, err
	}

	response, err := mapper.SourceToDTO(source)
	if err != nil {
		log.Printf("[ERROR] Source Service - CreateSource : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *sourceService) GetSources(ctx context.Context, p *pagination.Pagination) ([]response.Source, error) {
	sources, err := s.sourceRepo.GetSources(ctx, p)
	if err != nil {
		log.Printf("[ERROR] Source Service - GetSources : %v\n", err)
		return nil, err
	}

	responses := make([]response.Source, len(sources))
	for i, source := range sources {
		response, err := mapper.SourceToDTO(&source)
		if err != nil {
			log.Printf("[ERROR] Source Service - GetSources : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *sourceService) GetSourceById(ctx context.Context, id uuid.UUID) (*response.Source, error) {
	source, err := s.sourceRepo.GetSourceById(ctx, id)
	if err != nil {
		log.Printf("[ERROR] Source Service - GetSourceById : %v\n", err)
		return nil, err
	}

	response, err := mapper.SourceToDTO(source)
	if err != nil {
		log.Printf("[ERROR] Source Service - GetSourceById : %v\n", err)
		return nil, err
	}

	return response, nil
}

func (s *sourceService) GetSourcesByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]response.Source, error) {
	sources, err := s.sourceRepo.GetSourcesByIds(ctx, ids, p)
	if err != nil {
		log.Printf("[ERROR] Source Service - GetSourcesByIds : %v\n", err)
		return nil, err
	}

	responses := make([]response.Source, len(sources))
	for i, source := range sources {
		response, err := mapper.SourceToDTO(&source)
		if err != nil {
			log.Printf("[ERROR] Source Service - GetSourcesByIds : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *sourceService) UpdateSourceById(ctx context.Context, id uuid.UUID, body *request.SourceUpdate) (*response.Source, error) {
	source, err := s.sourceRepo.UpdateSourceById(ctx, id, &entity.Source{
		Label: body.Label,
		Color: body.Color,
	})
	if err != nil {
		log.Printf("[ERROR] Source Service - UpdateSourceById : %v\n", err)
		return nil, err
	}

	response, err := mapper.SourceToDTO(source)
	if err != nil {
		log.Printf("[ERROR] Source Service - UpdateSourceById : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *sourceService) DeleteSourceById(ctx context.Context, id uuid.UUID) error {
	if err := s.sourceRepo.DeleteSourceById(ctx, id); err != nil {
		log.Printf("[ERROR] Source Service - DeleteSourceById : %v\n", err)
		return err
	}

	return nil
}
