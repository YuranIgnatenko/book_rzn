package datatemp

import (
	"backend/config"
	"backend/models"
	"strings"
)

type DataTemp struct {
	config.Configuration
	TargetAll         []models.TargetCard
	TargetCards       []models.TargetCard
	// OrdersCardsCms    []models.OrdersCardsCms
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
		TargetAll:       ps,
		// OrdersCardsCms:  make([]models.OrdersCardsCms, 0),
		IsLogin:         false,
		NameLogin:       "Гость",
		NumberFastOrder: "",
		PageMenuTemplHtml: `
        <style>
        * {
            box-sizing: border-box;
        }

        body {
            margin: 0;
            background: #f2f2f2;
        }

        header {
            background: white;
            text-align: center;
        }

        header a {
            text-decoration: none;
            outline: none;
            display: block;
            transition: .3s ease-in-out;
        }

        nav {
            display: table;
            margin: 0 auto;
        }

        nav ul {
            list-style: none;
            margin: 0;
            padding: 0;
        }

        .topmenu:after {
            content: "";
            display: table;
            clear: both;
        }

        .topmenu {
            width: 100%;
        }

        .topmenu>li {
            width: 15%;
            float: left;
            position: relative;
            font-family: 'Open Sans', sans-serif;
        }

        .topmenu>li>a {
            text-transform: uppercase;
            font-size: 14px;
            font-weight: bold;
            color: #404040;
            padding: 15px 30px;
        }

        .topmenu li a:hover {
            color: #e66464;
        }

        .submenu-link:after {
            content: "\f107";
            font-family: "FontAwesome";
            color: inherit;
            margin-left: 10px;
        }

        .submenu {
            background: #273037;
            position: absolute;
            left: 0;
            top: 100%;
            z-index: 5;
            width: 180px;
            opacity: 0;
            transform: scaleY(0);
            transform-origin: 0 0;
            transition: .2s ease-in-out;
        }

        .submenu a {
            color: white;
            text-align: left;
            padding: 12px 15px;
            font-size: 13px;
            border-bottom: 1px solid rgba(255, 255, 255, .1);
        }

        .submenu li:last-child a {
            border-bottom: none;
        }

        .topmenu>li:hover .submenu {
            opacity: 1;
            transform: scaleY(1);
        }
    </style>

		            <nav>
                        <ul class="topmenu">
                            <li><a href="/home">Главная</a></li>
                            <li><a href="" class="submenu-link">Новинки</a>
                                <ul class="submenu">
                                    <li><a href="/new_basic">Базовые модули</a></li>
                                    <li><a href="/new_boxing">Системы хранения</a></li>
                                </ul>
                            </li>
                            <li><a href="" class="submenu-link">Школы</a>
                                <ul class="submenu">
                                    <li><a href="/sh_table">Рабочие столы</a></li>
                                    <li><a href="/sh_chair">Рабочие стулья</a></li>
                                    <li><a href="/sh_minitable">Тумба под доску</a></li>
                                </ul>
                            </li>
                            <li><a href="" class="submenu-link">Офисы</a>
                                <ul class="submenu">
                                    <li><a href="/office_table">Рабочие столы</a></li>
                                    <li><a href="/office_boxing">Системы хранения</a></li>
                                </ul>
                            </li>
                            <li><a href="/book_prosv">Книги</a></li>
                        </ul>
                    </nav>
		`,
	}
}

func (dt *DataTemp) FilterCards(data []models.TargetCard, mode string) []models.TargetCard {
	segm := make([]models.TargetCard, 0)

	for _, tc := range data {
		if mode == tc.Tag {
			segm = append(segm, tc)
		}
	}
	return segm
}

func (dt *DataTemp) FilterSearch(data []models.TargetCard, sub string) []models.TargetCard {
	segm := make([]models.TargetCard, 0)

	cvt := func(sub1, sub2 string) (string, string) {
		return strings.ToLower(sub1), strings.ToLower(sub2)
	}

	for _, tc := range data {
		if strings.Contains(cvt(tc.Autor, sub)) || strings.Contains(cvt(tc.Price, sub)) ||
			strings.Contains(cvt(tc.Title, sub)) || strings.Contains(cvt(tc.Source, sub)) {
			segm = append(segm, tc)
		}
	}
	return segm
}
