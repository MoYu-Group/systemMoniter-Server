package models

import (
	"github.com/spf13/viper"
)

var (
	Interval  = 1
	Cu        = "120.52.99.224"
	Ct        = "183.78.182.66"
	Cm        = "211.139.145.129"
	Porbeport = 80
	IsOpen    = true
)

type Config struct {
	interval  int
	cu        string
	ct        string
	cm        string
	isOpen    bool
	porbeport int
}

func NewConfig() Config {
	return Config{
		interval:  Interval,
		cu:        Cu,
		ct:        Ct,
		cm:        Cm,
		isOpen:    IsOpen,
		porbeport: Porbeport,
	}
}

func SetConfig() Config {
	interval := viper.GetInt("network.interval")
	isOpen := false
	if interval > 0 {
		isOpen = true
	} else {
		isOpen = false
	}
	return Config{
		interval: interval,
		cu:       viper.GetString("network.cu"),
		ct:       viper.GetString("network.ct"),
		cm:       viper.GetString("network.cm"),
		isOpen:   isOpen,
	}
}
