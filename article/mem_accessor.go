package article

import (
	"fmt"
)

// InMemoryAccessor is a simple in-memory database.
type InMemoryAccessor struct {
	articles map[ID]Article
	nextID   int64
}

// NewInMemoryAccessor returns a new MemoryDataAccess.
func NewInMemoryAccessor() Accessor {
	return &InMemoryAccessor{
		articles: map[ID]Article{},
		nextID:   int64(1),
	}
}

// Create adds a new
func (m *InMemoryAccessor) Create(a Article) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.articles[id] = a
	return id, nil
}

// Read returns a article with a given ID.
func (m *InMemoryAccessor) Read(id ID) (Article, error) {
	a, exists := m.articles[id]
	if !exists {
		return Article{}, ErrArticleNotExist
	}
	return a, nil
}

// Update updates a article with a given ID with a.
func (m *InMemoryAccessor) Update(id ID, a Article) error {
	if _, exists := m.articles[id]; !exists {
		return ErrArticleNotExist
	}
	m.articles[id] = a
	return nil
}

// Delete removes the article with a given ID.
func (m *InMemoryAccessor) Delete(id ID) error {
	if _, exists := m.articles[id]; !exists {
		return ErrArticleNotExist
	}
	delete(m.articles, id)
	return nil
}
