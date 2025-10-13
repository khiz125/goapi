package testdata

import "github.com/khiz125/goapi/domain"

var ArticleTestData = []domain.Article{

	domain.Article{
		ID:       1,
		Title:    "first post",
		Contents: "This is a test blog post",
		UserName: "test user",
		NiceNum:  4,
	},

	domain.Article{
		ID:       2,
		Title:    "2nd post",
		Contents: "2nd blog post",
		UserName: "test user",
		NiceNum:  1,
	},
}

var CommentTestData = []models.Comment{
	domain.Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "1st comment yeah",
	},
	domain.Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "welcome",
	},
}
