package services

import (
	"github.com/khiz125/goapi/domain"
	"github.com/khiz125/goapi/repositories"
)

func (s *AppService) PostCommentService(comment domain.Comment) (domain.Comment, error) {
	
  newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return domain.Comment{}, err
	}

	return newComment, nil
}
