package datatemp

import (
	"backend/config"
	"backend/models"
)

type DataTemp struct {
	config.Configuration
	ProsvCards []models.ProsvCard
}

func NewDataTemp(c config.Configuration, ps []models.ProsvCard) *DataTemp {
	return &DataTemp{
		Configuration: c,
		ProsvCards:    ps,
	}
}
