package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc(pathPrefix, apiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

const pathPrefix = "/api/v1/article/"

func apiHandler(w http.ResponseWriter, r *http.Request) {
	getID := func() (ID, error) {
		id := ID(r.URL.Path[len(pathPrefix):])
		if id == "" {
			return id, errors.New("apiHandler: Id is empty")
		}
		return id, nil
	}
	getArticles := func() ([]Article, error) {
		if err := r.ParseForm(); err != nil {
			return nil, err
		}
		encodedArticles, ok := r.PostForm["article"]
		if !ok {
			return nil, errors.New("article parameter expected")
		}
		var articles []Article
		for _, encodedArticle := range encodedArticles {
			var a Article
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
var m = NewMemoryDataAccess()

// Article is a data type for an article.
type Article struct {
	Title string
	Body  string
}

// ID is a data type to identify a article.
type ID string

// DataAccess is an interface to access articles.
type DataAccess interface {
	Create(a Article) (ID, error)
	Read(id ID) (Article, error)
	Update(id ID, a Article) error
	Delete(id ID) error
}

// MemoryDataAccess is a simple in-memory database.
type MemoryDataAccess struct {
	articles map[ID]Article
	nextID   int64
}

// NewMemoryDataAccess returns a new MemoryDataAccess.
func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess{
		articles: map[ID]Article{},
		nextID:   int64(1),
	}
}

// ErrArticleNotExist occurs when the article with the ID was not found.
var ErrArticleNotExist = errors.New("article does not exist")

// Create adds a new article.
func (m *MemoryDataAccess) Create(a Article) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.articles[id] = a
	return id, nil
}

// Read returns a article with a given ID.
func (m *MemoryDataAccess) Read(id ID) (Article, error) {
	a, exists := m.articles[id]
	if !exists {
		return Article{}, ErrArticleNotExist
	}
	return a, nil
}

// Update updates a article with a given ID with a.
func (m *MemoryDataAccess) Update(id ID, a Article) error {
	if _, exists := m.articles[id]; !exists {
		return ErrArticleNotExist
	}
	m.articles[id] = a
	return nil
}

// Delete removes the article with a given ID.
func (m *MemoryDataAccess) Delete(id ID) error {
	if _, exists := m.articles[id]; !exists {
		return ErrArticleNotExist
	}
	delete(m.articles, id)
	return nil
}

// ResponseError is the error for the JSON response.
type ResponseError struct {
	Err error
}

// MarshalJSON returns the JSON representation of the error.
func (err ResponseError) MarshalJSON() ([]byte, error) {
	if err.Err == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%v"`, err.Err)), nil
}

// Response is a struct for the JSON response.
type Response struct {
	ID      ID            `json:"id"`
	Article Article       `json:"article"`
	Error   ResponseError `json:"error"`
}
