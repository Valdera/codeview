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

type tagService struct {
	cfg     config.AppConfig
	tagRepo repository.TagRepository
}

func NewTagService(cfg config.AppConfig, tagRepo repository.TagRepository) service.TagService {
	return &tagService{
		cfg,
		tagRepo,
	}
}

func (s *tagService) CreateTag(ctx context.Context, body *request.TagCreate) (*response.Tag, error) {
	tag, err := s.tagRepo.CreateTag(ctx, &entity.Tag{
		Label: body.Label,
		Color: body.Color,
	})
	if err != nil {
		log.Printf("[ERROR] Tag Service - CreateTag : %v\n", err)
		return nil, err
	}

	response, err := mapper.TagToDTO(tag)
	if err != nil {
		log.Printf("[ERROR] Tag Service - CreateTag : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *tagService) GetTags(ctx context.Context, p *pagination.Pagination) ([]response.Tag, error) {
	tags, err := s.tagRepo.GetTags(ctx, p)
	if err != nil {
		log.Printf("[ERROR] Tag Service - GetTags : %v\n", err)
		return nil, err
	}

	responses := make([]response.Tag, len(tags))
	for i, tag := range tags {
		response, err := mapper.TagToDTO(&tag)
		if err != nil {
			log.Printf("[ERROR] Tag Service - GetTags : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *tagService) GetTagById(ctx context.Context, id uint) (*response.Tag, error) {
	tag, err := s.tagRepo.GetTagById(ctx, id)
	if err != nil {
		log.Printf("[ERROR] Tag Service - GetTagById : %v\n", err)
		return nil, err
	}

	response, err := mapper.TagToDTO(tag)
	if err != nil {
		log.Printf("[ERROR] Tag Service - GetTagById : %v\n", err)
		return nil, err
	}

	return response, nil
}

func (s *tagService) GetTagsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Tag, error) {
	tags, err := s.tagRepo.GetTagsByIds(ctx, ids, p)
	if err != nil {
		log.Printf("[ERROR] Tag Service - GetTagsByIds : %v\n", err)
		return nil, err
	}

	responses := make([]response.Tag, len(tags))
	for i, tag := range tags {
		response, err := mapper.TagToDTO(&tag)
		if err != nil {
			log.Printf("[ERROR] Tag Service - GetTagsByIds : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *tagService) UpdateTagById(ctx context.Context, id uint, body *request.TagUpdate) (*response.Tag, error) {
	tag, err := s.tagRepo.UpdateTagById(ctx, id, &entity.Tag{
		Label: body.Label,
		Color: body.Color,
	})
	if err != nil {
		log.Printf("[ERROR] Tag Service - UpdateTagById : %v\n", err)
		return nil, err
	}

	response, err := mapper.TagToDTO(tag)
	if err != nil {
		log.Printf("[ERROR] Tag Service - UpdateTagById : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *tagService) DeleteTagById(ctx context.Context, id uint) error {
	if err := s.tagRepo.DeleteTagById(ctx, id); err != nil {
		log.Printf("[ERROR] Tag Service - DeleteTagById : %v\n", err)
		return err
	}

	return nil
}
