package validation

import (
	"reflect"
	"testing"

	"github.com/16francs/gran/api/group/internal/application/request"
	"github.com/16francs/gran/api/group/internal/domain"
)

func TestGroupRequestValidation_CreateGroup(t *testing.T) {
	// Defined variables
	g := &request.CreateGroup{
		Name:        "テストグループ",
		Description: "説明",
	}

	// Start test
	target := NewGroupRequestValidation()

	want := []*domain.ValidationError(nil)

	got := target.CreateGroup(g)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}

func TestGroupRequestValidation_UpdateGroup(t *testing.T) {
	// Defined variables
	g := &request.UpdateGroup{
		Name:        "テストグループ",
		Description: "説明",
	}

	// Start test
	target := NewGroupRequestValidation()

	want := []*domain.ValidationError(nil)

	got := target.UpdateGroup(g)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
