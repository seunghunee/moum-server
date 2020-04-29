package accessor

import (
	"fmt"

	"github.com/seunghunee/moum-server/graph/model"
)

// InMemoryAccessor is a simple in-memory database.
type InMemoryAccessor struct {
	articles []*model.Article
	nextID   int64
}

// NewInMemoryAccessor returns a new MemoryDataAccess.
func NewInMemoryAccessor() Accessor {
	return &InMemoryAccessor{
		articles: []*model.Article{},
		nextID:   int64(1),
	}
}

// Create adds a new article.
func (m *InMemoryAccessor) Create(input model.AddArticleInput) (string, error) {
	id := fmt.Sprint(m.nextID)
	m.nextID++
	m.articles = append(m.articles, &model.Article{
		ID:    id,
		Title: input.Title,
		Body:  input.Body,
	})
	return id, nil
}

// Read returns a article with a given ID.
func (m *InMemoryAccessor) Read(id string) (model.Article, error) {
	for _, a := range m.articles {
		if a.ID == id {
			return *a, nil
		}
	}
	return model.Article{}, ErrArticleNotExist
}

// Update updates a article with input.
func (m *InMemoryAccessor) Update(input model.UpdateArticleInput) error {
	for i, a := range m.articles {
		if a.ID == input.ID {
			m.articles[i] = &model.Article{
				ID:    input.ID,
				Title: input.Title,
				Body:  input.Body,
			}
			return nil
		}
	}
	return ErrArticleNotExist
}

// Delete removes the article with a given ID.
func (m *InMemoryAccessor) Delete(id string) error {
	for i, a := range m.articles {
		if a.ID == id {
			m.articles = append(m.articles[:i], m.articles[i+1:]...)
			return nil
		}
	}
	return ErrArticleNotExist
}

// List returns all items of list.
func (m *InMemoryAccessor) List() ([]*model.Article, error) {
	return m.articles, nil
}
