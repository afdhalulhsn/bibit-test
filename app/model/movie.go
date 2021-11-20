package model

type GetListMovieRequest struct {
	Keyword string
	Page    int32
}
type DataFilmList struct {
	Title  string
	Year   string
	ImdbId string
	Type   string
	Poster string
	Detail *ResponseGetDetailMovie
}
type ResponseListFilm struct {
	ListFilm    []*DataFilmList
	TotalResult string
}

type GetDetailMovieRequest struct {
	ImdbId string
}
type ResponseGetDetailMovie struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Metascore  string
	ImdbRating string
	ImdbVotes  string
	ImdbID     string
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
	Rating     []*Ratings
}

type Ratings struct {
	Source string
	Value  string
}
