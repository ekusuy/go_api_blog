package api

import (
	"database/sql"
	"github.com/ekusuy/go_api_blog/api/middlewares"
	"github.com/ekusuy/go_api_blog/controllers"
	"github.com/ekusuy/go_api_blog/services"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	r.Use(middlewares.LoggingMiddleware)

	return r
}
