package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/khiz125/goapi/apperrors"
	"github.com/khiz125/goapi/controllers/services"
	"github.com/khiz125/goapi/domain"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment domain.Comment

	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
    err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
	}
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "failed to internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
