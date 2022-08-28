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

// CastServerCtx は echo.Context を *ServerCtx に Cast します
func CastServerCtx(c echo.Context) *ServerCtx {
	return c.(*ServerCtx)
}
