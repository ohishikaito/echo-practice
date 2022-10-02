//go:build wireinject
// +build wireinject

package wire_example

import (
	"github.com/google/wire"
)

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
