package biz

import "DTU/pkg"

type CacheList interface {
	Exists(DeviceCode string) bool
	Append(data pkg.DeviceList)
}

type CacheBiz struct {
	Cache CacheList
}

func NewCacheBiz(list CacheList) *CacheBiz {
	return &CacheBiz{list}
}

func (p *CacheBiz) TuyaListCallBack(data interface{}) {
	DeviceData, ok := data.(*pkg.DeviceList)
	if !ok {
		panic(ok)
	}
	if DevicelistResp.Cache.Exists(DeviceData.DeviceCode) {
		DevicelistResp.Cache.Append(*DeviceData)
	}
}
