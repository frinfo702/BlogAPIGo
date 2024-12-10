package services

import (
	"database/sql"
	"errors"
	"log"

	"github.com/frinfo702/MyApi/apperrors"
	"github.com/frinfo702/MyApi/models"
	"github.com/frinfo702/MyApi/repositories"
)

// ハンドラ層が Article 構造体関連で呼び出したい処理

// 指定IDの記事をデータベースから取得してくる
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	// make vertex to deal with (article, err) at the same time
	type articleResult struct {
		article models.Article
		err     error
	}
	// chan
	articleChan := make(chan articleResult)
	defer close(articleChan)

	// 1. repositories層の関数SelectArticleDetailで記事の詳細を取得
	var article models.Article
	var commentList []models.Comment
	var articleGetError, commentListGetError error

	go func(ch chan<- articleResult, db *sql.DB, articleID int) {
		// critical section
		article, err := repositories.SelectArticleDetail(db, articleID)
		ch <- articleResult{article, err}
	}(articleChan, s.db, articleID)

	// error have two types:
	// 1. not found error(article id not found)
	// 2. database error (Internal server error)

	// make vertex to deal with (commentList, err) at the same time
	type commentResult struct {
		commentList *[]models.Comment
		err         error
	}
	// chan
	CommentChan := make(chan commentResult)
	defer close(CommentChan)

	// 2. repositorries層の関数SelectCommentListでコメント一覧を取得
	go func(ch chan<- commentResult, db *sql.DB, articleID int) {
		// critical section
		commentList, err := repositories.SelectCommentList(db, articleID)
		ch <- commentResult{&commentList, err}
	}(CommentChan, s.db, articleID)

	for i := 0; i < 2; i++ {
		select {
		case ar := <-articleChan:
			article, articleGetError = ar.article, ar.err
		case cr := <-CommentChan:
			commentList, commentListGetError = *cr.commentList, cr.err
		}
	}

	if commentListGetError != nil {
		commentListGetError = apperrors.FetchDataFailed.Wrap(commentListGetError, "failed to exec select query")
		return models.Article{}, commentListGetError
	}
	if articleGetError != nil {
		if errors.Is(articleGetError, sql.ErrNoRows) {
			articleGetError = apperrors.EmptyData.Wrap(articleGetError, "choose article is not found")
			return models.Article{}, articleGetError
		}
		articleGetError = apperrors.FetchDataFailed.Wrap(articleGetError, "failed to exec select query")
		return models.Article{}, articleGetError
	}
	// 3. 2で得たコメント一覧を1で得たArticle構造体に紐づける
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// POST /article
func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		log.Printf("failed to exec insert article query: %v", err)
		err = apperrors.InsertDataFailed.Wrap(err, "failed to insert artile")
		return newArticle, err
	}
	return newArticle, nil
}

// GET /article/listハンドラに対するサービス層
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleArray, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.FetchDataFailed.Wrap(err, "failed to exex select query")
		log.Printf("データが取得できませんでした: %v", err)
		return nil, err
	}

	// if length of array is 0, return custom error
	if len(articleArray) == 0 {
		err = apperrors.EmptyData.Wrap(ErrEmptyData, "choose article is not found")
		return nil, err
	}

	return articleArray, nil
}

// PostNiceHandlerでの使用を想定
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	articleID := article.ID // handlerが受け取った構造体からIdのみを取り出す
	// error have two types:
	// 1. not found error(article id not found)
	// 2. database error (Internal server error)
	err := repositories.UpdateNiceNum(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.FetchDataFailed.Wrap(err, "failed to update a number of nices")
			log.Printf("failed to update a number of nices %v", err)
			return article, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "failed to update a number of nices")
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
