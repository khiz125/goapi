package domain

import (
	"time"
)

type Article struct {
	ID          int       `json:"article_id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	CommentList []Comment `json:"comments"`
	NiceNum     int       `json:"nice"`
	CreatedAt   time.Time `json:"created_at"`
}

type Comment struct {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
