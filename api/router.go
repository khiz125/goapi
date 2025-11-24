package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khiz125/goapi/api/middlewares"
	"github.com/khiz125/goapi/controllers"
	"github.com/khiz125/goapi/services"
)

func NewRouter(db *sql.DB) *mux.Router {
	service := services.NewAppService(db)
	articleController := controllers.NewArticleController(service)
	commentController := controllers.NewCommentController(service)
	r := mux.NewRouter()

	r.HandleFunc("/hello", articleController.HelloHandler).Methods(http.MethodGet)

	r.HandleFunc("/article", articleController.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", articleController.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", articleController.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", articleController.PostNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", commentController.PostCommentHandler).Methods(http.MethodPost)

	r.Use(middlewares.LoggingMiddleware)

	return r

}
