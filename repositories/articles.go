package repositories

import (
	"database/sql"

	"github.com/khiz125/goapi/domain"
)

const (
	articleNumPerPage = 5
)

// create new article
func InsertArticle(db *sql.DB, article domain.Article) (domain.Article, error) {
	const sqlStr = `
  insert into articles (
    title, contents, username, nice, created_at
  ) values (
    ?, ?, ?, 0, now()
  );
  `

	var newArticle domain.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return domain.Article{}, err
	}

	id, _ := result.LastInsertId()

	newArticle.ID = int(id)

	return newArticle, nil
}

// get articles
func SelectArticleList(db *sql.DB, page int) ([]domain.Article, error) {
	const sqlStr = `
  select select article_id, title, contents, username, nice 
  from articles limit ? offset ?;
  `

	rows, err := db.Query(sqlStr, articleNumPerPage, ((page - 1) * articleNumPerPage))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	articleArray := make([]domain.Article, 0)
	for rows.Next() {
		var article domain.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &article.CreatedAt)

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

// get article with id
func SelectArticleDetail(db *sql.DB, articleID int) (domain.Article, error) {
	const sqlStr = `
  select article_id, title, contents, username, nice, created_at 
  from articles where article_id = ?;
  `

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return domain.Article{}, err
	}

	var article domain.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return domain.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

// update nice number
func UpdateNiceNum(db *sql.DB, articleID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	const sqlStr = `
  select nice from articles where article_id = ?;
  `

	row := tx.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		tx.Rollback()
		return err
	}

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?;`
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
