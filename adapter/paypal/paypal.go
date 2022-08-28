package paypal

import (
	"github.com/ohishikaito/echo-practice/adapter/env"
)

type (
	paypalClient struct {
		env env.Env
	}
	PaypalClient interface{}
)

func NewPaypalClient(env env.Env) PaypalClient {
	return &paypalClient{
		env,
	}
}
