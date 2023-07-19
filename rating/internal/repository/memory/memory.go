package memory

import (
	"context"
	"movieapp/rating/internal/repository"
	"movieapp/rating/pkg/model"
)

type Repository struct {
	data map[model.RecordType](map[model.RecordID]([]model.Rating))
}

func New() *Repository {
	return &Repository{data: map[model.RecordType]map[model.RecordID][]model.Rating{}}
}

func (r *Repository) Get(ctx context.Context, recordId model.RecordID, recordType model.RecordType) ([]model.Rating, error) {
	if recordTypeMap, ok := r.data[recordType]; ok {
		if ratings, ok := recordTypeMap[recordId]; ok {
			return ratings, nil
		}
	}
	return nil, repository.ErrNotFound
}

func (r *Repository) Put(ctx context.Context, recordId model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	if _, ok := r.data[recordType]; !ok {
		r.data[recordType] = map[model.RecordID][]model.Rating{}
	}

	r.data[recordType][recordId] = append(r.data[recordType][recordId], *rating)

	return nil
}
