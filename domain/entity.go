package domain

import (
    "errors"
    "time"
)

type Article struct {
  ID          uint      `json:"id" gorm:"primaryKey"`
  Title       string    `json:"title"`
  Contents    string    `json:"contents"`
  UserName    string    `json:"user_name"`
  CommentList []Comment `json:"comments"`
  NiceNum     int       `json:"nice"`
  CreatedAt   time.Time `json:"created_at"`
}

func NewArticle(ID uint, title string, contents string) (Article, error) {
  if title == "" || contents == "" {
    return Article{}, errors.New("title and contents cannot be empty")
  }
  return Article {
    ID: ID,
    Title: title,
    Contents: contents,
    UserName: "default_user", // デフォルトのユーザー名
    CommentList: []Comment{},
    NiceNum: 0,
    CreatedAt: time.Now(),
  }, nil
}

type Comment struct {
  CommentID int       `json:"comment_id"`
  ArticleID int       `json:"article_id"`
  Message   string    `json:"message"`
  CreatedAt time.Time `json:"created_at"`
}
