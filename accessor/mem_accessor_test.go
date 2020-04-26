package accessor

import (
	"testing"

	"github.com/seunghunee/moum-server/graph/model"
)

func TestArticleCRUD(t *testing.T) {
	body1 := "test body1"
	body2 := "test body2"
	body3 := "test body3"
	input1 := model.AddArticleInput{Title: "test title1", Body: &body2}
	input2 := model.AddArticleInput{Title: "test title2", Body: &body1}
	input3 := model.AddArticleInput{Title: "test title3", Body: &body3}

	d := NewInMemoryAccessor()

	id, _ := d.Create(input1)
	readArticle, err := d.Read(id)
	if err != nil {
		t.Error(err.Error())
	}
	if !hasSameTitleAndBody(input1, readArticle) {
		t.Errorf("got: %v, want: %v", readArticle, input1)
	}
	id, _ = d.Create(input2)
	article2id := id
	readArticle, err = d.Read(id)
	if err != nil {
		t.Error(err.Error())
	}
	if !hasSameTitleAndBody(input2, readArticle) {
		t.Errorf("got: %v, want: %v", readArticle, input2)
	}
	id, _ = d.Create(input3)
	readArticle, err = d.Read(id)
	if err != nil {
		t.Error(err.Error())
	}
	if !hasSameTitleAndBody(input3, readArticle) {
		t.Errorf("got: %v, want: %v", readArticle, input3)
	}

	updatedBody2 := "updated body2"
	updatedArticleInput := model.AddArticleInput{Title: "updated title2", Body: &updatedBody2}
	if err := d.Update(article2id, updatedArticleInput); err != nil {
		t.Error(err.Error())
	}
	readArticle, err = d.Read(article2id)
	if err != nil {
		t.Error(err.Error())
	}
	if !hasSameTitleAndBody(updatedArticleInput, readArticle) {
		t.Errorf("got: %v, want: %v", readArticle, updatedArticleInput)
	}

	if err := d.Delete(article2id); err != nil {
		t.Error(err.Error())
	}
	if readArticle, err := d.Read(article2id); err == nil {
		t.Errorf("the article with id(%v) until exists: %v", article2id, readArticle)
	}
}

func hasSameTitleAndBody(input model.AddArticleInput, article model.Article) bool {
	if input.Title == article.Title && input.Body == article.Body {
		return true
	}
	return false
}
