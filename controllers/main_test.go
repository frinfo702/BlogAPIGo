package controllers_test

import (
	"testing"

	"github.com/frinfo702/MyApi/controllers"
	"github.com/frinfo702/MyApi/controllers/testdata"
	_ "github.com/go-sql-driver/mysql"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	service := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(service)

	m.Run()
}
