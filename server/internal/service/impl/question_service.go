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

type questionService struct {
	cfg          config.AppConfig
	questionRepo repository.QuestionRepository
}

func NewQuestionService(cfg config.AppConfig, questionRepo repository.QuestionRepository) service.QuestionService {
	return &questionService{
		cfg,
		questionRepo,
	}

}

func (s *questionService) CreateQuestion(ctx context.Context, body *request.QuestionCreate) (*response.Question, error) {
	question, err := s.questionRepo.CreateQuestion(ctx, &entity.Question{
		Content: body.Content})
	if err != nil {
		log.Printf("[ERROR] Question Service - CreateQuestion : %v\n", err)
		return nil, err
	}

	response, err := mapper.QuestionToDTO(question)
	if err != nil {
		log.Printf("[ERROR] Question Service - CreateQuestion : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *questionService) GetQuestions(ctx context.Context, p *pagination.Pagination) ([]response.Question, error) {
	questions, err := s.questionRepo.GetQuestions(ctx, p)
	if err != nil {
		log.Printf("[ERROR] Question Service - GetQuestions : %v\n", err)
		return nil, err
	}

	responses := make([]response.Question, len(questions))
	for i, question := range questions {
		response, err := mapper.QuestionToDTO(&question)
		if err != nil {
			log.Printf("[ERROR] Question Service - GetQuestions : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *questionService) GetQuestionById(ctx context.Context, id uuid.UUID) (*response.Question, error) {
	question, err := s.questionRepo.GetQuestionById(ctx, id)
	if err != nil {
		log.Printf("[ERROR] Question Service - GetQuestionById : %v\n", err)
		return nil, err
	}

	response, err := mapper.QuestionToDTO(question)
	if err != nil {
		log.Printf("[ERROR] Question Service - GetQuestionById : %v\n", err)
		return nil, err
	}

	return response, nil
}

func (s *questionService) GetQuestionsByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]response.Question, error) {
	questions, err := s.questionRepo.GetQuestionsByIds(ctx, ids, p)
	if err != nil {
		log.Printf("[ERROR] Question Service - GetQuestionsByIds : %v\n", err)
		return nil, err
	}

	responses := make([]response.Question, len(questions))
	for i, question := range questions {
		response, err := mapper.QuestionToDTO(&question)
		if err != nil {
			log.Printf("[ERROR] Question Service - GetQuestionsByIds : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *questionService) UpdateQuestionById(ctx context.Context, id uuid.UUID, body *request.QuestionUpdate) (*response.Question, error) {
	question, err := s.questionRepo.UpdateQuestionById(ctx, id, &entity.Question{
		Content: body.Content,
	})
	if err != nil {
		log.Printf("[ERROR] Question Service - UpdateQuestionById : %v\n", err)
		return nil, err
	}

	response, err := mapper.QuestionToDTO(question)
	if err != nil {
		log.Printf("[ERROR] Question Service - UpdateQuestionById : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *questionService) DeleteQuestionById(ctx context.Context, id uuid.UUID) error {
	if err := s.questionRepo.DeleteQuestionById(ctx, id); err != nil {
		log.Printf("[ERROR] Question Service - DeleteQuestionById : %v\n", err)
		return err
	}

	return nil
}
