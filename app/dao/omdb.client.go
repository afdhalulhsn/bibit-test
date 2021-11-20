package dao

import (
	"TEST/bibi_test/app/model"
	"TEST/bibi_test/app/repository"
	"TEST/bibi_test/internal/config"
	"strconv"
)

type ResponseData struct {
	ListMovie    []MovieData `json:"Search"`
	TotalResults string      `json:"totalResults"`
	Response     string      `json:"Response"`
}

type MovieData struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type MovieOmdbClient struct {
	db repository.LogApiRepo
}

func NewMovieOmdbCLient(db repository.LogApiRepo) *MovieOmdbClient {
	return &MovieOmdbClient{
		db: db,
	}
}

type ResponseDetail struct {
	Title      string    `json:"Title"`
	Year       string    `json:"Year"`
	Rated      string    `json:"Rated"`
	Released   string    `json:"Released"`
	Runtime    string    `json:"Runtime"`
	Genre      string    `json:"Genre"`
	Director   string    `json:"Director"`
	Writer     string    `json:"Writer"`
	Actors     string    `json:"Actors"`
	Plot       string    `json:"Plot"`
	Language   string    `json:"Language"`
	Country    string    `json:"Country"`
	Awards     string    `json:"Awards"`
	Poster     string    `json:"Poster"`
	Rating     []Ratings `json:"Ratings"`
	Metascore  string    `json:"Metascore"`
	ImdbRating string    `json:"imdbRating"`
	ImdbVotes  string    `json:"imdbVotes"`
	ImdbID     string    `json:"imdbID"`
	Type       string    `json:"Type"`
	Dvd        string    `json:"DVD"`
	BoxOffice  string    `json:"BoxOffice"`
	Production string    `json:"Production"`
	Website    string    `json:"Website"`
	Response   string    `json:"Response"`
}

type Ratings struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

func (a *MovieOmdbClient) GetListMovieOmdb(in interface{}) (interface{}, error) {
	cfg := config.GetConfig().Client.Omdb
	req := in.(*model.GetListMovieRequest)
	pages := strconv.Itoa(int(req.Page))
	param := map[string]string{
		"s":    req.Keyword,
		"page": pages,
	}
	cli := NewMoviewOmdbClient(cfg.Url, cfg.ApiKey, a.db)
	cli.SetQueryParams(param)
	var resp *ResponseData
	err := cli.GetData(&resp)
	if err != nil {
		return nil, err
	}
	var list []*model.DataFilmList
	for _, r := range resp.ListMovie {
		list = append(list, &model.DataFilmList{
			Type:   r.Type,
			Poster: r.Poster,
			Title:  r.Title,
			Year:   r.Year,
			ImdbId: r.ImdbID,
		})
	}
	out := &model.ResponseListFilm{
		ListFilm:    list,
		TotalResult: resp.TotalResults,
	}
	return out, nil
}

func (a *MovieOmdbClient) GetDeatailMovie(id string) (interface{}, error) {
	cfg := config.GetConfig().Client.Omdb
	param := map[string]string{
		"i": id,
	}
	cli := NewMoviewOmdbClient(cfg.Url, cfg.ApiKey, a.db)
	cli.SetQueryParams(param)
	var resp *ResponseDetail
	err := cli.GetData(&resp)
	if err != nil {
		return nil, err
	}
	var list []*model.Ratings
	for _, r := range resp.Rating {
		list = append(list, &model.Ratings{
			Source: r.Source,
			Value:  r.Value,
		})
	}
	out := &model.ResponseGetDetailMovie{
		Rating:     list,
		Year:       resp.Year,
		Title:      resp.Title,
		Type:       resp.Type,
		Poster:     resp.Poster,
		Actors:     resp.Actors,
		Awards:     resp.Awards,
		BoxOffice:  resp.BoxOffice,
		Country:    resp.Country,
		Director:   resp.Director,
		Genre:      resp.Genre,
		Response:   resp.Response,
		DVD:        resp.Dvd,
		ImdbRating: resp.ImdbRating,
		ImdbVotes:  resp.ImdbVotes,
		Language:   resp.Language,
		Metascore:  resp.Metascore,
		Plot:       resp.Plot,
		Production: resp.Production,
		Rated:      resp.Rated,
		Released:   resp.Released,
		Runtime:    resp.Runtime,
		Website:    resp.Website,
		Writer:     resp.Writer,
		ImdbID:     resp.ImdbID,
	}
	return out, nil
}
