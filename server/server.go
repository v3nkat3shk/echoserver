package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/v3nkat3shk/echoserver/api"
	"gorm.io/gorm"
)

type App interface {
	Start()
}

type EchoApp struct {
	App echo.Echo
	DB  *gorm.DB
}

func CreateNewEchoAPP(db *gorm.DB) EchoApp {
	return EchoApp{
		App: *echo.New(),
		DB:  db,
	}
}

func (e *EchoApp) Start() {
	e.initializeApplication()
	e.graceUp("8080")
}

func (e *EchoApp) initializeApplication() {
	apiV1 := e.App.Group("/api/v1")

	e.App.Pre(middleware.RemoveTrailingSlash())
	apiV1.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${host}${route} ${latency_human}\n",
	}))
	apiV1.Use(middleware.Recover())

	apiV1.GET("/check", api.HelloWorld)
	apiV1.GET("/users", api.GetUsers)
	apiV1.POST("/user", api.CreateUser)

}

func (e *EchoApp) graceUp(port string) {
	go func() {
		if err := e.App.Start(":" + port); err != nil && err != http.ErrServerClosed {
			e.App.Logger.Fatal("shutting down the interface")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.App.Shutdown(ctx); err != nil {
		e.App.Logger.Fatal(err)
	}
}
