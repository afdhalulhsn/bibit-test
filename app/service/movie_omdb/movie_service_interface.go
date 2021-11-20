package movie_omdb

type MovieImdbInputPort interface {
	GetListMovieData(in interface{}) (interface{}, error)
	GetDetailMovie(in string) (interface{}, error)
}
