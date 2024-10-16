package repositories

import (
	"database/sql"
	"fmt"

	"github.com/frinfo702/MyApi/models"
)

const articleNumPerPage = 5

// articleテーブルを操作する処理
// 新規投稿をデータベースにinsertする関数
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values
	(?, ?, ?, 0, now());
	`
	// 挿入したのはこんなデータだよーと教えるための構造体
	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err)
		return models.Article{}, err
	} else {
		insertID, _ := result.LastInsertId()
		newArticle.ID = int(insertID)
		return newArticle, nil

	}

}

// 変数pageで指定されたデータを取得する関数
func selectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlGetArticlesInPage = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;
	`

	rows, err := db.Query(sqlGetArticlesInPage, articleNumPerPage, (page-1)*articleNumPerPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 取得したデータを入れるための構造体配列
	articleArray := make([]models.Article, articleNumPerPage)
	// 取得したデータを構造体に格納していく
	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		articleArray = append(articleArray, article)
	}
	return articleArray, nil

}

// 指定されたIDの投稿を取得する関数
func selectArticleDetail(db *sql.DB, articleId int) (models.Article, error) {
	// クエリを定義
	const sqlGetArticle = `
		select *
		from articles
		where article_id = ?
	`

	// 取得したデータを取り扱うためのrows
	row := db.QueryRow(sqlGetArticle, articleId)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	// データを構造体に格納
	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(article.ID, article.Contents, article.UserName, article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

// 指定されたIDの投稿のいいね数を+1updateする関数
func updateNiceNum(db *sql.DB, artilceId int) error {
	// 取得と更新があるのでトランザクションを利用
	// トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// データ取得
	const sqlGetNiceNum = `
		select nice 
		from artilces
		where article_id = ?
	`
	// 取得したデータを取り扱う
	row := db.QueryRow(sqlGetNiceNum, artilceId)

	if err = row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var niceNum int
	if err := row.Scan(&niceNum); err != nil {
		tx.Rollback()
		return err
	}

	// データを更新
	const sqlUpdateNiceNum = `
		update articles set nice = ? where article_id = ?
	`
	if _, err = tx.Exec(sqlGetNiceNum, niceNum+1, artilceId); err != nil {
		tx.Rollback()
		return err
	}

	// トランザクション終了
	if err := tx.Commit(); err != nil {
		fmt.Println("failed to finish transaction")
	}

	return nil

}
