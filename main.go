package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Aldiwildan77/ko-fi-webhook/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	routes(e)

	go func() {
		if err := e.Start(":" + fmt.Sprint(config.Cfg.Port)); err != nil {
			e.Logger.Info("Shutting down the server.")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func routes(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Iam fine thanks")
	})

	g := e.Group("/api")

	r := g.Group("/webhook")
	r.POST("", webhook)
}
