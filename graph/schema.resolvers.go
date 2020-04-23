package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/seunghunee/moum-server/graph/generated"
	"github.com/seunghunee/moum-server/graph/model"
)

func (r *mutationResolver) AddArticle(ctx context.Context, input model.AddArticleInput) (*model.AddArticlePayload, error) {
	article := &model.Article{
		ID:    fmt.Sprintf("%d", rand.Int()),
		Title: input.Title,
		Body:  input.Body,
	}
	r.articles = append(r.articles, article)
	return &model.AddArticlePayload{Article: article}, nil
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	for _, a := range r.articles {
		if a.ID == id {
			return *a, nil
		}
	}
	return model.Article{}, errors.New("article does not exist")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
