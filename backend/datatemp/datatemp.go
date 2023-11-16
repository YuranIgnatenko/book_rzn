package datatemp

import (
	"backend/config"
	"backend/models"
)

type DataTemp struct {
	config.Configuration
	TargetCards    []models.TargetCard
	FavoritesCards []models.FavoritesCards
}

func NewDataTemp(c config.Configuration, ps []models.TargetCard) *DataTemp {
	return &DataTemp{
		Configuration: c,
		TargetCards:   ps,
	}
}
