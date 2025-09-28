package repositories_test

import (
	"testing"

	"github.com/khiz125/goapi/domain"
	"github.com/khiz125/goapi/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectCommentList(t *testing.T) {

	articleID := 1

	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("got articleID %d but expected should be %d comments\n", comment.ArticleID, articleID)
		}
	}
}

func TestInsertComment(t *testing.T) {
	comment := domain.Comment{
		Message:   "test comment",
		ArticleID: 1,
	}

	expecctedCommentMessage := "test comment"

	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Fatal(err)
	}
	if newComment.Message != expecctedCommentMessage {
		t.Errorf("new comment ID is expected %s but got %s\n", expecctedCommentMessage, newComment.Message)
	}

	t.Cleanup(func() {
		const sqlStr = `
    delete from comments where message = ? and article_id = ?`

		testDB.Exec(sqlStr, comment.Message, comment.ArticleID)
	})
}
