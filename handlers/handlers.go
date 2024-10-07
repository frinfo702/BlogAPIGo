package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/frinfo702/MyApi/models"
	"github.com/gorilla/mux"
)

// GET /helloのハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world\n")

}

// POST /articleのハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// リクエストヘッドから長さを取得
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		errMsg := fmt.Sprintln("failed to get contents-length")
		http.Error(w, errMsg, http.StatusBadRequest)
	}
	reqBodyBuffer := make([]byte, length) // 読み取った長さ分のバイト列を作成

	// 2. Readメソッドでリクエストボディを読み出し
	if _, err := req.Body.Read(reqBodyBuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "failed to get request body", http.StatusBadRequest)
		return
	}

	// 3. ボディをcloseする
	defer req.Body.Close()

	var reqArticle models.Article // デコードされた結果を受け取る構造体
	if err := json.Unmarshal(reqBodyBuffer, &reqArticle); err != nil {
		errMsg := fmt.Sprintln("failed to decord json")
		http.Error(w, errMsg, http.StatusBadRequest)
	}

	article := reqArticle
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "failed to encord json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
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
	jsonData, err := json.Marshal(articles)
	if err != nil {
		errMsg := fmt.Sprintf("failed to encord json page:%d\n", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
	}

	w.Write(jsonData)
}

// GET /article/{id}のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprintf("failed to encord json (Article ID: %d\n)", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
	}

	w.Write(jsonData)
}

// POST /article/niceのハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprintln("failed to encord json")
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	w.Write(jsonData)
}

// POST /commentのハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprintln("failed to encord json")
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
	w.Write(jsonData)
}
