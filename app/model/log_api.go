package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id string `jason:"id" gorm:"primaryKey"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New()
	base.Id = id.String()
	return nil
}

type LogApi struct {
	BaseModel
	EndPoint       string
	ReqHeader      string
	Body           string
	Duration       string
	Response       string
	ResponseHeader string
}
