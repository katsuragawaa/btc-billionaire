package transactions

import "github.com/labstack/echo/v4"

type Handlers interface {
	Create() echo.HandlerFunc
	GetPerHours() echo.HandlerFunc
	GetBalance() echo.HandlerFunc
}
