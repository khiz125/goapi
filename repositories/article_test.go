package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/khiz125/goapi/domain"
	"github.com/khiz125/goapi/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	expected := domain.Article{
		ID:       1,
		Title:    "first post",
		Contents: "This is a test blog post",
		UserName: "test user",
		NiceNum:  1,
	}

	got, err := repositories.SelectArticleDetail(db, expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if got.ID != expected.ID {
		t.Errorf("get %d but actual data should be %d\n", got.ID, expected.ID)
	}
	if got.Title != expected.Title {
		t.Errorf("get %s but actual data should be %s\n", got.Title, expected.Title)
	}
	if got.Contents != expected.Contents {
		t.Errorf("get %s but actual data should be %s\n", got.Contents, expected.Contents)
	}
	if got.UserName != expected.UserName {
		t.Errorf("get %s but actual data should be %s\n", got.UserName, expected.UserName)
	}
	if got.NiceNum != expected.NiceNum {
		t.Errorf("get %d but actual data should be %d\n", got.NiceNum, expected.NiceNum)
	}

}
