package main

import (
	"fmt"

	"github.com/seunghunee/moum-server/article"
)

// Response is a struct for the JSON response.
type Response struct {
	ID      article.ID      `json:"id"`
	Article article.Article `json:"article"`
	Error   ResponseError   `json:"error"`
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
