package main

import (
	"github.com/CHIIIPLOW-922/ginblog-be/dbconfig"
	"github.com/CHIIIPLOW-922/ginblog-be/router"
)

func main() {

	dbconfig.InitDB()

	router.InitRouter()
}
