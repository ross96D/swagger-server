package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/ross96d/simple_swagger_server/models"
	"github.com/ross96d/simple_swagger_server/shared"
)

var artistss = []models.Artist{
	{
		ID:   1,
		Name: "Carlo",
	},
	{
		ID:   2,
		Name: "Camilo",
	},
}

var tag string

func Register(api huma.API, base string) {
	tag = base

	if !strings.HasPrefix(base, "/") {
		base = "/" + base
	}
	if !strings.HasSuffix(base, "/") {
		base += "/"
	}

	path := func(s string) string {
		return fmt.Sprintf("%s%s", base, s)
	}

	huma.Register[shared.EmptyInput, shared.DefResp[[]models.Artist]](api, huma.Operation{
		Method:  http.MethodGet,
		Summary: "Get all artists",
		Path:    path(""),
		Tags:    []string{tag},
	}, func(ctx context.Context, a *shared.EmptyInput) (*shared.DefResp[[]models.Artist], error) {
		return artists(ctx)
	})

	huma.Register[artistInput, shared.DefResp[models.Artist]](api, huma.Operation{
		Method:  http.MethodGet,
		Summary: "Get artist by id",
		Path:    path("{id}"),
		Tags:    []string{tag},
	}, func(ctx context.Context, id *artistInput) (*shared.DefResp[models.Artist], error) {
		return artist(ctx, id)
	})
}

// get all artists
func artists(ctx context.Context) (*shared.DefResp[[]models.Artist], error) {
	return &shared.DefResp[[]models.Artist]{
		Body: artistss,
	}, nil
}

type artistInput struct {
	ID int `path:"id"`
}

func artist(ctx context.Context, id *artistInput) (*shared.DefResp[models.Artist], error) {
	if id.ID > len(artistss) || id.ID <= 0 {
		return nil, errors.New("no artist found")
	}
	fmt.Printf("RETURNING %+v\n", artistss[id.ID])
	return &shared.DefResp[models.Artist]{
		Body: artistss[id.ID],
	}, nil
}
