package testdata

import "github.com/khiz125/goapi/domain"



var articleTestData = []domain.Article{
  domain.Article{
    ID: 1,
    Title: "firstPost",
    Contents: "This is my first blog",
    UserName: "testuser",
    NiceNum: 2,
    CommentList: commentTestData,
  },
  domain.Article{
ID: 2,
    Title: "2ndPost",
    Contents: "Second blog post",
    UserName: "testuser",
    NiceNum: 4,
  },
}

var commentTestData = []domain.Comment{
  domain.Comment{
    CommentID: 1,
    ArticleID: 1,
    Message: "1st comment",
  },
  domain.Comment{
    CommentID: 2,
    ArticleID: 1,
    Message: "hello",
  },
}
