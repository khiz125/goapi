package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/khiz125/goapi/domain"
	"github.com/khiz125/goapi/services"
)

type AppController struct {
	service *services.AppService
}

func NewAppController(s *services.AppService) *AppController {
	return &AppController{service: s}
}

func (c *AppController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello go world!\n")
}

func (c *AppController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle domain.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "failed internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *AppController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "failed internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articleList)
}

func (c *AppController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])

	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "failed to internal exex\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)

}

func (c *AppController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle domain.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "failed to internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

func (c *AppController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment domain.Comment

	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
	}
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "failed to internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
