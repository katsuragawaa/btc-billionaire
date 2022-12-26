package server

import (
	"net/http"

	"github.com/katsuragawaa/btc-billionaire/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// MapHandlers Map the Server handlers
func (s *Server) MapHandlers(e *echo.Echo) error {
	docs.SwaggerInfo.Title = "BTC Billionaire"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/api/v1")

	ping := v1.Group("/ping")
	ping.GET("", p())

	return nil
}

// @Tags Health
// @Summary Ping app
// @Description Ping server for health check
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /ping [get]
func p() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	}
}
