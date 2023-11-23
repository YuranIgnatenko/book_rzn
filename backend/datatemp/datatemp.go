package datatemp

import (
	"backend/config"
	"backend/models"
	"fmt"
)

type DataTemp struct {
	config.Configuration
	TargetAll         []models.TargetCard
	TargetCards       []models.TargetCard
	FavoritesCards    []models.FavoritesCards
	OrdersCards       []models.OrdersCards
	OrdersRows        []models.OrdersRows
	SearchTarget      []models.TargetCard
	FastOrdersList    []models.DataFastOrderOne
	NumberFastOrder   string
	IsLogin           bool
	NameLogin         string
	PageMenuTemplHtml string
}

func NewDataTemp(c config.Configuration, ps []models.TargetCard) *DataTemp {
	return &DataTemp{
		Configuration:   c,
		TargetCards:     ps,
		IsLogin:         false,
		NameLogin:       "Гость",
		NumberFastOrder: "",
		PageMenuTemplHtml: `
		<style>
                        .nav-bar-btn {
                            float: left;
                        }

                        .btn {
                            background-color: #fca783;
                            border-radius: 5px;
                            border: 0;
                            margin: 5px;

                        }

                        .btn:hover {
                            background-color: #e18057;
                        }
                    </style>
                    <div class="nav-bar">

                        <div class="nav-bar-btn">
                            <form action="/home">
                                <input class="btn" type="submit" value="Главная" />
                            </form>
                        </div>

                        <div class="nav-bar-btn">
                            <form action="/for_school">
                                <input class="btn" type="submit" value="Для школы" />
                            </form>
                        </div>

                        <div class="nav-bar-btn">
                            <form action="for_office">
                                <input class="btn" type="submit" value="Для офиса" />
                            </form>
                        </div>
                    </div>
		`,
	}
}

func (dt *DataTemp) FilterCards(data []models.TargetCard, mode string) []models.TargetCard {
	segm := make([]models.TargetCard, 0)

	for _, tc := range dt.TargetCards {
		fmt.Println(tc.Tag)
		if mode == tc.Tag {
			segm = append(segm, tc)
		}
	}
	return segm
}
