package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/seunghunee/moum-server/graph/generated"
	"github.com/seunghunee/moum-server/graph/model"
)

func (r *mutationResolver) AddArticle(ctx context.Context, input model.AddArticleInput) (*model.AddArticlePayload, error) {
	id, err := r.Accessor.Create(input)
	if err != nil {
		return &model.AddArticlePayload{}, err
	}
	a, err := r.Accessor.Read(id)
	if err != nil {
		return &model.AddArticlePayload{}, err
	}
	return &model.AddArticlePayload{Article: &a}, nil
}

func (r *mutationResolver) UpdateArticle(ctx context.Context, input model.UpdateArticleInput) (*model.UpdateArticlePayload, error) {
	if err := r.Accessor.Update(input); err != nil {
		return &model.UpdateArticlePayload{}, err
	}
	a, err := r.Accessor.Read(input.ID)
	if err != nil {
		return &model.UpdateArticlePayload{}, err
	}
	return &model.UpdateArticlePayload{Article: &a}, nil
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	a, err := r.Accessor.Read(id)
	if err != nil {
		return model.Article{}, err
	}
	return a, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
