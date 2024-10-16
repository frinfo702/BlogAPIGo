package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*
		r := mux.NewRouter()

		// ハンドラの登録
		r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
		r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
		r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
		r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
		r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
		r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

		log.Println("server start at port 8080")
		log.Fatal(http.ListenAndServe(":8080", r))
	*/

	// データベースユーザーの設定
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	// データベースの起動
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// 起動のテスト
	if err := db.Ping(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("successfully connect to DB")
	}
}
