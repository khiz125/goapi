package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khiz125/goapi/api/middlewares"
	"github.com/khiz125/goapi/config"
	"github.com/khiz125/goapi/controllers"

	"github.com/khiz125/goapi/infrastructure/oauth"
	"github.com/khiz125/goapi/infrastructure/repositories"
	"github.com/khiz125/goapi/services"
	"github.com/khiz125/goapi/services/auth"
)

func NewRouter(db *sql.DB) *mux.Router {
	cfg := config.LoadGoogleOAuthConfig()

	// auth NewRouter
	uow := repositories.NewUnitOfWork(db)
	googleClient := oauth.NewGoogleClient(cfg)
	authService := auth.NewAuthService(uow, googleClient)

	// article router
	service := services.NewAppService(db)
	articleController := controllers.NewArticleController(service)
	commentController := controllers.NewCommentController(service)

	googleController := controllers.NewGoogleAuthController(cfg, authService)
	r := mux.NewRouter()

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/google/login", googleController.Login).Methods("GET")
	authRouter.HandleFunc("/google/callback", googleController.CallBack).Methods("GET")

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
