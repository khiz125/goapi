package domain

type Article struct {
  ID          uint      `json:"id" gorm:"primaryKey"`
  Title       string    `json:"title"`
  Contents    string    `json:"contents"`
  UserName    string    `json:"user_name"`
  CommentList []Comment `json:"comments"`
  NiceNum     int       `json:"nice"`
  CreatedAt   time.Time `json:"created_at"`
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

type Commnet struct {
  CommentID int       `json:"comment_id"`
  ArticleID int       `json:"article_id"`
  Message   string    `json:"message"`
  CreatedAt time.Time `json:"created_at"`
}
