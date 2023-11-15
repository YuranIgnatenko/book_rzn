package datatemp

import (
	"backend/config"
	"backend/models"
)

type DataTemp struct {
	config.Configuration
	ProsvCards     []models.ProsvCard
	FavoritesCards []models.FavoritesCards
}

func NewDataTemp(c config.Configuration, ps []models.ProsvCard) *DataTemp {
	return &DataTemp{
		Configuration: c,
		ProsvCards:    ps,
	}
}
