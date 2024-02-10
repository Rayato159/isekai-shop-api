package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Rayato159/isekai-shop-api/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
	app *echo.Echo
	db  *gorm.DB
	cfg *config.AppConfig
}

var (
	appServer *echoServer
	once      sync.Once
)

func NewAppServer(cfg *config.AppConfig, db *gorm.DB) Server {
	once.Do(func() {
		appServer = &echoServer{
			app: echo.New(),
			db:  db,
			cfg: cfg,
		}
	})

	return appServer
}

func (s *echoServer) Start() {
	// Initialize all middlewares
	timeOutMiddleware := getTimeOutMiddleware(s.cfg.ServerConfig.Timeout)
	corsMiddleware := getCorsMiddleware(s.cfg.ServerConfig.AllowOrigins)
	bodyLimitMiddleware := getBodyLimitMiddleware(s.cfg.ServerConfig.BodyLimit)

	s.app.Use(timeOutMiddleware)
	s.app.Use(corsMiddleware)
	s.app.Use(bodyLimitMiddleware)

	// Initialzie all routers

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefulShutdown(quit)

	s.httpListening()

	s.app.Use(middleware.Logger())
}

func (s *echoServer) httpListening() {
	serverUrl := fmt.Sprintf(":%d", s.cfg.ServerConfig.Port)

	if err := s.app.Start(serverUrl); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("Error: %v", err)
	}
}

func (s *echoServer) gracefulShutdown(quit <-chan os.Signal) {
	ctx := context.Background()

	<-quit
	s.app.Logger.Infof("Shutting down service...")

	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}

func getTimeOutMiddleware(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Error: Request timeout.",
		Timeout:      timeout * time.Second,
	})
}

func getCorsMiddleware(allowOrigins []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: allowOrigins,
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
	})
}

func getBodyLimitMiddleware(bodyLimit string) echo.MiddlewareFunc {
	return middleware.BodyLimit(bodyLimit)
}
