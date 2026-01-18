package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khiz125/goapi/api/middlewares"
	"github.com/khiz125/goapi/config"
	"github.com/khiz125/goapi/controllers"
	"github.com/khiz125/goapi/controllers/auth"
	"github.com/khiz125/goapi/services"
)

func NewRouter(db *sql.DB) *mux.Router {
	service := services.NewAppService(db)
	articleController := controllers.NewArticleController(service)
	commentController := controllers.NewCommentController(service)

	cfg := config.LoadGoogleOAuthConfig()
	googleController := auth.NewGoogleAuthController(cfg)
	r := mux.NewRouter()

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/google/login", googleController.Login).Methods("GET")
	authRouter.HandleFunc("/google/callback", googleController.Callback).Methods("GET")

	r.HandleFunc("/hello", articleController.HelloHandler).Methods(http.MethodGet)

	r.HandleFunc("/article", articleController.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", articleController.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", articleController.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", articleController.PostNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", commentController.PostCommentHandler).Methods(http.MethodPost)

	r.Use(middlewares.LoggingMiddleware)
	r.Use(middlewares.AuthMiddleware)
	return r

}
