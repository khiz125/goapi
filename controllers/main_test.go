package controllers_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khiz125/goapi/controllers"
	"github.com/khiz125/goapi/controllers/testdata"
)

var articleCon *controllers.ArticleController

func TestMain(m *testing.M) {
	service := testdata.NewServiceMock()
	articleCon = controllers.NewArticleController(service)

	m.Run()
}
