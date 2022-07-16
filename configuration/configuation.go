package configuration

import "github.com/BurntSushi/toml"

type Config struct {
	Title string `toml:"title"`
	Base  base   `toml:"base"`
}

type base struct {
	Port      int32  `toml:"port"`
	ImagePath string `toml:"image_path"`
}

func Configuarion() Config {
	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}
	return config
}
