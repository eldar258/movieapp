package movie

import (
	"context"
	"errors"
	metadatamodel "movieapp/metadata/pkg/model"
	"movieapp/movie/internal/gateway"
	"movieapp/movie/pkg/model"
	ratingmodel "movieapp/rating/pkg/model"
)

var ErrNotFound = errors.New("movie metadata not found")

type ratingGateway interface {
	GetAggregatingRating(ctx context.Context, recordId ratingmodel.RecordID, recordType ratingmodel.RecordType) (float64, error)
	PutRating(ctx context.Context, recordId ratingmodel.RecordID, recordType ratingmodel.RecordType, rating ratingmodel.Rating) error
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadatamodel.Metadata, error)
}

type Controller struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

func New(ratingGateway ratingGateway, metadataGateway metadataGateway) *Controller {
	return &Controller{ratingGateway: ratingGateway, metadataGateway: metadataGateway}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	metadata, err := c.metadataGateway.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gateway.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	details := &model.MovieDetails{Metadata: *metadata}
	rating, err := c.ratingGateway.GetAggregatingRating(ctx, ratingmodel.RecordID(id), ratingmodel.RecordTypeMovie)
	if err != nil {
		if !errors.Is(err, gateway.ErrNotFound) {
			//TODO
			//Just proceed in this case, it's ok not to have ratings yet.
		}
		return nil, err
	}

	details.Rating = &rating
	return details, nil
}
