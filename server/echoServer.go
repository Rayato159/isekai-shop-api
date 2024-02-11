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
	_oauth2Repository "github.com/Rayato159/isekai-shop-api/modules/oauth2/repository"
	_oauth2Service "github.com/Rayato159/isekai-shop-api/modules/oauth2/service"
	_playerRepository "github.com/Rayato159/isekai-shop-api/modules/player/repository"
	"github.com/Rayato159/isekai-shop-api/server/customMiddlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// Initialize all custom middlewares
	middlewares := s.getMiddlewares()
	_ = middlewares

	// Initialzie all routers
	s.baseRouter.GET("/health", s.healthCheck)

	s.initOAuth2Router()

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
	fmt.Println(pctx.Cookie("_authorization"))
	return pctx.String(http.StatusOK, "OK")
}

func (s *echoServer) getMiddlewares() customMiddlewares.CustomMiddleware {
	oauth2Repository := _oauth2Repository.NewGoogleOAuth2Repository(s.db, s.app.Logger)
	playerRepository := _playerRepository.NewPlayerRepositoryImpl(s.db, s.app.Logger)

	oauth2Service := _oauth2Service.NewGoogleOAuth2Service(
		oauth2Repository,
		playerRepository,
		s.app.Logger,
	)

	return customMiddlewares.NewCustomMiddlewaresImpl(
		s.conf.OAuth2Config,
		s.app.Logger,
		oauth2Service,
	)
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
