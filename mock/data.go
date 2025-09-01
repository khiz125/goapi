package mock

import (
  "time"
  "github.com/khiz125/goapi/domain"
)

var (
  Comment1 = domain.Comment {
    CommentID: 1,
    ArticleID: 1,
    Message: "test comment1",
    CreatedAt: time.Now(),
  }

  Comment2 = domain.Comment {
    CommentID: 2,
    ArticleID: 1,
    Message: "second comment",
    CreatedAt: time.Now(),
  }
)

var (
  Article1 = domain.Article { 
    ID: 1,
    Title: "first article",
    Contents: "This is the test article.",
    UserName: "saki",
    NiceNum: 1,
    CommentList: []domain.Comment{Comment1, Comment2},
    CreatedAt: time.Now(),
  }

  Article2 = domain.Article {
    ID: 2, 
    Title: "second article",
    Contents: "This is the test article.",
    UserName: "saki",
    NiceNum: 2,
    CommentList: []domain.Comment{Comment1, Comment2},
    CreatedAt: time.Now(),
  }
)

