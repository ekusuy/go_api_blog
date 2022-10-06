package main

import (
	"database/sql"
	"fmt"
	"github.com/ekusuy/go_api_blog/controllers"
	"github.com/ekusuy/go_api_blog/router"
	"github.com/ekusuy/go_api_blog/services"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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
		log.Println("fail to connect DB")
		return
	}

	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)

	r := router.NewRouter(con)
	log.Println("server_start")
	log.Fatal(http.ListenAndServe(":8080", r))
}
