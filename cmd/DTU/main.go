// +build wireinject
package DTU

import "github.com/google/wire"

func initAppProviderSet() {
	wire.Build(AppProviderSet)
}
