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

func (s *problemService) CreateQuestion(ctx context.Context, body *request.QuestionCreate) (*response.Question, error) {
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

func (s *problemService) GetQuestions(ctx context.Context, p *pagination.Pagination) ([]response.Question, error) {
	questions, err := s.problemRepo.GetQuestions(ctx, p)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetQuestions : %v\n", err)
		return nil, err
	}

	responses := make([]response.Question, len(questions))
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

func (s *problemService) GetQuestionById(ctx context.Context, id uint) (*response.Question, error) {
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

func (s *problemService) GetQuestionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Question, error) {
	questions, err := s.problemRepo.GetQuestionsByIds(ctx, ids, p)
	if err != nil {
		log.Printf("[ERROR] Problem Service - GetQuestionsByIds : %v\n", err)
		return nil, err
	}

	responses := make([]response.Question, len(questions))
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

func (s *problemService) UpdateQuestionById(ctx context.Context, id uint, body *request.QuestionUpdate) (*response.Question, error) {
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
