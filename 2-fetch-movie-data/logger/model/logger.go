package model

import (
	"github.com/ramailh/stockbit/2-fetch-movie-data/logger/repository/postgres"
	"gorm.io/gorm"
)

type Logger struct {
	gorm.Model
	Method     string  `json:"method"`
	ClientIP   string  `json:"client_ip"`
	Latency    float64 `json:"latency"`
	Path       string  `json:"path"`
	StatusCode int     `json:"status_code"`
	TimeStamp  int64   `json:"timestamp"`
}

func (lgr Logger) Insert() error {
	mgtr := postgres.DB.Migrator()
	if !mgtr.HasTable(&lgr) {
		mgtr.CreateTable(&lgr)
	}

	return postgres.DB.Create(&lgr).Error
}
