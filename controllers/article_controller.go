package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/frinfo702/MyApi/apperrors"
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
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		log.Println(err)
		apperrors.ErrorHandler(w, req, err)
	}

	// post article received
	insertedArticle, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
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
			err = apperrors.BadParam.Wrap(err, "query parameter must be number")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		page = 1
	}

	articles, err := c.service.GetArticleListService(page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	// jsonにエンコード
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		apperrors.ErrorHandler(w, req, err)
	}

}

// GET /article/{id}のハンドラ
func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "query parameter must be numbers")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		apperrors.ErrorHandler(w, req, err)
	}
}

// POST /article/niceのハンドラ
func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	// 受け取ったjsonをデコード
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "failed to decode json")
		apperrors.ErrorHandler(w, req, err)
	}

	// Niceをupdate
	updatedArticle, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	// 再度jsonにエンコードしてレスポンスを返す
	if err := json.NewEncoder(w).Encode(updatedArticle); err != nil {
		apperrors.ErrorHandler(w, req, err)
	}
}
