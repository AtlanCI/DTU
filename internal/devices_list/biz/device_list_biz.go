package biz

import (
	"DTU/pkg"
	"github.com/pkg/errors"
)

type CacheList interface {
	Exists(DeviceCode string) bool
	Append(data pkg.DeviceList)
}

type CacheBiz struct {
	Code  int       `json:"code"`
	Cache CacheList `json:"data"`
	Msg   string    `json:"msg"`
}

func NewCacheBiz(list CacheList) *CacheBiz {
	return &CacheBiz{Cache: list}
}

func (p *CacheBiz) TuyaListCallBack(data interface{}) {
	DeviceData, ok := data.(*pkg.DeviceList)
	if !ok {
		pkg.LogSugar.Sugar().Warn(errors.New("read channel for Tuya Device failed. unknown object"))
	}
	if DevicelistResp.Cache.Exists(DeviceData.DeviceCode) {
		DevicelistResp.Cache.Append(*DeviceData)
	}
}
