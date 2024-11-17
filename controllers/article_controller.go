package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/frinfo702/MyApi/controllers/services"
	"github.com/frinfo702/MyApi/models"
	"github.com/gorilla/mux"
)

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleController(service services.ArticleServicer) *ArticleController {
	return &ArticleController{service: service}
}

// POST /articleのハンドラ
func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article // デコードされた結果を受け取る構造体

	// 受け取ったjsonを構造体にデコード
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decord json\n", http.StatusBadRequest)
	}

	// post article received
	insertedArticle, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "failed to exec post article \n", http.StatusInternalServerError)
		return
	}
	// 挿入したArticleを再度jsonにエンコード
	json.NewEncoder(w).Encode(insertedArticle)
}

// GET /article/list のハンドラ
func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
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

	articles, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "failed to exec\n", http.StatusInternalServerError)
		return
	}

	// jsonにエンコード
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		http.Error(w, "failed to encord json\n", http.StatusBadRequest)
	}

}

// GET /article/{id}のハンドラ
func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "failed to exec\n", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "failed to encord json\n", http.StatusBadRequest)
	}
}

// POST /article/niceのハンドラ
func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	// 受け取ったjsonをデコード
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
	}

	// Niceをupdate
	updatedArticle, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "failed to exec update nice num query\n", http.StatusInternalServerError)
		return
	}

	// 再度jsonにエンコードしてレスポンスを返す
	if err := json.NewEncoder(w).Encode(updatedArticle); err != nil {
		http.Error(w, "failed to encode json\n", http.StatusBadRequest)
	}
}