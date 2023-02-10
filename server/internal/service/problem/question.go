package service

import (
	"codeview/internal/dto"
	"codeview/internal/entity"
	problemMapper "codeview/internal/mapper/problem"
	"codeview/utils/pagination"
	"context"
	"log"
)

func (s *problemService) CreateQuestion(ctx context.Context, body *dto.QuestionCreate) (*dto.QuestionResponse, error) {
	question, err := s.problemRepo.CreateQuestion(ctx, &entity.Question{
		Content: body.Content})
	if err != nil {
		log.Printf("[ERROR] Problem Service - CreateQuestion : %v\n", err)
		return nil, err
	}

	response, err := problemMapper.QuestionToDTO(question)
	if err != nil {
		log.Printf("[ERROR] Problem Service - CreateQuestion : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *problemService) GetQuestions(ctx context.Context, p *pagination.Pagination) ([]dto.QuestionResponse, error) {
	questions, err := s.problemRepo.GetQuestions(ctx, p)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetQuestions : %v\n", err)
		return nil, err
	}

	responses := make([]dto.QuestionResponse, len(questions))
	for i, question := range questions {
		response, err := problemMapper.QuestionToDTO(&question)
		if err != nil {
			log.Printf("[ERROR] Problem Service - GetQuestions : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *problemService) GetQuestionById(ctx context.Context, id uint) (*dto.QuestionResponse, error) {
	question, err := s.problemRepo.GetQuestionById(ctx, id)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetQuestionById : %v\n", err)
		return nil, err
	}

	response, err := problemMapper.QuestionToDTO(question)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetQuestionById : %v\n", err)
		return nil, err
	}

	return response, nil
}

func (s *problemService) GetQuestionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]dto.QuestionResponse, error) {
	questions, err := s.problemRepo.GetQuestionsByIds(ctx, ids, p)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetQuestionsByIds : %v\n", err)
		return nil, err
	}

	responses := make([]dto.QuestionResponse, len(questions))
	for i, question := range questions {
		response, err := problemMapper.QuestionToDTO(&question)
		if err != nil {
			log.Printf("[ERROR] Problem Service - GetQuestionsByIds : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *problemService) UpdateQuestionById(ctx context.Context, id uint, body *dto.QuestionUpdate) (*dto.QuestionResponse, error) {
	question, err := s.problemRepo.UpdateQuestionById(ctx, id, &entity.Question{
		Content: body.Content,
	})
	if err != nil {
		log.Printf("[ERROR] Problem Service - UpdateQuestionById : %v\n", err)
		return nil, err
	}

	response, err := problemMapper.QuestionToDTO(question)
	if err != nil {
		log.Printf("[ERROR] Problem Service - UpdateQuestionById : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *problemService) DeleteQuestionById(ctx context.Context, id uint) error {
	if err := s.problemRepo.DeleteQuestionById(ctx, id); err != nil {
		log.Printf("[ERROR] Problem Service - DeleteQuestionById : %v\n", err)
		return err
	}

	return nil
}
