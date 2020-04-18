package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/seunghunee/moum-server/article"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	getID := func() (article.ID, error) {
		id := article.ID(r.URL.Path[len(pathPrefix):])
		if id == "" {
			return id, errors.New("apiHandler: Id is empty")
		}
		return id, nil
	}
	getArticles := func() ([]article.Article, error) {
		if err := r.ParseForm(); err != nil {
			return nil, err
		}
		encodedArticles, ok := r.PostForm["article"]
		if !ok {
			return nil, errors.New("article parameter expected")
		}
		var articles []article.Article
		for _, encodedArticle := range encodedArticles {
			var a article.Article
			if err := json.Unmarshal([]byte(encodedArticle), &a); err != nil {
				return nil, err
			}
			articles = append(articles, a)
		}
		return articles, nil
	}

	switch r.Method {
	case "POST":
		articles, err := getArticles()
		if err != nil {
			log.Println(err)
			return
		}
		for _, a := range articles {
			id, err := m.Create(a)
			err = json.NewEncoder(w).Encode(Response{
				ID:      id,
				Article: a,
				Error:   ResponseError{err},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	case "GET":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		a, err := m.Read(id)
		err = json.NewEncoder(w).Encode(Response{
			ID:      id,
			Article: a,
			Error:   ResponseError{err},
		})
		if err != nil {
			log.Println(err)
			return
		}
	case "PUT":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		articles, err := getArticles()
		if err != nil {
			log.Println(err)
			return
		}
		for _, a := range articles {
			err = m.Update(id, a)
			err = json.NewEncoder(w).Encode(Response{
				ID:      id,
				Article: a,
				Error:   ResponseError{err},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	case "DELETE":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		err = m.Delete(id)
		err = json.NewEncoder(w).Encode(Response{
			ID:    id,
			Error: ResponseError{err},
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// FIXME: m is NOT thread-safe
var m = article.NewInMemoryAccessor()
