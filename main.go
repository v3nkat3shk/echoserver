package main

import (
	"github.com/v3nkat3shk/echoserver/database"
	"github.com/v3nkat3shk/echoserver/server"
)

func main() {
	db := database.InitDb()
	echoApp := server.CreateNewEchoAPP(db)
	echoApp.App.HideBanner = true
	echoApp.Start()
}
