package datatemp

import (
	"backend/config"
	"backend/models"
)

type DataTemp struct {
	config.Configuration
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
		PageMenuTemplHtml: `<div class="page-menu">
		<ul id="menu_list" class="level_1">
			<li class="li"><a href="http://localhost:8082/home" class=""
				title="О компании">Главная</a></li>
			 <li class="li"><a href="http://localhost:8082/804" class="" 
				title="О компании">Приказ804</a></li>
			<li class="li"><a href="http://localhost:8082/prosv" class=""
					title="О компании">Просвещение</a></li>
			<li class="li"><a href="http://localhost:8082/agat" class=""
					title="О компании">Агат</a></li>
		</ul>
	</div>
	<div class="mob_menu"><i data-feather="menu"></i>Навигация</div>
	<div class="menu_popup_mob">
		<div class="catalog_subtitle">Навигация по сайту</div>
		<ul id="menu_mob">
			 <li class="li"><a href="http://localhost:8082/804" class=""
					title="О компании">Приказ804</a></li>
			<li class="li"><a href="http://localhost:8082/prosv" class=""
					title="О компании">Просвещение</a></li>
			<li class="li"><a href="http://localhost:8082/agat" class=""
					title="О компании">Агат</a></li>
		</ul>
	</div>`,
	}
}
