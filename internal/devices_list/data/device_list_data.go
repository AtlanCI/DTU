package data

import "DTU/pkg"

type DeviceListCache []pkg.DeviceList

func NewDeviceListCache() DeviceListCache {
	return []pkg.DeviceList{}
}

//exists is false  if not exists return trun
func (p DeviceListCache) Exists(DeviceCode string) bool {
	for _, v := range p {
		if DeviceCode == v.DeviceCode {
			return false
		}
	}
	return true
}

// append to DeviceListCache
func (p DeviceListCache) Append(data pkg.DeviceList) {
	p = append(p, data)
}
