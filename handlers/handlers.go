package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/frinfo702/MyApi/models"
	"github.com/frinfo702/MyApi/services"
	"github.com/gorilla/mux"
)

// POST /articleのハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article // デコードされた結果を受け取る構造体

	// 受け取ったjsonを構造体にデコード
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decord json\n", http.StatusBadRequest)
	}

	// post article received

	article := reqArticle

	// 構造体を再度jsonにエンコード
	json.NewEncoder(w).Encode(article)
}

// GET /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int

	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	articles := []models.Article{models.Article1, models.Article2}

	// jsonにエンコード
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		http.Error(w, "failed to encord json\n", http.StatusBadRequest)
	}

	// logが未定義になるのを避けるための処理
	log.Println(page)
}

// GET /article/{id}のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article, err := services.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "failed to exec\n", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "failed to encord json\n", http.StatusBadRequest)
	}
}

// POST /article/niceのハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	// 受け取ったjsonをデコード
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
	}

	// 再度jsonに直して表示(機能の検証のためなので無意味)
	article := reqArticle
	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "failed to encode json\n", http.StatusBadRequest)
	}
}

// POST /commentのハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// jsonをデコード
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
	}

	// 機能検証のため再度エンコード
	comment := reqComment
	json.NewEncoder(w).Encode(comment)

}
