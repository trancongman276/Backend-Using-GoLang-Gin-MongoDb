package main

import (
	"test/routers"
	"test/utils"
)

func main() {
	router := routers.InitRoute()
	port := utils.EnvVar("localhost", ":27017")
	router.Run(port)
}
