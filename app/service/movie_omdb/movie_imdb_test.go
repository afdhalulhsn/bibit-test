package movie_omdb

import (
	"TEST/bibi_test/app/model"
	m "TEST/bibi_test/app/shared/mock"
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
			client.On("GetListMovieOmdb", req).Return(mock.Anything, nil)
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
		})
	})
}
