package main

import (
	"github.com/inact25/movieplusplus/configs"
	"github.com/inact25/movieplusplus/masters/apis"
	utils "github.com/inact25/movieplusplus/utilities"
)

func main() {
	conf := configs.NewAppConfig()
	db, err := configs.InitDB(conf)
	utils.ErrorCheck(err, "Print")
	myRoute := configs.CreateRouter()
	apis.Init(myRoute, db)
	configs.RunServer(myRoute)
}
