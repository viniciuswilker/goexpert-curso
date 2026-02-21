package main

import (
	"1-api/configs"
)

func main() {
	config, _ := configs.CarregarConfig(".")
	println(config.DBDriver)

}
