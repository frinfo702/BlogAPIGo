package services

import "github.com/frinfo702/MyApi/models"

type MyAppServicer interface {
	GetArticleService(articleID int) (models.Article, error)
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)

	PostCommentService(comment models.Comment) (models.Comment, error)
}

