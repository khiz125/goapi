package repositories_test

import (
	"testing"

	"github.com/khiz125/goapi/domain"
	"github.com/khiz125/goapi/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {

	tests := []struct {
		testTitle string
		expected  domain.Article
	}{
		{
			testTitle: "subtest1",
			expected: domain.Article{
				ID:       1,
				Title:    "first post",
				Contents: "This is a test blog post",
				UserName: "test user",
				NiceNum:  1,
			},
		},
		{
			testTitle: "subtest2",
			expected: domain.Article{
				ID:       2,
				Title:    "2nd post",
				Contents: "2nd blog post",
				UserName: "test user",
				NiceNum:  1,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("get %d but actual data should be %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("get %s but actual data should be %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("get %s but actual data should be %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("get %s but actual data should be %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("get %d but actual data should be %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

func TestSelectArticleList(t *testing.T) {

	expectedNum := 4

	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("got %d but got should %d articles\n", num, expectedNum)
	}
}

func TestInsertArticle(t *testing.T) {
	article := domain.Article{
		Title:    "insertTest",
		Contents: "test content",
		UserName: "test",
  }

  expectedArticleNum := 5
  newArticle, err := repositories.InsertArticle(testDB, article)
  if err != nil {
    t.Error(err)
  }
  if newArticle.ID != expectedArticleNum {
    t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum, newArticle.ID)
  }

  t.Cleanup(func() {
    const sqlStr = `
    delete from articles where title = ? and contents = ? and username = ?;
    `

    testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
  })
}
