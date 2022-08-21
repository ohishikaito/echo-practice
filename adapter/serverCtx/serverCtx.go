package serverCtx

import (
	"github.com/labstack/echo"
	"github.com/ohishikaito/echo-practice/dject"
)

type (
	ServerCtx struct {
		echo.Context
		dject.Container
	}
)
