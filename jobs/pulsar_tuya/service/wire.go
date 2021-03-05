// +build wireinject

package service

import (
	"github.com/google/wire"
)

func initializeNewGreeterService() *GreeterService {
	wire.Build(ServiceProviderSet)
	return &GreeterService{}
}
