package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/khiz125/goapi/controllers"
	"github.com/khiz125/goapi/services"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("failed to connect DB")
		return
	}

	service := services.NewAppService(db)
	controller := controllers.NewAppController(service)
	r := mux.NewRouter()

	r.HandleFunc("/hello", controller.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", controller.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", controller.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", controller.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", controller.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", controller.PostCommentHandler).Methods(http.MethodPost)

	log.Println("Server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
