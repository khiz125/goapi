package repositories_test

import (
	"testing"

	"github.com/khiz125/goapi/domain"
	"github.com/khiz125/goapi/repositories"
	"github.com/khiz125/goapi/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {

	tests := []struct {
		testTitle string
		expected  domain.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
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
				t.Errorf("get %s but actual  data should be %s\n", got.Contents, test.expected.Contents)
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
		t.Errorf("got %d but expected should %d articles\n", num, expectedNum)
	}
}

func TestInsertArticle(t *testing.T) {
	article := domain.Article{
		Title:    "insertTest",
		Contents: "test content",
		UserName: "test",
	}

	expectedArticleTitle := "insertTest"
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.Title != expectedArticleTitle {
		t.Errorf("new article id is expected %s but got %s\n", expectedArticleTitle, newArticle.Title)
	}

	t.Cleanup(func() {
		const sqlStr = `
    delete from articles where title = ? and contents = ? and username = ?;
    `

		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

func TestUpdateNiceNum(t *testing.T) {

	articleID := 3
	before, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal("failed to get before data")
	}

	err = repositories.UpdateNiceNum(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	after, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal("failed to get after data")
	}

	if after.NiceNum != (before.NiceNum + 1) {
		t.Error("failed to update nice num")
	}
}
