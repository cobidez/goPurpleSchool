package api

import (
	"bin/config"
)

func getKey() string {
	return config.NewConfig().Key
}
