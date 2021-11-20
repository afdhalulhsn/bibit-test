package repository

type MoveiOmdbRepo interface {
	GetListMovieOmdb(interface{}) (interface{}, error)
	GetDeatailMovie(string2 string) (interface{}, error)
}
