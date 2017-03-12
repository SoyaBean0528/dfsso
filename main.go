package main

import (
	"flag"

	"dreamfish/dfsso/models/initDBModel"
	_ "dreamfish/dfsso/routers"

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
