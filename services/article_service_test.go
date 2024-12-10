package services_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/frinfo702/MyApi/services"

	_ "github.com/go-sql-driver/mysql"
)

var aSer *services.MyAppService

func TestMain(m *testing.M) {
	// setup database
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("db connection:", err)
		os.Exit(1)
	}

	aSer = services.NewAppService(db)

	// excute individual benchmark tests
	m.Run()
}

func BenchmarkGetArticleService(b *testing.B) {
	articleID := 1

	b.ResetTimer() // start timer until here (not include pre-processing database)
	for i := 0; i < b.N; i++ {
		_, err := aSer.GetArticleService(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}
