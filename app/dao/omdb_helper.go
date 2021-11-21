package dao

import (
	"bibit/app/model"
	"bibit/app/repository"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type MovieOmdbHelper struct {
	url        string
	apiKey     string
	parameters map[string]string
	db         repository.LogApiRepo
}

func NewMoviewOmdbClient(url, apikey string, db repository.LogApiRepo) *MovieOmdbHelper {
	return &MovieOmdbHelper{
		url: url,
		parameters: map[string]string{
			"apikey": apikey,
		},
		db: db,
	}
}

func (a *MovieOmdbHelper) SetQueryParams(req map[string]string) {
	for key, val := range req {
		a.parameters[key] = val
	}
}

func (a *MovieOmdbHelper) GetData(out interface{}) error {
	cli := resty.New()
	times := time.Now()
	log := &model.LogApi{
		EndPoint: a.url,
		Body:     "*",
	}
	resp, err := cli.R().SetQueryParams(a.parameters).Get(a.url)
	if err != nil {
		log.Response = err.Error()
		go a.db.InsertLogApi(log)
		return status.Error(codes.Internal, "Error To Get Data")
	}
	end := fmt.Sprint(time.Now().Sub(times).Milliseconds())
	log.Duration = end
	log.ReqHeader = fmt.Sprintf("%v", resp.Request.Header)
	if resp.StatusCode() != 200 {
		log.ResponseHeader = fmt.Sprintf("%v", resp.RawResponse.Header)
		log.Response = fmt.Sprintf("%v", resp.Body())
		go a.db.InsertLogApi(log)
		return status.Errorf(codes.Internal, "HTPP Status Not Success")
	}

	var body = resp.Body()
	if err = json.Unmarshal(body, &out); err != nil {
		log.Response = err.Error()
		go a.db.InsertLogApi(log)
		return status.Errorf(codes.Internal, "Error Marshaal Data")
	}
	log.ResponseHeader = fmt.Sprintf("%v", resp.RawResponse.Header)
	log.Response = fmt.Sprintf("%v", string(resp.Body()))
	go a.db.InsertLogApi(log)
	return nil
}
