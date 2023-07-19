package memory

import (
	"context"
	"movieapp/metadata/internal/repository"
	"movieapp/metadata/pkg/model"
	"sync"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

func New() *Repository {
	return &Repository{data: map[string]*model.Metadata{}}
}

func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()

	if m, ok := r.data[id]; ok {
		return m, nil
	} else {
		return nil, repository.ErrNotFound
	}
}

func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()

	r.data[id] = metadata

	return nil
}
