package metadata

import (
	"context"
	"errors"
	"movieapp/metadata/internal/repository"
	"movieapp/metadata/pkg/model"
)

var ErrNotFound = errors.New("not found")

type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

type Controller struct {
	repo metadataRepository
}

func New(repo metadataRepository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
	if res, err := c.repo.Get(ctx, id); err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	} else {
		return res, err
	}
}
