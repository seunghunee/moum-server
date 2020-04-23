package graph

import "github.com/seunghunee/moum-server/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the root graph resolver type.
type Resolver struct {
	articles []*model.Article
}
