//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/mundo/pkg/adapterx"
	"github.com/blackhorseya/mundo/pkg/transports/httpx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var providerSet = wire.NewSet(httpx.NewServer)

func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(newService, providerSet))
}

func NewRestful() (adapterx.Restful, error) {
	panic(wire.Build(newRestful, providerSet))
}
