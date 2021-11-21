package movie_omdb

import (
	"bibit/app/model"
	m "bibit/app/shared/mock"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestMovieImdbServiceImpl_GetDetailMovie(t *testing.T) {
	Convey("GetDetail Moovier", t, func() {
		Convey("Positive Scenario", func() {
			client := &m.MovieImdbMockRepo{}
			uc := NewMovieImdbServiceImpl(client)
			client.On("GetDeatailMovie", "Hehehe").Return(mock.Anything, nil)
			_, err := uc.GetDetailMovie("Hehehe")
			So(err, ShouldBeNil)
		})
		Convey("Negative Scenario", func() {
			Convey("Request Empty", func() {
				client := &m.MovieImdbMockRepo{}
				uc := NewMovieImdbServiceImpl(client)
				client.On("GetDeatailMovie", "").Return(mock.Anything, nil)
				_, err := uc.GetDetailMovie("")
				So(err, ShouldNotBeNil)
			})
			Convey("Error From Res Imdb", func() {
				client := &m.MovieImdbMockRepo{}
				uc := NewMovieImdbServiceImpl(client)
				client.On("GetDeatailMovie", "H").Return(nil, errors.New("Error"))
				_, err := uc.GetDetailMovie("H")
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestMovieImdbServiceImpl_GetListMovieData(t *testing.T) {
	Convey("GetListMovie Data", t, func() {
		Convey("Positive Scenario", func() {
			client := &m.MovieImdbMockRepo{}
			uc := NewMovieImdbServiceImpl(client)
			req := &model.GetListMovieRequest{
				Page:    2,
				Keyword: "2",
			}
			var lis []*model.DataFilmList

			lis = append(lis, &model.DataFilmList{
				ImdbId: "1",
			})
			out := &model.ResponseListFilm{
				ListFilm:    lis,
				TotalResult: "1",
			}
			client.On("GetDeatailMovie", "1").Return(&model.ResponseGetDetailMovie{}, nil)
			client.On("GetListMovieOmdb", req).Return(out, nil)
			_, err := uc.GetListMovieData(req)
			So(err, ShouldBeNil)
		})
		Convey("Negative Scenario", func() {
			Convey("Request Empty", func() {
				client := &m.MovieImdbMockRepo{}
				uc := NewMovieImdbServiceImpl(client)
				_, err := uc.GetListMovieData(nil)
				So(err, ShouldNotBeNil)
			})
			Convey("Error Decode Response", func() {
				client := &m.MovieImdbMockRepo{}
				uc := NewMovieImdbServiceImpl(client)
				req := &model.GetListMovieRequest{
					Page:    2,
					Keyword: "2",
				}
				client.On("GetListMovieOmdb", req).Return([]string{"1"}, nil)
				_, err := uc.GetListMovieData(req)
				So(err, ShouldNotBeNil)
			})
			Convey("Error From Res Imdb", func() {
				client := &m.MovieImdbMockRepo{}
				uc := NewMovieImdbServiceImpl(client)
				req := &model.GetListMovieRequest{
					Page:    2,
					Keyword: "2",
				}
				client.On("GetListMovieOmdb", req).Return(nil, errors.New("Error"))
				_, err := uc.GetListMovieData(req)
				So(err, ShouldNotBeNil)
			})
			Convey("Error Decode Request", func() {
				client := &m.MovieImdbMockRepo{}
				uc := NewMovieImdbServiceImpl(client)
				_, err := uc.GetListMovieData([]string{"1"})
				So(err, ShouldNotBeNil)
			})
			Convey("When Get Detail Error,", func() {
				client := &m.MovieImdbMockRepo{}
				uc := NewMovieImdbServiceImpl(client)
				req := &model.GetListMovieRequest{
					Page:    2,
					Keyword: "2",
				}
				var lis []*model.DataFilmList

				lis = append(lis, &model.DataFilmList{
					ImdbId: "1",
				})
				out := &model.ResponseListFilm{
					ListFilm:    lis,
					TotalResult: "1",
				}
				client.On("GetDeatailMovie", "1").Return(nil, errors.New("Err"))
				client.On("GetListMovieOmdb", req).Return(out, nil)
				_, err := uc.GetListMovieData(req)
				So(err, ShouldBeNil)
			})
		})
	})
}
