package controller

import (
	proto "TEST/bibi_test/app/infrastructure/grpc/proto/movie"
	"TEST/bibi_test/app/service/movie_omdb"
	m "TEST/bibi_test/app/shared/mock"
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
	Convey("GetDetail Movie", t, func() {
		Convey("Positive Scenario", func() {
			client := &m.MovieImdbMockRepo{}
			uc := movie_omdb.NewMovieImdbServiceImpl(client)
			req := &proto.GetListMovieRequest{
				Keyword: "1",
			}
			client.On("GetListMovieOmdb", mock.Anything).Return(&proto.ResponseListFilm{}, nil)
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
				client := &m.MovieImdbMockRepo{}
				uc := movie_omdb.NewMovieImdbServiceImpl(client)
				req := &proto.GetListMovieRequest{}
				client.On("GetListMovieOmdb", mock.Anything).Return([]string{"1"}, nil)
				svc := NewMovieIMdbController(uc)
				_, err := svc.GetListMovieOmdb(context.Background(), req)

				So(err, ShouldNotBeNil)
			})
		})
	})
}
