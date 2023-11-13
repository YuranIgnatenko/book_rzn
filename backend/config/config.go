package config

import (
	"backend/parsing"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var path_config = "/home/yu/Рабочий стол/code/book_rzn/backend/config/config.json"
var path_datatemp = "/home/yu/Рабочий стол/code/book_rzn/backend/config/datatemp.json"
var path_datatempCMS = "/home/yu/Рабочий стол/code/book_rzn/backend/config/cms.json"

type Configuration struct {
	Url_prefix    string `json:"url_prefix"`
	Ip            string `json:"ip"`
	Port          string `json:"port"`
	Split_ip_port string `json:"split_ip_port"`
	Path_prefix   string `json:"path_prefix"`
	Path_frontend string `json:"path_frontend"`
	Path_backend  string `json:"path_backend"`
	Path_config   string `json:"path_config"`
	Path_static   string `json:"path_static"`
	Path_bd       string `json:"path_bd"`
	Bd_admin_list string `json:"bd_admin_list"`
	Bd_users_list string `json:"bd_users_list"`
	Full_url_addr string
	HostAndPort   string
}

type DataTemp struct {
	Phone              string   `json:"phone"`
	Email              string   `json:"email"`
	Footer_years       string   `json:"footer_years"`
	Copyright          string   `json:"copyright"`
	Location_office    string   `json:"location_office"`
	Location_shops     []string `json:"location_shops"`
	Path_logo_company  string   `json:"path_logo_company"`
	Path_banner_home   string   `json:"path_banner_home"`
	Full_url_addr      string   `json:"full_url_addr"`
	Company_name       string   `json:"compny_name"`
	Text_banner        string   `json:"text_banner"`
	Description_banner string   `json:"description_banner"`

	Cards_prosv     []string
	Cards_naura     []string
	Cards_stronikum []string
	Cards_agat      []string

	Prosv_cards []parsing.ProsvCard
}

type Cms struct {
	Url_prefix        string   `json:"url_prefix"`
	Ip                string   `json:"ip"`
	Port              string   `json:"port"`
	Split_ip_port     string   `json:"split_ip_port"`
	Path_prefix       string   `json:"path_prefix"`
	Path_frontend     string   `json:"path_frontend"`
	Path_backend      string   `json:"path_backend"`
	Path_config       string   `json:"path_config"`
	Path_static       string   `json:"path_static"`
	Path_bd           string   `json:"path_bd"`
	Bd_admin_list     string   `json:"bd_admin_list"`
	Bd_users_list     string   `json:"bd_users_list"`
	Phone             string   `json:"phone"`
	Email             string   `json:"email"`
	Footer_years      string   `json:"footer_years"`
	Copyright         string   `json:"copyright"`
	Location_office   string   `json:"location_office"`
	Location_shops    []string `json:"location_shops"`
	Path_logo_company string   `json:"path_logo_company"`
	Path_banner_home  string   `json:"path_banner_home"`
	Company_name      string   `json:"compny_name"`
	Full_url_addr     string
	HostAndPort       string
}

func NewConfiguration() Configuration {
	file, _ := os.Open(path_config)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	var conf Configuration
	json.Unmarshal(byteValue, &conf)
	conf.Full_url_addr = conf.Url_prefix + conf.Ip + conf.Split_ip_port + conf.Port
	conf.HostAndPort = conf.Ip + conf.Split_ip_port + conf.Port
	return conf
}

func NewDataTemp() DataTemp {
	file, _ := os.Open(path_datatemp)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	var dt DataTemp
	json.Unmarshal(byteValue, &dt)
	// dt.Prosv_cards = []ProsvCard{
	// 	ProsvCard{"name", "desc", "100.90 r"},
	// 	ProsvCard{"name", "desc", "100.90 r"},
	// 	ProsvCard{"name", "desc", "100.90 r"},
	// 	ProsvCard{"name", "desc", "100.90 r"},
	// }

	dt.Prosv_cards = parsing.GetLinks()
	return dt
}

func NewCms() Cms {
	file, _ := os.Open(path_datatempCMS)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	var dtCMS Cms
	json.Unmarshal(byteValue, &dtCMS)
	dtCMS.Full_url_addr = dtCMS.Url_prefix + dtCMS.Ip + dtCMS.Split_ip_port + dtCMS.Port
	dtCMS.HostAndPort = dtCMS.Ip + dtCMS.Split_ip_port + dtCMS.Port
	fmt.Printf("%#v\n", dtCMS)
	return dtCMS
}
