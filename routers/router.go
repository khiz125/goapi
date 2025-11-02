package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khiz125/goapi/controllers"
)

func NewRouter(controller *controllers.AppController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/hello", controller.HelloHandler).Methods(http.MethodGet)

	r.HandleFunc("/article", controller.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", controller.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", controller.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", controller.PostNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", controller.PostCommentHandler).Methods(http.MethodPost)

	return r

}
