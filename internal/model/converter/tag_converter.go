package converter

import (
	"github.com/jordanmarcelino/go-article-api/internal/entity"
	"github.com/jordanmarcelino/go-article-api/internal/model"
)

func TagToResponse(tag *entity.Tag) *model.TagResponse {
	return &model.TagResponse{
		Id:   tag.ID,
		Name: tag.Name,
	}
}

func TagsToResponse(tags []entity.Tag) []model.TagResponse {
	var tagResponses []model.TagResponse

	for _, tag := range tags {
		tagResponses = append(tagResponses, *TagToResponse(&tag))
	}

	return tagResponses
}
