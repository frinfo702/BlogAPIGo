package services

import (
	"log"

	"github.com/frinfo702/MyApi/apperrors"
	"github.com/frinfo702/MyApi/models"
	"github.com/frinfo702/MyApi/repositories"
)

// ハンドラ層が Comment 構造体関連で呼び出したい処理
// PostCommenthandlerでの使用を想定
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {

	insertedComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to insert new comment")
		log.Printf("failed to insert new comment: %v", err)
		return comment, err
	}

	return insertedComment, nil
}
