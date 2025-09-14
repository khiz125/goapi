package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/khiz125/goapi/domain"
	"github.com/khiz125/goapi/handlers"
)

func main() {

	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	//dbConn := fmt.Sprintf("%s:%s@tcp(db-for-go:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	const sqlStr = `
  select *
  from articles;
  `

	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	articleArray := make([]domain.Article, 0)
	for rows.Next() {
		var article domain.Article
		var createdTime sql.NullTime
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Contents,
			&article.UserName,
			&article.NiceNum,
			&createdTime,
		)

		if err != nil {
			fmt.Println(err)
		} else {
			articleArray = append(articleArray, article)
		}
	}

	fmt.Printf("%+v\n", articleArray)

	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("Server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
