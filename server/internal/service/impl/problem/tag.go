package problemservice

import (
	"codeview/internal/dto/request"
	"codeview/internal/dto/response"
	"codeview/internal/entity"
	problemMapper "codeview/internal/mapper/problem"
	"codeview/utils/pagination"
	"context"
	"log"
)

func (s *problemService) CreateTag(ctx context.Context, body *request.TagCreate) (*response.Tag, error) {
	tag, err := s.problemRepo.CreateTag(ctx, &entity.Tag{
		Label: body.Label,
		Color: body.Color,
	})
	if err != nil {
		log.Printf("[ERROR] Tag Service - CreateTag : %v\n", err)
		return nil, err
	}

	response, err := problemMapper.TagToDTO(tag)
	if err != nil {
		log.Printf("[ERROR] Tag Service - CreateTag : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *problemService) GetTags(ctx context.Context, p *pagination.Pagination) ([]response.Tag, error) {
	tags, err := s.problemRepo.GetTags(ctx, p)
	if err != nil {
		log.Printf("[ERROR] Tag Service - GetTags : %v\n", err)
		return nil, err
	}

	responses := make([]response.Tag, len(tags))
	for i, tag := range tags {
		response, err := problemMapper.TagToDTO(&tag)
		if err != nil {
			log.Printf("[ERROR] Tag Service - GetTags : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *problemService) GetTagById(ctx context.Context, id uint) (*response.Tag, error) {
	tag, err := s.problemRepo.GetTagById(ctx, id)
	if err != nil {
		log.Printf("[ERROR] Tag Service - GetTagById : %v\n", err)
		return nil, err
	}

	response, err := problemMapper.TagToDTO(tag)
	if err != nil {
		log.Printf("[ERROR] Tag Service - GetTagById : %v\n", err)
		return nil, err
	}

	return response, nil
}

func (s *problemService) GetTagsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Tag, error) {
	tags, err := s.problemRepo.GetTagsByIds(ctx, ids, p)
	if err != nil {
		log.Printf("[ERROR] Tag Service - GetTagsByIds : %v\n", err)
		return nil, err
	}

	responses := make([]response.Tag, len(tags))
	for i, tag := range tags {
		response, err := problemMapper.TagToDTO(&tag)
		if err != nil {
			log.Printf("[ERROR] Tag Service - GetTagsByIds : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *problemService) UpdateTagById(ctx context.Context, id uint, body *request.TagUpdate) (*response.Tag, error) {
	tag, err := s.problemRepo.UpdateTagById(ctx, id, &entity.Tag{
		Label: body.Label,
		Color: body.Color,
	})
	if err != nil {
		log.Printf("[ERROR] Tag Service - UpdateTagById : %v\n", err)
		return nil, err
	}

	response, err := problemMapper.TagToDTO(tag)
	if err != nil {
		log.Printf("[ERROR] Tag Service - UpdateTagById : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *problemService) DeleteTagById(ctx context.Context, id uint) error {
	if err := s.problemRepo.DeleteTagById(ctx, id); err != nil {
		log.Printf("[ERROR] Tag Service - DeleteTagById : %v\n", err)
		return err
	}

	return nil
}
