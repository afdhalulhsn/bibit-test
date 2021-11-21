package dao

import (
	"bibit/app/model"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type LogApiClient struct {
	db *gorm.DB
}

func NewLogApiClient(conn *gorm.DB) *LogApiClient {
	return &LogApiClient{
		db: conn,
	}
}

func (a *LogApiClient) InsertLogApi(in interface{}) {
	var et *model.LogApi
	_ = mapstructure.Decode(in, &et)
	err := a.db.Create(&et).Error
	if err != nil {
		fmt.Println("Error Insert")
	}
}
