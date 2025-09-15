package repositories

import (
	"database/sql"

	"github.com/khiz125/goapi/domain"
)

// create new comment
func InsertComment(db *sql.DB, comment domain.Comment) (domain.Comment, error) {
	const sqlStr = `
  insert into comments (
    article_id, message, created_at
  ) values (
    ?, ?, now()
  );
  `
	var newComment domain.Comment
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return domain.Comment{}, err
	}

	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)

	return newComment, nil
}

// get comments
func SelectCommentList(db *sql.DB, articleID int) ([]domain.Comment, error) {
	const sqlStr = `
  select * from comments where article_id = ?;
  `

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	commentArray := make([]domain.Comment, 0)

	for rows.Next() {
		var comment domain.Comment
		var createdTime sql.NullTime
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
