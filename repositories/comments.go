package repositories

import (
	"database/sql"

	"github.com/frinfo702/MyApi/models"
)

// commentsテーブルを操作する関連処理を実装する

// 新規コメントを追加する関数
func insertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	// クエリの定義
	const sqlStr = `
		insert into comments (article_id, message, created_at)
		values (?, ?, now())
	`

	var newComment models.Comment
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message
	result, err := db.Exec(sqlStr, newComment.ArticleID, newComment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	newCommentID, _ := result.LastInsertId()
	newComment.CommentID = int(newCommentID)

	return newComment, nil

}

// TODO
// 指定された記事IDについているコメントを取得する関数
