package movie

import (
	"context"
	"github.com/tunx321/movie-microservices/internal/domain/models"
	"time"
)

//  rpc GetMovieByID(GetMovieRequest) returns (Movie);
//  rpc InsertMovie(Movie) returns (MovieResponse);
//  rpc DeleteMovie(GetMovieRequest) returns (MovieResponse);

type Movie interface {
	GetMovieByID(
		ctx context.Context,
		id int64,
	) (models.Movie, error)

	InsertMovie(
		ctx context.Context,
		title string,
		author string,
		duration time.Duration,
		description string,
	) (id int64, err error)

	DeleteMovie(ctx context.Context, id int64) (int64, error)
}

type MovieAPI struct {
}
