package article

import (
	"reflect"
	"testing"
)

func TestArticleCRUD(t *testing.T) {
	a2 := Article{Title: "test title2", Body: "test body2"}
	a1 := Article{Title: "test title1", Body: "test body1"}
	a3 := Article{Title: "test title3", Body: "test body3"}

	d := NewInMemoryAccessor()

	id, _ := d.Create(a1)
	readArticle, err := d.Read(id)
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(readArticle, a1) {
		t.Errorf("got: %q, want: %s", readArticle, a1)
	}
	id, _ = d.Create(a2)
	article2id := id
	readArticle, err = d.Read(id)
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(readArticle, a2) {
		t.Errorf("got: %q, want: %s", readArticle, a2)
	}
	id, _ = d.Create(a3)
	readArticle, err = d.Read(id)
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(readArticle, a3) {
		t.Errorf("got: %q, want: %s", readArticle, a3)
	}

	updatedArticle := Article{Title: "updated title2", Body: "updated body2"}
	if err := d.Update(article2id, updatedArticle); err != nil {
		t.Error(err.Error())
	}
	readArticle, err = d.Read(article2id)
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(readArticle, updatedArticle) {
		t.Errorf("got: %q, want: %s", readArticle, updatedArticle)
	}

	if err := d.Delete(article2id); err != nil {
		t.Error(err.Error())
	}
	if readArticle, err := d.Read(article2id); err == nil {
		t.Errorf("the article with id(%q) until exists: %q", article2id, readArticle)
	}
}
