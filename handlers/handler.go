package handlers

import (
  "fmt"
  "io"
  "net/http"
  "strconv"
  "encoding/json"
  "errors"

  "github.com/gorilla/mux"
  
  "github.com/khiz125/goapi/mock"
  "github.com/khiz125/goapi/domain"
)


func HelloHandler(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "Hello go world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {

  length, err := strconv.Atoi(req.Header.Get("Content-Length"))
  if err != nil {
    http.Error(w, "cannot get content length\n", http.StatusBadRequest)
    return
  }

  reqBodybuffer := make([]byte, length)

  if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
    http.Error(w, "failed to get request body\n", http.StatusBadRequest)
    return
  }

  defer req.Body.Close()

  var reqArticle domain.Article

  if err := json.Unmarshal(reqBodybuffer, &reqArticle); err != nil {
    http.Error(w, "failed to decode json\n", http.StatusBadRequest)
    return
  }

  article := reqArticle
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
  
  w.Write(jsonData)
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
    errMsg := fmt.Sprintf("failed to encode json (articleID %d)\n", articleID)
    http.Error(w, errMsg, http.StatusInternalServerError)
    return
  }

  w.Write(jsonData)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
  article := mock.Article1
  jsonData, err := json.Marshal(article)

  if err != nil {
    http.Error(w, "failed to encode json\n", http.StatusInternalServerError)
    return
  }

  w.Write(jsonData)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
  comment := mock.Comment1
  jsonData, err := json.Marshal(comment)
  if err != nil {
    http.Error(w, "failed to encode json\n", http.StatusInternalServerError)
    return
  }
  w.Write(jsonData)
}

