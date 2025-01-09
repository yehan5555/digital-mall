package main

import (
	"test_mysql/conf"
	"test_mysql/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
