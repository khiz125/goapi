package services

import (
	"github.com/khiz125/goapi/domain"
	"github.com/khiz125/goapi/repositories"
)

func GetArticleService(articleID int) (domain.Article, error) {

	db, err := connectDB()
	if err != nil {
		return domain.Article{}, err
	}
	defer db.Close()

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return domain.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return domain.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func GetArticleListService(page int) ([]domain.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	artcleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}

	return artcleList, nil
}

func PostArticleService(article domain.Article) (domain.Article, error) {
	db, err := connectDB()
	if err != nil {
		return domain.Article{}, err
	}

	defer db.Close()

	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return domain.Article{}, err
	}

	return newArticle, nil
}

func PostNiceService(article domain.Article) (domain.Article, error) {
	db, err := connectDB()
	if err != nil {
		return domain.Article{}, err
	}
	defer db.Close()

	err = repositories.UpdateNiceNum(db, article.ID)
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
