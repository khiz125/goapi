package services

import "github.com/khiz125/goapi/domain"

type ArticleServicer interface {
	PostArticleService(article domain.Article) (domain.Article, error)
	GetArticleListService(page int) ([]domain.Article, error)
	GetArticleService(articleID int) (domain.Article, error)
	PostNiceService(article domain.Article) (domain.Article, error)
}

type CommentServicer interface {
	PostCommentService(comment domain.Comment) (domain.Comment, error)
}
