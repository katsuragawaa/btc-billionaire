package server

import (
	"net/http"

	"github.com/katsuragawaa/btc-billionaire/docs"
	transactionsHttp "github.com/katsuragawaa/btc-billionaire/internal/transactions/delivery/http"
	"github.com/katsuragawaa/btc-billionaire/pkg/logger"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// MapHandlers Map the Server handlers
func (s *Server) MapHandlers(e *echo.Echo) error {
	docs.SwaggerInfo.Title = "BTC Billionaire"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/api/v1")

	ping := v1.Group("/ping")
	ping.GET("", pingHandler(s.logger))

	transactionsGroup := v1.Group("/transactions")
	transactionsHandlers := transactionsHttp.NewTransactionsHandlers(s.cfg, s.logger)
	transactionsHttp.MapTransactionsRoutes(transactionsGroup, transactionsHandlers)

	return nil
}

// @Tags Health
// @Summary Ping app
// @Description Ping server for health check
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /ping [get]
func pingHandler(logger logger.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.Info("Health check ping")
		return c.JSON(http.StatusOK, "pong")
	}
}
