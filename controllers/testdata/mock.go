package testdata

import "github.com/khiz125/goapi/domain"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) PostArticleService(article domain.Article) (domain.Article, error) {
	return articleTestData[1], nil
}

func (s *serviceMock) GetArticleListService(page int) ([]domain.Article, error) {
	return articleTestData, nil
}

func (s *serviceMock) GetArticleService(articleID int) (domain.Article, error) {
	return articleTestData[0], nil
}

func (s *serviceMock) PostNiceService(article domain.Article) (domain.Article, error) {
	return articleTestData[0], nil
}

func (s *serviceMock) PostCommentService(comment domain.Comment) (domain.Comment, error) {
	return commentTestData[0], nil
}
