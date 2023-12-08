package parsing

import (
	"backend/config"
	"backend/connector"
	"backend/models"
	"fmt"
)

func TargetHash(autor, title, price, image string) string {
	return fmt.Sprintf("%v%v%v%v", autor, title, price, image)
}

type ParsingService struct {
	ListServices        []models.ServiceScraper
	ListTargetCardCache []models.TargetCard

	config.Configuration
	connector.Connector
}

func NewParsingService(c config.Configuration, conn connector.Connector) *ParsingService {
	lss := []models.ServiceScraper{
		NewServiceProsv([]string{
			// 	"https://shop.prosv.ru/homepage-categorynewproducts-185",
			// 	"https://shop.prosv.ru/homepage-categorynewproducts-185?pagenumber=2",
			// 	"https://shop.prosv.ru/homepage-categorynewproducts-185?pagenumber=3",
			// 	"https://shop.prosv.ru/homepage-categorynewproducts-185?pagenumber=4",
			// 	"https://shop.prosv.ru/homepage-categorynewproducts-185?pagenumber=5",
			"https://shop.prosv.ru/homepage-categorynewproducts-185?pagenumber=6"}, "prosv_new"),

		NewServiceProsv([]string{
			// "https://shop.prosv.ru/srednee-specialnoe-obrazovanie4415",
			// "https://shop.prosv.ru/srednee-specialnoe-obrazovanie4415?pagenumber=2",
			// "https://shop.prosv.ru/srednee-specialnoe-obrazovanie4415?pagenumber=3",
			// "https://shop.prosv.ru/srednee-specialnoe-obrazovanie4415?pagenumber=4",
			// "https://shop.prosv.ru/srednee-specialnoe-obrazovanie4415?pagenumber=5",
			"https://shop.prosv.ru/srednee-specialnoe-obrazovanie4415?pagenumber=6"}, "prosv_sh_middle"),

		NewServiceProsv([]string{
			// "https://shop.prosv.ru/doshkolnoe-obrazovanie105",
			// "https://shop.prosv.ru/doshkolnoe-obrazovanie105?pagenumber=2",
			// "https://shop.prosv.ru/doshkolnoe-obrazovanie105?pagenumber=3",
			// "https://shop.prosv.ru/doshkolnoe-obrazovanie105?pagenumber=4",
			// "https://shop.prosv.ru/doshkolnoe-obrazovanie105?pagenumber=5",
			"https://shop.prosv.ru/doshkolnoe-obrazovanie105?pagenumber=6"}, "prosv_do_sh"),

		NewServiceProsv([]string{
			// "https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101",
			// "https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumber=2",
			// "https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumber=3",
			// "https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumber=4",
			// "https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumber=5",
			"https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumber=6"}, "prosv_1_4"),

		NewServiceProsv([]string{
			// "https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101",
			// "https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumaber=2",
			// "https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumaber=3",
			// "https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumaber=4",
			// "https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumaber=5",
			"https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumaber=6"}, "prosv_5_9"),

		NewServiceProsv([]string{
			// "https://shop.prosv.ru/srednee-obrazovanie-10-11-klassy91?pagenumber=2",
			// "https://shop.prosv.ru/srednee-obrazovanie-10-11-klassy91?pagenumber=3",
			// "https://shop.prosv.ru/srednee-obrazovanie-10-11-klassy91?pagenumber=4",
			// "https://shop.prosv.ru/srednee-obrazovanie-10-11-klassy91?pagenumber=5",
			// "https://shop.prosv.ru/srednee-obrazovanie-10-11-klassy91?pagenumber=6",
			"https://shop.prosv.ru/srednee-obrazovanie-10-11-klassy91"}, "prosv_10_11"),

		NewServiceProsv([]string{
			// "https://shop.prosv.ru/obuchenie-detej-s-ovz102?pagenumber=2",
			// "https://shop.prosv.ru/obuchenie-detej-s-ovz102?pagenumber=3",
			// "https://shop.prosv.ru/obuchenie-detej-s-ovz102?pagenumber=4",
			// "https://shop.prosv.ru/obuchenie-detej-s-ovz102?pagenumber=5",
			// "https://shop.prosv.ru/obuchenie-detej-s-ovz102?pagenumber=6",
			"https://shop.prosv.ru/obuchenie-detej-s-ovz102"}, "prosv_ovz"),

		NewServiceProsv([]string{
			// "https://shop.prosv.ru/xudozhestvennaya-literatura103?pagenumber=2",
			// "https://shop.prosv.ru/xudozhestvennaya-literatura103?pagenumber=3",
			// "https://shop.prosv.ru/xudozhestvennaya-literatura103?pagenumber=4",
			// "https://shop.prosv.ru/xudozhestvennaya-literatura103?pagenumber=5",
			// "https://shop.prosv.ru/xudozhestvennaya-literatura103?pagenumber=6",
			"https://shop.prosv.ru/xudozhestvennaya-literatura103"}, "prosv_artistic"),

		NewServiceProsv([]string{
			// "https://shop.prosv.ru/elektronnye-knigi182?pagenumber=2",
			// "https://shop.prosv.ru/elektronnye-knigi182?pagenumber=3",
			// "https://shop.prosv.ru/elektronnye-knigi182?pagenumber=4",
			// "https://shop.prosv.ru/elektronnye-knigi182?pagenumber=5",
			// "https://shop.prosv.ru/elektronnye-knigi182?pagenumber=6",
			"https://shop.prosv.ru/elektronnye-knigi182"}, "prosv_digit_books"),

		NewServiceAgat([]string{"https://agatmk.ru/moduli-sistemyi-xraneniya-fresh "}, "new_basic"),
		// 	NewServiceAgat([]string{"https://agatmk.ru/stolyi-rabochie-fresh "}, "new_table"),
		// 	NewServiceAgat([]string{"https://agatmk.ru/sistema-xraneniya "}, "new_boxing"),
		// 	NewServiceAgat([]string{"https://agatmk.ru/uchenicheskie-stolyi "}, "sh_table"),
		// 	NewServiceAgat([]string{"https://agatmk.ru/uchenicheskie-stulya "}, "sh_chair"),
		// 	NewServiceAgat([]string{"https://agatmk.ru/stolyi-rabochie-optima "}, "office_table"),
		// 	NewServiceAgat([]string{"https://agatmk.ru/sistema-xraneniya-optima "}, "office_boxing"),
		// 	NewServiceAgat([]string{"https://agatmk.ru/tumbyi-pod-dosku "}, "sh_minitable"),
	}

	ps := ParsingService{
		Configuration: c,
		Connector:     conn,
		ListServices:  lss,
	}

	tc_all := conn.TableTargets.GetListTargets()

	if len(tc_all) <= 1 {
		fmt.Println("Launched scrapper - started")
		tc_all = RangeScrapServices(lss)

		for _, tc_temp := range tc_all {
			conn.TableTargets.SaveParsingService(tc_temp)
		}
		ps.ListTargetCardCache = tc_all
		fmt.Println("Launched scrapper -- finished")
	} else {
		fmt.Println("Launched scrapper -- no (getting from BD)")
		ps.ListTargetCardCache = tc_all
	}
	fmt.Println("All targets in BD :", len(ps.ListTargetCardCache))

	return &ps
}

func RangeScrapServices(data []models.ServiceScraper) []models.TargetCard {
	tc := make([]models.TargetCard, 0)

	for _, service := range data {
		tc = append(tc, service.ScrapSource()...)
	}
	return tc
}
