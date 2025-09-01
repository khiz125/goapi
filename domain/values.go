
package domain

import (
  "errors"
)

type ArticleNameVal string
type ArticleName struct {
	val ArticleNameVal
}
func NewArticleName(val ArticleNameVal) (ArticleName, error) {
	if val == "" {
		return ArticleName{}, errors.New("invalid Article Name")
	}
	return ArticleName{val: val}, nil
}

type ArticleBodyVal string
type ArticleBody struct {
	val ArticleBodyVal
}
func NewArticleBody(val ArticleBodyVal) (ArticleBody, error) {
	if val == "" {
		return ArticleBody{}, errors.New("invalid Article Body")
	}
	return ArticleBody{val: val}, nil
}

