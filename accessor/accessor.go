package accessor

import (
	"errors"

	"github.com/seunghunee/moum-server/graph/model"
)

// Accessor is an interface to access articles.
type Accessor interface {
	Create(input model.AddArticleInput) (string, error)
	Read(id string) (model.Article, error)
	Update(input model.UpdateArticleInput) error
	Delete(id string) error
	List() ([]*model.Article, error)
}

// ErrArticleNotExist occurs when the article with the ID was not found.
var ErrArticleNotExist = errors.New("article does not exist")

// ErrInvalidArg occurs when the arguments was invalid.
var ErrInvalidArg = errors.New("invalid arguments")
