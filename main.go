package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/frinfo702/MyApi/controllers"

	"github.com/frinfo702/MyApi/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

	ser := services.NewAppService(db)
	con := controllers.NewMyAppController(ser)

	r := routers.NewRouter(con)


	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
