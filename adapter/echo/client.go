package echo

import "github.com/labstack/echo/v4"

type (
	echoClient struct {
		E *echo.Echo
	}
)

func NewEchoClient() *echoClient{
	e := echo.New()
	return &echoClient{e}
}