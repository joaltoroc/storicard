package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	"github/joaltoroc/storicard/config"
	"github/joaltoroc/storicard/pkg/database"
	customMiddleware "github/joaltoroc/storicard/pkg/middleware"
)

type App struct {
	db    *gorm.DB
	cfg   config.Config
	echo  *echo.Echo
	local bool
}

func NewApp(ctx context.Context) *App {
	env := os.Getenv("env")
	if env == "" {
		env = "local"
	}

	cfg, err := config.LoadConfig(env)
	if err != nil {
		panic(err)
	}

	db, err := database.NewDatabase(cfg)
	if err != nil {
		panic(err)
	}

	return &App{
		cfg:   cfg,
		db:    db,
		echo:  echo.New(),
		local: env == "local",
	}
}

func (app *App) Run() error {
	if err := app.startService(); err != nil {
		return err
	}

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)

	go func() {
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Close DB Connection
		dbInstance, _ := app.db.DB()
		_ = dbInstance.Close()

		app.echo.Shutdown(ctx)
	}()

	app.echo.Debug = app.cfg.Server.Debug
	app.echo.Use(customMiddleware.AppCors())
	app.echo.Use(customMiddleware.CacheWithRevalidation)

	app.echo.Use(middleware.RequestID())
	app.echo.Use(middleware.Secure())

	if app.local {
		app.echo.Use(middleware.Logger())
	}

	return app.echo.Start(fmt.Sprintf(":%s", app.cfg.Server.Port))
}

func PingPong(domain *echo.Group) {
	domain.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong")
	})
}
