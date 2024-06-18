package config

type Application struct {
	Server     Server     `yaml:"server"`
	DataSource DataSource `yaml:"datasource"`
}
