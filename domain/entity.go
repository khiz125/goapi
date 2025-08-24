package domain

type Article struct {
  ID          uint    `json:"id" gorm:"primaryKey"`
  Name        string  `json:"name"`
  Body        string  `json:"body"`
  Nice        int     `json:"nice"`
}

func NewArticle(ID ArticleID, name ArticleName, body ArticleBody) (Article, err) {
  if err != nil {
    return Article{}, errors.New(err)
  }
  return Article{
    name: name,
    body: body,
    nice: 0
  }
}
