package main

import (
	"CloudHoliday/conf"
	"CloudHoliday/routers"
)

var port = conf.HttpPort

func main() {
	router := routers.InitRouter()
	router.Run(port)
}