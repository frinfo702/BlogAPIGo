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

// Handler for POST /article
func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article // Struct to receive decoded result

	// Decode received JSON into struct
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		log.Println(err)
		apperrors.ErrorHandler(w, req, err)
	}

	// Post article received
	insertedArticle, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	// Encode the inserted Article back to JSON
	json.NewEncoder(w).Encode(insertedArticle)
}

// Handler for GET /article/list
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

	// Encode to JSON
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		apperrors.ErrorHandler(w, req, err)
	}

}

// Handler for GET /article/{id}
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

// Handler for POST /article/nice
func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	// Decode received JSON
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "failed to decode json")
		apperrors.ErrorHandler(w, req, err)
	}

	// Update Nice
	updatedArticle, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	// Encode back to JSON and return response
	if err := json.NewEncoder(w).Encode(updatedArticle); err != nil {
		apperrors.ErrorHandler(w, req, err)
	}
}
