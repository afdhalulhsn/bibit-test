package controller

import (
	proto "bibit/app/infrastructure/grpc/proto/movie"
	"bibit/app/model"
	"bibit/app/service/movie_omdb"
	m "bibit/app/shared/mock"
	"context"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestMovieImdbController_GetDatilMovie(t *testing.T) {
	Convey("GetDetail Movie", t, func() {
		Convey("Positive Scenario", func() {
			client := &m.MovieImdbMockRepo{}
			uc := movie_omdb.NewMovieImdbServiceImpl(client)
			client.On("GetDeatailMovie", "1").Return(&proto.ResponseGetDetailMovie{}, nil)
			svc := NewMovieIMdbController(uc)
			_, err := svc.GetDatilMovie(context.Background(), &proto.GetDetailMovieRequest{
				ImdbId: "1",
			})
			So(err, ShouldBeNil)
		})
		Convey("Negative Scenario", func() {
			Convey("Error From Service", func() {
				client := &m.MovieImdbMockRepo{}
				uc := movie_omdb.NewMovieImdbServiceImpl(client)
				client.On("GetDeatailMovie", "1").Return(nil, errors.New("Error"))
				svc := NewMovieIMdbController(uc)
				_, err := svc.GetDatilMovie(context.Background(), &proto.GetDetailMovieRequest{ImdbId: "1"})
				So(err, ShouldNotBeNil)
			})
			Convey("Error Decode Response From Res", func() {
				client := &m.MovieImdbMockRepo{}
				uc := movie_omdb.NewMovieImdbServiceImpl(client)
				client.On("GetDeatailMovie", "1").Return([]string{"1"}, nil)
				svc := NewMovieIMdbController(uc)
				_, err := svc.GetDatilMovie(context.Background(), &proto.GetDetailMovieRequest{
					ImdbId: "1",
				})
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestMovieImdbController_GetListMovieOmdb(t *testing.T) {
	Convey("Get List Movie", t, func() {
		Convey("Positive Scenario", func() {
			client := &m.MovieImdbMockRepo{}
			uc := movie_omdb.NewMovieImdbServiceImpl(client)
			req := &proto.GetListMovieRequest{
				Keyword: "1",
			}
			list := []*proto.DataFilmList{
				&proto.DataFilmList{
					ImdbId: "1",
				},
			}
			client.On("GetDeatailMovie", "1").Return(&model.ResponseGetDetailMovie{}, nil)
			client.On("GetListMovieOmdb", mock.Anything).Return(&proto.ResponseListFilm{
				ListFilm: list,
			}, nil)
			svc := NewMovieIMdbController(uc)
			_, err := svc.GetListMovieOmdb(context.Background(), req)

			So(err, ShouldBeNil)
		})
		Convey("Negative Scenario", func() {
			Convey("Error From Service", func() {
				client := &m.MovieImdbMockRepo{}
				uc := movie_omdb.NewMovieImdbServiceImpl(client)
				req := &proto.GetListMovieRequest{}
				client.On("GetListMovieOmdb", mock.Anything).Return(nil, errors.New("Error"))
				svc := NewMovieIMdbController(uc)
				_, err := svc.GetListMovieOmdb(context.Background(), req)
				So(err, ShouldNotBeNil)
			})
			Convey("Error Decode Response From Res", func() {
				client := &m.MovieImdbMockSevice{}
				//uc := movie_omdb.NewMovieImdbServiceImpl(client)
				req := &proto.GetListMovieRequest{}
				client.On("GetListMovieData", mock.Anything).Return([]string{"1"}, nil)
				svc := NewMovieIMdbController(client)
				_, err := svc.GetListMovieOmdb(context.Background(), req)

				So(err, ShouldNotBeNil)
			})
		})
	})
}
