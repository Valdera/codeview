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
	"errors"
	"log"

	"github.com/google/uuid"
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
		Label:   body.Label,
		Color:   body.Color,
		TagType: entity.TagType(body.TagType),
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

func (s *tagService) GetTagById(ctx context.Context, id uuid.UUID) (*response.Tag, error) {
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

func (s *tagService) GetTagsByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]response.Tag, error) {
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

func (s *tagService) UpdateTagById(ctx context.Context, id uuid.UUID, body *request.TagUpdate) (*response.Tag, error) {
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

func (s *tagService) DeleteTagById(ctx context.Context, id uuid.UUID) error {
	if err := s.tagRepo.DeleteTagById(ctx, id); err != nil {
		log.Printf("[ERROR] Tag Service - DeleteTagById : %v\n", err)
		return err
	}

	return nil
}

func (s *tagService) GetTagsByType(ctx context.Context, tagType string, p *pagination.Pagination) ([]response.Tag, error) {
	switch entity.TagType(tagType) {
	case entity.TagTypeCollection, entity.TagTypeNote, entity.TagTypeProblem:
		{
			tags, err := s.tagRepo.GetTagsByType(ctx, entity.TagType(tagType), p)
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
	default:
		err := errors.New("invalid tag type")
		log.Printf("[ERROR] Tag Service - GetTags : %v\n", err)
		return nil, err
	}

}
