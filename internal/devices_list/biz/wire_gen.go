// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package biz

import (
	"DTU/internal/devices_list/data"
	"github.com/google/wire"
)

// Injectors from getter.go:

func initDeviceDataSet() *CacheBiz {
	deviceListCache := data.NewDeviceListCache()
	cacheBiz := NewCacheBiz(deviceListCache)
	return cacheBiz
}

// getter.go:

var DeviceDataSet = wire.NewSet(data.NewDeviceListCache, wire.Bind(new(CacheList), new(*data.DeviceListCache)), NewCacheBiz)
