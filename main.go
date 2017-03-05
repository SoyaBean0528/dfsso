package main

import (
	"flag"

	_ "dreamfish/dfsso/routers"
	"dreamfish/dfsso/models/initDBModel"

	"github.com/astaxie/beego"
)

func main() {
	// flags
	initdb := flag.Bool("initdb", false, "Should initdb.")
	flag.Parse()

	// database
	if *initdb {
		initDBModel.InitDB()
		return
	}
	initDBModel.RegisterDB()

	// start
	beego.Run()
}
