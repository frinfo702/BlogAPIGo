package models

import "time"

/*
まだデータベースが実装できていないので、この中には
データベースから取り出されたデータが入っているという想定
*/

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "This is the first test comment",
		CreatedAt: time.Now(),
	}
	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "This is the second test comment",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		ID:          1,
		Title:       "First Article",
		Contents:    "This is the first test article",
		UserName:    "ken",
		NiceNum:     1,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}
	Article2 = Article{
		ID:        2,
		Title:     "Second Article",
		Contents:  "This is the second test article",
		UserName:  "ken",
		NiceNum:   2,
		CreatedAt: time.Now(),
	}
)
