package paypal

import (
	"github.com/ohishikaito/echo-practice/adapter/environment"
)

type (
	paypalClient struct {
		env environment.Env
	}
	PaypalClient interface{}
)

func NewPaypalClient(env environment.Env) PaypalClient {
	return &paypalClient{
		env,
	}
}
