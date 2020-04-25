package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/seunghunee/moum-server/graph/model"

// Resolver is the root graph resolver type.
type Resolver struct {
	articles []*model.Article
}
