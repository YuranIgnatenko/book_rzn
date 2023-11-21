package datatemp

import (
	"backend/config"
	"backend/models"
)

type DataTemp struct {
	config.Configuration
	TargetCards     []models.TargetCard
	FavoritesCards  []models.FavoritesCards
	OrdersCards     []models.OrdersCards
	OrdersRows      []models.OrdersRows
	SearchTarget    []models.TargetCard
	FastOrdersList  []models.DataFastOrderOne
	NumberFastOrder string
	IsLogin         bool
	NameLogin       string
}

func NewDataTemp(c config.Configuration, ps []models.TargetCard) *DataTemp {
	return &DataTemp{
		Configuration:   c,
		TargetCards:     ps,
		IsLogin:         false,
		NumberFastOrder: "",
	}
}
