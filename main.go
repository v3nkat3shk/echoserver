package main

import (
	"example/echoserver/database"
	"example/echoserver/server"
)

func main() {
	db := database.InitDb()
	echoApp := server.CreateNewEchoAPP(db)
	echoApp.App.HideBanner = true
	echoApp.Start()
}
