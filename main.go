package main

import (
	"github.com/ekusuy/go_api_blog/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server_start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
