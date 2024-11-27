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
		apperrors.ErrorHandler(w, req, err)
	}

	// コメントを投稿
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	// 再度jsonにしてレスポンスに書き込む
	json.NewEncoder(w).Encode(comment)

}
