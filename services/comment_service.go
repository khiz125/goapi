package services

import (
	"github.com/khiz125/goapi/domain"
	"github.com/khiz125/goapi/repositories"
)

func PostCommentService(comment domain.Comment) (domain.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return domain.Comment{}, err
	}

	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return domain.Comment{}, err
  }

  return newComment, nil
}
