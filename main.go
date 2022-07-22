package main

import (
	"deviceApp/model/orm"
	"deviceApp/router"
)

func main() {

	orm.InitDB()
	router.InitRouter()

}
