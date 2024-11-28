package controllers_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/frinfo702/MyApi/controllers"
	"github.com/frinfo702/MyApi/services"
	_ "github.com/go-sql-driver/mysql"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("failed to connect db")
		os.Exit(1)
	}

	service := services.NewAppService(db)
	aCon = controllers.NewArticleController(service)

	m.Run()
}
