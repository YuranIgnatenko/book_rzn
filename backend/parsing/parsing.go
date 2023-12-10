package parsing

import (
	"backend/config"
	"backend/connector"
	"backend/models"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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
			"https://shop.prosv.ru/homepage-categorynewproducts-185",
			"https://shop.prosv.ru/homepage-categorynewproducts-185?pagenumber=2",
			"https://shop.prosv.ru/homepage-categorynewproducts-185?pagenumber=3",
			"https://shop.prosv.ru/homepage-categorynewproducts-185?pagenumber=4",
			"https://shop.prosv.ru/homepage-categorynewproducts-185?pagenumber=5",
			"https://shop.prosv.ru/homepage-categorynewproducts-185?pagenumber=6"}, "book_new"),

		NewServiceProsv([]string{
			"https://shop.prosv.ru/srednee-specialnoe-obrazovanie4415",
			"https://shop.prosv.ru/srednee-specialnoe-obrazovanie4415?pagenumber=2",
			"https://shop.prosv.ru/srednee-specialnoe-obrazovanie4415?pagenumber=3",
			"https://shop.prosv.ru/srednee-specialnoe-obrazovanie4415?pagenumber=4",
			"https://shop.prosv.ru/srednee-specialnoe-obrazovanie4415?pagenumber=5",
			"https://shop.prosv.ru/srednee-specialnoe-obrazovanie4415?pagenumber=6"}, "book_sh_middle"),

		NewServiceProsv([]string{
			"https://shop.prosv.ru/doshkolnoe-obrazovanie105",
			"https://shop.prosv.ru/doshkolnoe-obrazovanie105?pagenumber=2",
			"https://shop.prosv.ru/doshkolnoe-obrazovanie105?pagenumber=3",
			"https://shop.prosv.ru/doshkolnoe-obrazovanie105?pagenumber=4",
			"https://shop.prosv.ru/doshkolnoe-obrazovanie105?pagenumber=5",
			"https://shop.prosv.ru/doshkolnoe-obrazovanie105?pagenumber=6"}, "book_do_sh"),

		NewServiceProsv([]string{
			"https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101",
			"https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumber=2",
			"https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumber=3",
			"https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumber=4",
			"https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumber=5",
			"https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumber=6"}, "book_1_4"),

		NewServiceProsv([]string{
			"https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101",
			"https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumaber=2",
			"https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumaber=3",
			"https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumaber=4",
			"https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumaber=5",
			"https://shop.prosv.ru/nachalnoe-obrazovanie-1-4-klassy101?pagenumaber=6"}, "book_5_9"),

		NewServiceProsv([]string{
			"https://shop.prosv.ru/srednee-obrazovanie-10-11-klassy91?pagenumber=2",
			"https://shop.prosv.ru/srednee-obrazovanie-10-11-klassy91?pagenumber=3",
			"https://shop.prosv.ru/srednee-obrazovanie-10-11-klassy91?pagenumber=4",
			"https://shop.prosv.ru/srednee-obrazovanie-10-11-klassy91?pagenumber=5",
			"https://shop.prosv.ru/srednee-obrazovanie-10-11-klassy91?pagenumber=6",
			"https://shop.prosv.ru/srednee-obrazovanie-10-11-klassy91"}, "book_10_11"),

		NewServiceProsv([]string{
			"https://shop.prosv.ru/obuchenie-detej-s-ovz102?pagenumber=2",
			"https://shop.prosv.ru/obuchenie-detej-s-ovz102?pagenumber=3",
			"https://shop.prosv.ru/obuchenie-detej-s-ovz102?pagenumber=4",
			"https://shop.prosv.ru/obuchenie-detej-s-ovz102?pagenumber=5",
			"https://shop.prosv.ru/obuchenie-detej-s-ovz102?pagenumber=6",
			"https://shop.prosv.ru/obuchenie-detej-s-ovz102"}, "book_ovz"),

		NewServiceProsv([]string{
			"https://shop.prosv.ru/xudozhestvennaya-literatura103?pagenumber=2",
			"https://shop.prosv.ru/xudozhestvennaya-literatura103?pagenumber=3",
			"https://shop.prosv.ru/xudozhestvennaya-literatura103?pagenumber=4",
			"https://shop.prosv.ru/xudozhestvennaya-literatura103?pagenumber=5",
			"https://shop.prosv.ru/xudozhestvennaya-literatura103?pagenumber=6",
			"https://shop.prosv.ru/xudozhestvennaya-literatura103"}, "book_artistic"),

		NewServiceProsv([]string{
			"https://shop.prosv.ru/elektronnye-knigi182?pagenumber=2",
			"https://shop.prosv.ru/elektronnye-knigi182?pagenumber=3",
			"https://shop.prosv.ru/elektronnye-knigi182?pagenumber=4",
			"https://shop.prosv.ru/elektronnye-knigi182?pagenumber=5",
			"https://shop.prosv.ru/elektronnye-knigi182?pagenumber=6",
			"https://shop.prosv.ru/elektronnye-knigi182"}, "book_digit_books"),

		NewServiceAgat([]string{"https://agatmk.ru/moduli-sistemyi-xraneniya-fresh "}, "new_basic"),
		NewServiceAgat([]string{"https://agatmk.ru/stolyi-rabochie-fresh "}, "new_table"),
		NewServiceAgat([]string{"https://agatmk.ru/sistema-xraneniya "}, "new_boxing"),
		NewServiceAgat([]string{"https://agatmk.ru/uchenicheskie-stolyi "}, "sh_table"),
		NewServiceAgat([]string{"https://agatmk.ru/uchenicheskie-stulya "}, "sh_chair"),
		NewServiceAgat([]string{"https://agatmk.ru/stolyi-rabochie-optima "}, "office_table"),
		NewServiceAgat([]string{"https://agatmk.ru/sistema-xraneniya-optima "}, "office_boxing"),
		NewServiceAgat([]string{"https://agatmk.ru/tumbyi-pod-dosku "}, "sh_minitable"),

		NewServiceStronikum([]string{"https://stronikum.ru"}, "str_top"),
		NewServiceStronikum([]string{"https://stronikum.ru/5940_Kabinet_psihologa"}, "str_psiholog"),
		NewServiceStronikum([]string{"https://stronikum.ru/4409_Vtoraya_mladshaya_gruppa_3_4"}, "str_do_sh_3_4"),
		NewServiceStronikum([]string{"https://stronikum.ru/4442_Srednyaya_gruppa_4_5"}, "str_do_sh_4_5"),
		NewServiceStronikum([]string{"https://stronikum.ru/4472_Starshaya_gruppa_5_6"}, "str_do_sh_5_6"),
		NewServiceStronikum([]string{"https://stronikum.ru/4504_Podgotovitelnaya_gruppa_6_7"}, "str_do_sh_6_7"),
		NewServiceStronikum([]string{"https://stronikum.ru/4273_Nachalnaya_shkola"}, "str_sh_started"),
		NewServiceStronikum([]string{"https://stronikum.ru/1061_Fizika"}, "str_phisic"),
		NewServiceStronikum([]string{"https://stronikum.ru/1383_Himiya"}, "str_himiya"),
		NewServiceStronikum([]string{"https://stronikum.ru/1627_Biologiya"}, "str_biologiya"),
		NewServiceStronikum([]string{"https://stronikum.ru/2438_Literatura"}, "str_litra"),
		NewServiceStronikum([]string{"https://stronikum.ru/2486_Russkiy_yazik"}, "str_ru_lang"),
		NewServiceStronikum([]string{"https://stronikum.ru/2517_Inostranniy_yazik"}, "str_other_lang"),
		NewServiceStronikum([]string{"https://stronikum.ru/2554_Istoriya"}, "str_history"),
		NewServiceStronikum([]string{"https://stronikum.ru/2691_Geografiya"}, "str_geograph"),
		NewServiceStronikum([]string{"https://stronikum.ru/3334_Matematika"}, "str_math"),
		NewServiceStronikum([]string{"https://stronikum.ru/3531_Informatika"}, "str_info"),
		NewServiceStronikum([]string{"https://stronikum.ru/4029_OBG"}, "str_obg"),
		NewServiceStronikum([]string{"https://stronikum.ru/4133_Ekologiya"}, "str_eco"),
		NewServiceStronikum([]string{"https://stronikum.ru/3480_Izobrazitelnoe_iskusstvo"}, "str_izo"),
		NewServiceStronikum([]string{"https://stronikum.ru/3535_Muzika"}, "str_music"),
		NewServiceStronikum([]string{"https://stronikum.ru/3549_Tehnologiya"}, "str_tehno"),
		NewServiceStronikum([]string{"https://stronikum.ru/3650_Plakati_dlya_PROFOBRAZOVANIYA"}, "str_posters"),
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

	for n, service := range data {
		tc = append(tc, service.ScrapSource()...)
		fmt.Println("Parse service start --> link n:", n, "all count --> ", len(tc))
	}
	return tc
}

func htmlToString(link string) string {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func getHTMLContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
