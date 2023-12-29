package main

import (
	"ginblog-be/model"
	"ginblog-be/router"
)

func main() {
	model.InitDb()
	router.InitRouter()
}
