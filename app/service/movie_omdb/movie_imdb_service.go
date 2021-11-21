package movie_omdb

import (
	"bibit/app/model"
	"bibit/app/repository"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
)

type MovieImdbServiceImpl struct {
	repo repository.MoveiOmdbRepo
}

func NewMovieImdbServiceImpl(repo repository.MoveiOmdbRepo) *MovieImdbServiceImpl {
	return &MovieImdbServiceImpl{
		repo: repo,
	}
}

func (a *MovieImdbServiceImpl) GetListMovieData(in interface{}) (interface{}, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "Request Cannot Be Nil")
	}
	var req *model.GetListMovieRequest
	err := mapstructure.Decode(in, &req)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error Decoding Request")
	}
	resp, err := a.repo.GetListMovieOmdb(req)
	if err != nil {
		return nil, err
	}
	var out *model.ResponseListFilm
	err = mapstructure.Decode(resp, &out)
	wg := &sync.WaitGroup{}

	for _, dt := range out.ListFilm {
		wg.Add(1)
		go func(in *model.DataFilmList) {
			defer wg.Done()
			res, err := a.repo.GetDeatailMovie(in.ImdbId)
			if err != nil {
				fmt.Println("Error In Data =>", in.ImdbId)
			}
			in.Detail = res.(*model.ResponseGetDetailMovie)
		}(dt)
	}

	wg.Wait()

	return resp, nil
}

func (a *MovieImdbServiceImpl) GetDetailMovie(in string) (interface{}, error) {
	if in == "" {
		return nil, status.Error(codes.InvalidArgument, "ID Cannot Be Empty")
	}

	resp, err := a.repo.GetDeatailMovie(in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
