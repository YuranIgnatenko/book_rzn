package config

import (
	"encoding/json"
	"io"
	"os"
)

var path_config = "/home/yu/Desktop/code/book_rzn/backend/config/config.json"

// var path_config_default = "/home/yu/Desktop/code/book_rzn/backend/config/default.json"

type Configuration struct {
	Url_prefix         string   `json:"url_prefix"`
	Ip                 string   `json:"ip"`
	Port               string   `json:"port"`
	Split_ip_port      string   `json:"split_ip_port"`
	Path_prefix        string   `json:"path_prefix"`
	Path_frontend      string   `json:"path_frontend"`
	Path_backend       string   `json:"path_backend"`
	Path_config        string   `json:"path_config"`
	Path_static        string   `json:"path_static"`
	Path_bd            string   `json:"path_bd"`
	Bd_admin_list      string   `json:"bd_admin_list"`
	Bd_users_list      string   `json:"bd_users_list"`
	Bd_favorites       string   `json:"bd_favorites"`
	Bd_orders          string   `json:"bd_orders"`
	Bd_prosv           string   `json:"bd_prosv"`
	Phone              string   `json:"phone"`
	Email              string   `json:"email"`
	Footer_years       string   `json:"footer_years"`
	Copyright          string   `json:"copyright"`
	Location_office    string   `json:"location_office"`
	Location_shops     []string `json:"location_shops"`
	Path_logo_company  string   `json:"path_logo_company"`
	Path_banner_home   string   `json:"path_banner_home"`
	Path_banner_home_2 string   `json:"path_banner_home_2"`
	Path_banner_home_3 string   `json:"path_banner_home_3"`
	Company_name       string   `json:"company_name"`
	Text_banner        string   `json:"text_banner"`
	Description_banner string   `json:"description_banner"`
	DB_user            string   `json:"db_user"`
	DB_password        string   `json:"db_password"`
	DB_hostname        string   `json:"db_hostname"`
	DB_database        string   `json:"db_database"`
}

func NewConfiguration() *Configuration {
	file, _ := os.Open(path_config)
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	var conf Configuration
	json.Unmarshal(byteValue, &conf)
	conf.DB_user = "bookrzn"
	conf.DB_password = "password"
	return &conf
}
