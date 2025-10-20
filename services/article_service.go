package services

import (
	"github.com/khiz125/goapi/domain"
	"github.com/khiz125/goapi/repositories"
)

func (s *AppService) GetArticleService(articleID int) (domain.Article, error) {

	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return domain.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return domain.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *AppService) GetArticleListService(page int) ([]domain.Article, error) {

	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return nil, err
	}

	return articleList, nil
}

func (s *AppService) PostArticleService(article domain.Article) (domain.Article, error) {

	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return domain.Article{}, err
	}

	return newArticle, nil
}

func (s *AppService) PostNiceService(article domain.Article) (domain.Article, error) {

	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		return domain.Article{}, err
	}

	return domain.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
