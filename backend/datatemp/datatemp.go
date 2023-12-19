package datatemp

import (
	"backend/config"
	"backend/models"
	"strings"
)

type DataTemp struct {
	config.Configuration
	TargetAll []models.TargetCard
	// TargetCards          []models.TargetCard
	PageTarget           models.PageTarget
	MenuCards            []models.MenuCard
	ListOrdersTargetCard models.ListOrdersTargetCard
	FavoritesCards       []models.Favorites
	OrdersCards          []models.OrdersCards
	OrdersRows           []models.OrdersRows
	SearchTarget         []models.TargetCard
	FastOrdersList       []models.DataFastOrderOne
	NumberFastOrder      string
	IsLogin              bool
	IsAdmin              bool
	LastValueSearch      string
	NameLogin            string
	PageMenuTemplHtml    string
	MenuMap              []models.MenuCard
	MenuFiles            []models.MenuCard
}

func NewDataTemp(c config.Configuration, ps []models.TargetCard) *DataTemp {
	return &DataTemp{
		Configuration: c,
		TargetAll:     ps,
		PageTarget:    models.PageTarget{},
		// TargetCards:          []models.TargetCard{},
		ListOrdersTargetCard: models.ListOrdersTargetCard{},
		FavoritesCards:       []models.Favorites{},
		OrdersCards:          []models.OrdersCards{},
		OrdersRows:           []models.OrdersRows{},
		SearchTarget:         []models.TargetCard{},
		FastOrdersList:       []models.DataFastOrderOne{},
		NumberFastOrder:      "",
		LastValueSearch:      "",
		IsLogin:              false,
		IsAdmin:              false,
		NameLogin:            "Гость",
		MenuFiles: []models.MenuCard{
			models.MenuCard{
				Title:  "Приказ 804",
				Link:   "static/logo_docs.png",
				PathTo: "home_804",
			},
			models.MenuCard{
				Title:  "Приказ 804",
				Link:   "static/logo_docs.png",
				PathTo: "home_804",
			},
			models.MenuCard{
				Title:  "Приказ 804",
				Link:   "static/logo_docs.png",
				PathTo: "home_804",
			},
			models.MenuCard{
				Title:  "Приказ 804",
				Link:   "static/logo_docs.png",
				PathTo: "home_804",
			},
			models.MenuCard{
				Title:  "Приказ 804",
				Link:   "static/logo_docs.png",
				PathTo: "home_804",
			},
			models.MenuCard{
				Title:  "Приказ 804",
				Link:   "static/logo_docs.png",
				PathTo: "home_804",
			},
			models.MenuCard{
				Title:  "Приказ 804",
				Link:   "static/logo_docs.png",
				PathTo: "home_804",
			},
			models.MenuCard{
				Title:  "Приказ 804",
				Link:   "static/logo_docs.png",
				PathTo: "home_804",
			},
			models.MenuCard{
				Title:  "Приказ 804",
				Link:   "static/logo_docs.png",
				PathTo: "home_804",
			},
			models.MenuCard{
				Title:  "Приказ 804",
				Link:   "static/logo_docs.png",
				PathTo: "home_804",
			},
		},
		MenuMap: []models.MenuCard{
			models.MenuCard{
				Title:  "Новинки",
				Link:   "static/logo_news.jpg",
				PathTo: "home_news",
			},
			models.MenuCard{
				Title:  "Приказ 804 (pdf)",
				Link:   "static/logo_docs.png",
				PathTo: "home_804",
			},
			models.MenuCard{
				Title:  "Посмотреть нас в VK",
				Link:   "static/logo_vk_group.png",
				PathTo: "home_vk",
			},
			models.MenuCard{
				Title:  "Контакты и Адреса",
				Link:   "static/logo_address.png",
				PathTo: "home_contacts_address",
			},

			models.MenuCard{
				Title:  "Документы и информация",
				Link:   "static/logo_info.png",
				PathTo: "home_docs_info",
			},
		},
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
					background-color: #3584a4;
		            color:  rgb(255,255,255);
					
		        }mainContrastSchemeColor
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
					
						<li><a href="/home" class="submenu-link">Главная страница</a>
						</li>

						<li><a href="/new_basic" class="submenu-link">Мебельные новинки</a>
							<ul class="submenu">
								<li><a href="/new_basic">Базовые модули</a></li>
								<li><a href="/new_boxing">Системы хранения</a></li>
							</ul>
						</li>

						<li><a href="/sh_table" class="submenu-link">Мебель для школ</a>
							<ul class="submenu">
								<li><a href="/sh_table">Рабочие столы</a></li>
								<li><a href="/sh_chair">Рабочие стулья</a></li>
								<li><a href="/sh_minitable">Тумба под доску</a></li>
							</ul>
						</li>

						<li><a href="/office_table" class="submenu-link">Мебель в офис</a>
							<ul class="submenu">
								<li><a href="/office_table">Рабочие столы</a></li>
								<li><a href="/office_boxing">Системы хранения</a></li>
							</ul>
						</li>

						<li><a href="/book_1_4" class="submenu-link">Книги и Учебники</a>
							<ul class="submenu">
								<li><a href="/book_new">Новинки</a></li>
								<li><a href="/book_sh_middle">Для среднего специального образования</a></li>
								<li><a href="/book_do_sh">Для дошкольников</a></li>
								<li><a href="/book_1_4">Для 1-4 классов</a></li>
								<li><a href="/book_5_9">Для 5-9 классов</a></li>
								<li><a href="/book_10_11">Для 10-11 классов</a></li>
								<li><a href="/book_ovz">Для детей с ОВЗ</a></li>
								<li><a href="/book_artistic">Художественная литература</a></li>
							</ul>
						</li>

						<li><a href="/book_digit_books" class="submenu-link">Электронная библиотека</a>
						</li>

						<li><a href="/str_do_sh_3_4" class="submenu-link">Оборудование Дошкольное</a>
							<ul class="submenu">
								<li><a href="/str_do_sh_3_4">Дошкольники 3-4 лет</a></li>
								<li><a href="/str_do_sh_4_5">Дошкольники 4-5 лет</a></li>
								<li><a href="/str_do_sh_5_6">Дошкольники 5-6 лет</a></li>
								<li><a href="/str_do_sh_6_7">Дошкольники 6-7 лет</a></li>
								<li><a href="/str_sh_started">Начальная школа</a></li>
							</ul>
						</li>

						<li><a href="/str_psiholog" class="submenu-link">Оборудование Предметное</a>
							<ul class="submenu">
								<li><a href="/str_psiholog">Психология</a></li>
								<li><a href="/str_phisic">Физика</a></li>
								<li><a href="/str_himiya">Химия</a></li>
								<li><a href="/str_biologiya">Биология</a></li>
								<li><a href="/str_litra">Литература</a></li>
								<li><a href="/str_ru_lang">Русский язык</a></li>
								<li><a href="/str_other_lang">Иностранный язык</a></li>
								<li><a href="/str_history">История</a></li>
							</ul>
						</li>
		
						<li><a href="/str_math" class="submenu-link">Оборудование Дополнительное</a>
							<ul class="submenu">
								<li><a href="/str_geograph">География</a></li>
								<li><a href="/str_math">Математика</a></li>
								<li><a href="/str_info">Информатика</a></li>
								<li><a href="/str_obg">ОБЖ</a></li>
								<li><a href="/str_eco">Экология</a></li>
								<li><a href="/str_izo">Изобразительное искусство</a></li>
								<li><a href="/str_music">Музыка</a></li>
								<li><a href="/str_tehno">Технология</a></li>
							</ul>
						</li>

						<li><a href="/str_posters" class="submenu-link">Плакаты для ПРОФ.образования</a>
						</li>

						<li><a href="/naura" class="submenu-link">naura</a>
						</li>

					</ul>
				</nav>
				`}
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

// func (dt *DataTemp) GetMenuFromMap(path_menu string) []models.MenuCard {
// 	segm := make([]models.MenuCard, 0)

// 	cvt := func(sub1, sub2 string) (string, string) {
// 		return strings.ToLower(sub1), strings.ToLower(sub2)
// 	}

// 	for _, tc := range data {
// 		if strings.Contains(cvt(tc.Autor, sub)) || strings.Contains(cvt(tc.Price, sub)) ||
// 			strings.Contains(cvt(tc.Title, sub)) || strings.Contains(cvt(tc.Source, sub)) {
// 			segm = append(segm, tc)
// 		}
// 	}
// 	return segm
// }
