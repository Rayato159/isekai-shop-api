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
	_adminRepository "github.com/Rayato159/isekai-shop-api/modules/admin/repository"
	_oauth2Controller "github.com/Rayato159/isekai-shop-api/modules/oauth2/controller"
	_oauth2Service "github.com/Rayato159/isekai-shop-api/modules/oauth2/service"
	_playerRepository "github.com/Rayato159/isekai-shop-api/modules/player/repository"
	"github.com/Rayato159/isekai-shop-api/packages/state"
	"github.com/Rayato159/isekai-shop-api/server/customMiddleware"
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

func NewEchoServer(conf *config.AppConfig, db *gorm.DB) *echoServer {
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
	loggerMiddleware := getLoggerMiddleware(s.app.Logger)
	timeOutMiddleware := getTimeOutMiddleware(s.conf.ServerConfig.Timeout)
	corsMiddleware := getCorsMiddleware(s.conf.ServerConfig.AllowOrigins)
	bodyLimitMiddleware := getBodyLimitMiddleware(s.conf.ServerConfig.BodyLimit)

	s.app.Use(loggerMiddleware)
	s.app.Use(timeOutMiddleware)
	s.app.Use(corsMiddleware)
	s.app.Use(bodyLimitMiddleware)

	// Initialize all custom middlewares
	customerMiddleware := s.getCustomMiddleware()

	// Initialzie all routers
	s.baseRouter.GET("/health", s.healthCheck)

	s.initOAuth2Router()
	s.initPlayerRouter(customerMiddleware)
	s.initItemRouter()
	s.initAdminRouter(customerMiddleware)
	s.initPaymentRouter(customerMiddleware)
	s.initInventoryRouter(customerMiddleware)

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

func (s *echoServer) getCustomMiddleware() customMiddleware.CustomMiddleware {
	stateConfig := s.conf.StateConfig
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

	controller := _oauth2Controller.NewGoogleOAuth2Controller(
		oauth2Service,
		s.conf.OAuth2Config,
		stateProvider,
		s.app.Logger,
	)

	return customMiddleware.NewCustomMiddlewaresImpl(
		controller,
		s.conf.OAuth2Config,
		s.app.Logger,
	)
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
