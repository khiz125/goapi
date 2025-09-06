package handlers

import (
  "fmt"
  "io"
  "net/http"
  "strconv"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/khiz125/goapi/mock"
)


func HelloHandler(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "Hello go world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
  article := mock.Article1
  jsonData, err := json.Marshal(article)
  if err != nil {
    http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
    return
  }

  w.Write(jsonData)
}


func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
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
  
  articleList := []domain.Article{ mock.Article1, mock.Article2 }
  jsonData, err := json.Marshal(articleList)

  if err != nil {
    errMsg := fmt.Sprintf("failed to encode json (page %d)\n", page)
    http.Error(w, errMsg, http.StatusInternalServerError)
    return
  }
  
  w.Write(jsonData[page])
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
  articleID, err := strconv.Atoi(mux.Vars(req)["id"])
  if err != nil {
    http.Error(w, "Invalid query parameter", http.StatusBadRequest)
    return
  }
  article := mock.Article1
  jsonData, err := json.Marshal(article)

  if err != nil {
    errMsg := fmtSprintf("failed to encode json (articleID %d)\n", articleID)
    http.Error(w, errMsg, http.StatusInternalServerError)
    return
  }

  w.Write(jsonData)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
  article := mock.Article1
  jsonData, err := json.Marshal(article)

  if err != nil {
    http.Error(w, "failed to encode json\n", httpStatusInternalServerError)
    return
  }

  w.Write(jsonData)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
  comment = mock.Comment1
  jsonData, err := json.Marshal(comment)
  if err != nil {
    http.Error(w, "failed to encode json\n", http.StatusInternalServerError)
    return
  }
  w.Write(jsonData)
}

