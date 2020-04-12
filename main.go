package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, 세계!\n")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

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
