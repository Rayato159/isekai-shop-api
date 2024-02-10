package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Rayato159/isekai-shop-api/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type echoServer struct {
	app        *echo.Echo
	baseRouter *echo.Group
	db         *gorm.DB
	conf       *config.AppConfig
}

var (
	server *echoServer
	once   sync.Once
)

func NewEchoServer(conf *config.AppConfig, db *gorm.DB) Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	baseRouter := echoApp.Group("/v1")

	once.Do(func() {
		server = &echoServer{
			app:        echoApp,
			baseRouter: baseRouter,
			db:         db,
			conf:       conf,
		}
	})

	return server
}

func (s *echoServer) Start() {
	// Initialize all middlewares
	timeOutMiddleware := getTimeOutMiddleware(s.conf.ServerConfig.Timeout)
	corsMiddleware := getCorsMiddleware(s.conf.ServerConfig.AllowOrigins)
	bodyLimitMiddleware := getBodyLimitMiddleware(s.conf.ServerConfig.BodyLimit)

	s.app.Use(timeOutMiddleware)
	s.app.Use(corsMiddleware)
	s.app.Use(bodyLimitMiddleware)

	// Initialzie all routers
	s.baseRouter.GET("/health", s.healthCheck)

	s.initOauth2Router()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefulShutdown(quit)

	s.httpListening()
}

func (s *echoServer) httpListening() {
	serverUrl := fmt.Sprintf(":%d", s.conf.ServerConfig.Port)

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

func (s *echoServer) healthCheck(pctx echo.Context) error {
	return pctx.String(http.StatusOK, "OK")
}
