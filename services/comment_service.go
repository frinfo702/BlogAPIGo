package services

import (
	"log"

	"github.com/frinfo702/MyApi/models"
	"github.com/frinfo702/MyApi/repositories"
)

// ハンドラ層が Comment 構造体関連で呼び出したい処理
// PostCommenthandlerでの使用を想定
func (s *MyappService) PostCommentService(comment models.Comment) (models.Comment, error) {
	insetedComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		log.Printf("failed to insert new comment: %v", err)
		return comment, err
	}

	return insetedComment, nil
}
