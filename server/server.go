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
	_adminRepository "github.com/Rayato159/isekai-shop-api/pkg/admin/repository"
	_oauth2Controller "github.com/Rayato159/isekai-shop-api/pkg/oauth2/controller"
	_oauth2Service "github.com/Rayato159/isekai-shop-api/pkg/oauth2/service"
	"github.com/Rayato159/isekai-shop-api/pkg/oauth2/state"
	_playerRepository "github.com/Rayato159/isekai-shop-api/pkg/player/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type echoServer struct {
	app  *echo.Echo
	db   *gorm.DB
	conf *config.Config
}

var (
	server *echoServer
	once   sync.Once
)

func NewEchoServer(conf *config.Config, db *gorm.DB) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	once.Do(func() {
		server = &echoServer{
			app:  echoApp,
			db:   db,
			conf: conf,
		}
	})

	return server
}

func (s *echoServer) Start() {
	// Initialize all middlewares
	loggerMiddleware := getLoggerMiddleware(s.app.Logger)
	timeOutMiddleware := getTimeOutMiddleware(s.conf.Server.Timeout)
	corsMiddleware := getCorsMiddleware(s.conf.Server.AllowOrigins)
	bodyLimitMiddleware := getBodyLimitMiddleware(s.conf.Server.BodyLimit)

	// Prevent application from crashing
	s.app.Use(middleware.Recover())

	s.app.Use(loggerMiddleware)
	s.app.Use(timeOutMiddleware)
	s.app.Use(corsMiddleware)
	s.app.Use(bodyLimitMiddleware)

	// Initialize all custom middlewares
	customerMiddleware := s.getCustomMiddleware()

	// Initialzie all routers
	s.app.GET("/v1/health", s.healthCheck)

	s.initOAuth2Router()
	s.initInventoryRouter(customerMiddleware)
	s.initItemShopRouter(customerMiddleware)
	s.initItemManagingRouter(customerMiddleware)
	s.initPlayerCoinRouter(customerMiddleware)

	// Graceful shutdown
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefullyShutdown(quitCh)

	s.httpListening()
}

func (s *echoServer) httpListening() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("Error: %v", err)
	}
}

func (s *echoServer) gracefullyShutdown(quitCh <-chan os.Signal) {
	ctx := context.Background()

	<-quitCh
	s.app.Logger.Infof("Shutting down service...")

	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}

func (s *echoServer) healthCheck(pctx echo.Context) error {
	return pctx.String(http.StatusOK, "OK")
}

func (s *echoServer) getCustomMiddleware() *customMiddleware {
	stateConfig := s.conf.State
	stateProvider := state.NewJwtState(
		[]byte(stateConfig.Secret),
		stateConfig.ExpiresAt,
		stateConfig.Issuer,
	)

	playerRepository := _playerRepository.NewPlayerRepositoryImpl(s.db, s.app.Logger)
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db, s.app.Logger)

	oauth2Service := _oauth2Service.NewGoogleOAuth2Service(
		playerRepository,
		adminRepository,
	)

	oauth2Controller := _oauth2Controller.NewGoogleOAuth2Controller(
		oauth2Service,
		s.conf.OAuth2,
		stateProvider,
		s.app.Logger,
	)

	return &customMiddleware{
		oauth2Controller: oauth2Controller,
		oauth2Conf:       s.conf.OAuth2,
		logger:           s.app.Logger,
	}
}

func getLoggerMiddleware(logger echo.Logger) echo.MiddlewareFunc {
	return middleware.Logger()
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
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})
}

func getBodyLimitMiddleware(bodyLimit string) echo.MiddlewareFunc {
	return middleware.BodyLimit(bodyLimit)
}
