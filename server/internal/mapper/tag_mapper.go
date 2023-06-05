package mapper

import (
	"codeview/internal/dto/response"
	"codeview/internal/entity"
)

func TagToDTO(tag *entity.Tag) (*response.Tag, error) {
	res := response.Tag{
		ID:        tag.ID,
		Label:     tag.Label,
		Color:     tag.Color,
		TagType:   string(tag.TagType),
		CreatedAt: tag.CreatedAt,
		UpdatedAt: tag.UpdatedAt,
	}

	return &res, nil
}
