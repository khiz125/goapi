package main

import (
  "log"
  "net/http"

  "github.com/gorilla/mux"
  "github.com/khiz125/goapi/handlers"
)

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
  r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodGet)
  r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
  r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
  r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodGet)
  r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodGet)

  log.Println("Server start at port 8080")
  log.Fatal(http.ListenAndServe(":8080", r))
}

