package controllers_test

import (
	"github.com/ekusuy/go_api_blog/controllers"
	"github.com/ekusuy/go_api_blog/controllers/testdata"
	"testing"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
