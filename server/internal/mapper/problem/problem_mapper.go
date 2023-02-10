package mapper

import (
	"codeview/internal/dto"
	"codeview/internal/entity"
)

func ProblemToDTO(problem *entity.Problem) (*dto.ProblemResponse, error) {
	difficultyDto, err := DifficultyToDTO(&problem.Difficulty)
	if err != nil {
		return nil, err
	}

	tagsDto := make([]dto.TagResponse, len(problem.Tags))
	for i, tag := range problem.Tags {
		tagDto, err := TagToDTO(&tag)
		if err != nil {
			return nil, err
		}
		tagsDto[i] = *tagDto
	}

	sourcesDto := make([]dto.SourceResponse, len(problem.Sources))
	for i, source := range problem.Sources {
		sourceDto, err := SourceToDTO(&source)
		if err != nil {
			return nil, err
		}
		sourcesDto[i] = *sourceDto
	}

	solutionsDto := make([]dto.SolutionResponse, len(problem.Solutions))
	for i, solution := range problem.Solutions {
		solutionDto, err := SolutionToDTO(&solution)
		if err != nil {
			return nil, err
		}
		solutionsDto[i] = *solutionDto
	}

	questionsDto := make([]dto.QuestionResponse, len(problem.Questions))
	for i, question := range problem.Questions {
		questionDto, err := QuestionToDTO(&question)
		if err != nil {
			return nil, err
		}
		questionsDto[i] = *questionDto
	}

	return &dto.ProblemResponse{
		ID:         problem.ID,
		Title:      problem.Title,
		Rating:     *problem.Rating,
		Difficulty: *difficultyDto,
		Tags:       tagsDto,
		Sources:    sourcesDto,
		Solutions:  solutionsDto,
		Questions:  questionsDto,
		CreatedAt:  problem.CreatedAt,
		UpdatedAt:  problem.UpdatedAt,
	}, nil
}

func TagToDTO(tag *entity.Tag) (*dto.TagResponse, error) {
	res := dto.TagResponse{
		ID:        tag.ID,
		Label:     tag.Label,
		Color:     tag.Color,
		CreatedAt: tag.CreatedAt,
		UpdatedAt: tag.UpdatedAt,
	}

	return &res, nil
}

func SourceToDTO(source *entity.Source) (*dto.SourceResponse, error) {
	res := dto.SourceResponse{
		ID:        source.ID,
		Label:     source.Label,
		Color:     source.Color,
		CreatedAt: source.CreatedAt,
		UpdatedAt: source.UpdatedAt,
	}

	return &res, nil
}

func DifficultyToDTO(difficulty *entity.Difficulty) (*dto.DifficultyResponse, error) {
	res := dto.DifficultyResponse{
		ID:        difficulty.ID,
		Label:     difficulty.Label,
		Color:     difficulty.Color,
		CreatedAt: difficulty.CreatedAt,
		UpdatedAt: difficulty.UpdatedAt,
	}

	return &res, nil
}

func SolutionToDTO(solution *entity.Solution) (*dto.SolutionResponse, error) {
	res := dto.SolutionResponse{
		ID:        solution.ID,
		Content:   solution.Content,
		CreatedAt: solution.CreatedAt,
		UpdatedAt: solution.UpdatedAt,
	}

	return &res, nil
}

func QuestionToDTO(question *entity.Question) (*dto.QuestionResponse, error) {
	res := dto.QuestionResponse{
		ID:        question.ID,
		Content:   question.Content,
		CreatedAt: question.CreatedAt,
		UpdatedAt: question.UpdatedAt,
	}

	return &res, nil
}
