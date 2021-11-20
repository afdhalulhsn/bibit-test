package controller

import (
	proto "TEST/bibi_test/app/infrastructure/grpc/proto/movie"
	"TEST/bibi_test/app/service/movie_omdb"
	"context"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MovieImdbController struct {
	svc movie_omdb.MovieImdbInputPort
}

func NewMovieIMdbController(svc movie_omdb.MovieImdbInputPort) *MovieImdbController {
	return &MovieImdbController{
		svc: svc,
	}
}

func (a *MovieImdbController) GetListMovieOmdb(ctx context.Context, in *proto.GetListMovieRequest) (*proto.ResponseListFilm, error) {
	res, err := a.svc.GetListMovieData(in)
	if err != nil {
		return nil, err
	}
	var out *proto.ResponseListFilm
	err = mapstructure.Decode(res, &out)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error Decode Response In Controller")
	}
	return out, nil
}

func (a *MovieImdbController) GetDatilMovie(ctx context.Context, in *proto.GetDetailMovieRequest) (*proto.ResponseGetDetailMovie, error) {
	res, err := a.svc.GetDetailMovie(in.GetImdbId())
	if err != nil {
		return nil, err
	}
	var out *proto.ResponseGetDetailMovie
	err = mapstructure.Decode(res, &out)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error Decode Response In Controller")
	}
	return out, nil
}
