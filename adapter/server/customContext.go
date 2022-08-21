package server

import (
	"github.com/labstack/echo"
	"github.com/ohishikaito/echo-practice/dject"
)

type (
	CustomContext struct {
		echo.Context
		dject.Container
	}
)
