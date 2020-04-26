package accessor

import (
	"fmt"

	"github.com/seunghunee/moum-server/graph/model"
)

// InMemoryAccessor is a simple in-memory database.
type InMemoryAccessor struct {
	articles map[string]model.Article
	nextID   int64
}

// NewInMemoryAccessor returns a new MemoryDataAccess.
func NewInMemoryAccessor() Accessor {
	return &InMemoryAccessor{
		articles: map[string]model.Article{},
		nextID:   int64(1),
	}
}

// Create adds a new article.
func (m *InMemoryAccessor) Create(input model.AddArticleInput) (string, error) {
	id := fmt.Sprint(m.nextID)
	m.nextID++
	m.articles[id] = model.Article{
		ID:    id,
		Title: input.Title,
		Body:  input.Body,
	}
	return id, nil
}

// Read returns a article with a given ID.
func (m *InMemoryAccessor) Read(id string) (model.Article, error) {
	a, exists := m.articles[id]
	if !exists {
		return model.Article{}, ErrArticleNotExist
	}
	return a, nil
}

// Update updates a article with input.
func (m *InMemoryAccessor) Update(input model.UpdateArticleInput) error {
	if _, exists := m.articles[input.ID]; !exists {
		return ErrArticleNotExist
	}
	m.articles[input.ID] = model.Article{
		ID:    input.ID,
		Title: input.Title,
		Body:  input.Body,
	}
	return nil
}

// Delete removes the article with a given ID.
func (m *InMemoryAccessor) Delete(id string) error {
	if _, exists := m.articles[id]; !exists {
		return ErrArticleNotExist
	}
	delete(m.articles, id)
	return nil
}
