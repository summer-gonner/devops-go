package config

import "strconv"

type DataSource struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

func (ds *DataSource) Dsn() string {
	if ds.Type == "mysql" {
		return ds.Username + ":" + ds.Password + "@tcp(" + ds.Host + ":" + strconv.Itoa(ds.Port) + ")/" + ds.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	} else {
		return ""
	}
}
