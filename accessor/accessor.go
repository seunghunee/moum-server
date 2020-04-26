package accessor

import (
	"errors"

	"github.com/seunghunee/moum-server/graph/model"
)

// Accessor is an interface to access articles.
type Accessor interface {
	Create(input model.AddArticleInput) (string, error)
	Read(id string) (model.Article, error)
	Update(id string, input model.AddArticleInput) error
	Delete(id string) error
}

// ErrArticleNotExist occurs when the article with the ID was not found.
var ErrArticleNotExist = errors.New("article does not exist")