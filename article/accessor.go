package article

import "errors"

// Accessor is an interface to access articles.
type Accessor interface {
	Create(a Article) (ID, error)
	Read(id ID) (Article, error)
	Update(id ID, a Article) error
	Delete(id ID) error
}

// ID is a data type to identify a article.
type ID string

// ErrArticleNotExist occurs when the article with the ID was not found.
var ErrArticleNotExist = errors.New("article does not exist")
