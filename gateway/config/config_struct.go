package config

type Config struct {
	Gateway struct {
		Routes []Route `mapstructure:"routes"`
	} `mapstructure:"gateway"`
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
}

type Route struct {
	ID         string   `mapstructure:"id"`
	URI        string   `mapstructure:"uri"`
	Target     string   `mapstructure:"target"`
	Predicates []string `mapstructure:"predicates"`
	Filters    []string `mapstructure:"filters"`
	Metadata   Metadata `mapstructure:"metadata"`
}

type Metadata struct {
	LB LBConfig `mapstructure:"lb"`
}

type LBConfig struct {
	Name string `mapstructure:"name"`
}
