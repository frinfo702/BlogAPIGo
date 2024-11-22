package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/frinfo702/MyApi/apperrors"
	"github.com/frinfo702/MyApi/controllers/services"
	"github.com/frinfo702/MyApi/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// POST /commentのハンドラ
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// jsonをデコード
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "failed to decode json")
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)

	}

	// コメントを投稿
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "failed to exec post comment query\n", http.StatusInternalServerError)
		return
	}

	// 再度jsonにしてレスポンスに書き込む
	json.NewEncoder(w).Encode(comment)

}
