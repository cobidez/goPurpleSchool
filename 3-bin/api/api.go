package api

import (
	"bin/config"
)

func getKey() string {
	conf := config.NewConfig()
	return conf.GetKey()
}
