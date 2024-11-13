package services

import (
	"log"

	"github.com/frinfo702/MyApi/models"
	"github.com/frinfo702/MyApi/repositories"
)

// ハンドラ層が Article 構造体関連で呼び出したい処理

// 指定IDの記事をデータベースから取得してくる
func (s *MyappService) GetArticleService(articleID int) (models.Article, error) {
	// 1. repositories層の関数SelectArticleDetailで記事の詳細を取得
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	// 2. repositorries層の関数SelectCommentListでコメント一覧を取得
	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	// 3. 2で得たコメント一覧を1で得たArticle構造体に紐づける
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// POST /article
func (s *MyappService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		log.Printf("failed to exec insert article query: %v", err)
		return newArticle, err
	}
	return newArticle, nil
}

// GET /article/listハンドラに対するサービス層
func (s *MyappService) GetArticleListService(page int) ([]models.Article, error) {
	articleArray, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		log.Printf("データが取得できませんでした: %v", err)
		return nil, err
	}

	return articleArray, nil
}

// PostNiceHandlerでの使用を想定
func (s *MyappService) PostNiceService(article models.Article) (models.Article, error) {
	articleID := article.ID // handlerが受け取った構造体からIdのみを取り出す
	if err := repositories.UpdateNiceNum(s.db, articleID); err != nil {
		log.Printf("failed to update a number of nices %v", err)
		return article, err
	}

	// return updated vertex Article
	return models.Article{
		ID:          article.ID,
		Title:       article.Title,
		Contents:    article.Contents,
		UserName:    article.UserName,
		NiceNum:     article.NiceNum + 1,
		CommentList: article.CommentList,
		CreatedAt:   article.CreatedAt,
	}, nil
}
