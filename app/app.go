package app

import (
	"context"
	accountHandler "ekolo/account/handler"
	account "ekolo/account/service"
	"ekolo/app/config"
	generic "ekolo/pkg/echogeneric"
	"ekolo/pkg/storage"
	"ekolo/pkg/xlog"
	tag "ekolo/tag/service"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "ekolo/app/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

type App struct {
	Opts config.Config
}

func New() App {
	opts := config.New()
	return App{Opts: opts}
}

// @title Ekolo Swagger UI
// @version 1.0
// @description Ekolo
// @contact.name API Support
// @termsOfService demo.com
// @contact.url http://demo.com/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @BasePath /
// @Schemes http https
// @query.collection.format multi
// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func (a App) Run() {
	store, err := storage.NewStore(a.Opts.GetDBDSN())
	if err != nil {
		xlog.Error("error while initializing storage", "err", err)
		return
	}
	e := echo.New()
	e.Debug = true
	// e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogError:     true,
		LogRemoteIP:  true,
		LogMethod:    true,
		LogURIPath:   true,
		LogRoutePath: true,
		LogHost:      true,
		LogProtocol:  true,
		HandleError:  true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			xlog.Info("request", "values", values)
			return nil
		},
	}))

	// Shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello !")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Organization CRUD endpoints
	generic.MountService(e, account.New(store))
	// User CRUD endpoints
	generic.MountService(e, account.NewUserService(store))
	// User extra endpoints
	userH := accountHandler.NewUserHandler(store)
	e.GET("/user/types", userH.GetUserTypes(ctx))
	// Tag CRUD endpoints
	generic.MountService(e, tag.New(store))

	// Run migrations
	models := []any{}
	models = append(models, account.GetModels()...)
	models = append(models, tag.GetModels()...)
	store.RunMigrations(models...)

	xlog.Debug("routes", "values", e.Routes())

	// Start server
	go func() {
		if err := e.Start(a.Opts.HTTPAddr); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
