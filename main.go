package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/khiz125/goapi/controllers"
	"github.com/khiz125/goapi/routers"
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
	r := routers.NewRouter(controller)

	log.Println("Server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
