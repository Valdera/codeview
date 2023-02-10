package service

import (
	"codeview/internal/dto"
	"codeview/internal/entity"
	problemMapper "codeview/internal/mapper/problem"
	"codeview/utils/pagination"
	"context"
	"log"
)

func (s *problemService) CreateTag(ctx context.Context, body *dto.TagCreate) (*dto.TagResponse, error) {
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

func (s *problemService) GetTags(ctx context.Context, p *pagination.Pagination) ([]dto.TagResponse, error) {
	tags, err := s.problemRepo.GetTags(ctx, p)
	if err != nil {
		log.Printf("[ERROR] Tag Service - GetTags : %v\n", err)
		return nil, err
	}

	responses := make([]dto.TagResponse, len(tags))
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

func (s *problemService) GetTagById(ctx context.Context, id uint) (*dto.TagResponse, error) {
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

func (s *problemService) GetTagsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]dto.TagResponse, error) {
	tags, err := s.problemRepo.GetTagsByIds(ctx, ids, p)
	if err != nil {
		log.Printf("[ERROR] Tag Service - GetTagsByIds : %v\n", err)
		return nil, err
	}

	responses := make([]dto.TagResponse, len(tags))
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

func (s *problemService) UpdateTagById(ctx context.Context, id uint, body *dto.TagUpdate) (*dto.TagResponse, error) {
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
