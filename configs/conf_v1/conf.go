package conf_v1

type DialOption func(*Config)

func ConfigEcEventInterface(p string) DialOption {
	return func(s *Config) {
		s.ConnConfig.EcConfig.EventUpdateInterface = p
	}
}

func ConfigEcHealthInterface(p string) DialOption {
	return func(s *Config) {
		s.ConnConfig.EcConfig.HealthUpdateInterface = p
	}
}

func ConfigEyesIpcIntervalTime(p int32) DialOption {
	return func(s *Config) {
		s.ConnConfig.AlleyesConfig.GetGettingIPCIntervalTime = p
	}
}

func ConfigEyesIpcIntreface(p string) DialOption {
	return func(s *Config) {
		s.ConnConfig.AlleyesConfig.IPCStatusGettingInterface = p

	}
}

func ConfigEyesEventInterface(p string) DialOption {
	return func(s *Config) {
		s.ConnConfig.AlleyesConfig.EventInfoRecvInterface = p
	}
}

func ConfigTuyaEmailInterface(p string) DialOption {
	return func(s *Config) {
		s.ConnConfig.TuyaConfig.MerchantEmailInterface = p
	}
}

func ConfigTuyaAccessId(p string) DialOption {
	return func(s *Config) {
		s.ConnConfig.TuyaConfig.TuyaIotCloudAccessId = p

	}
}

func ConfigTuyaAccessCRET(p string) DialOption {
	return func(s *Config) {
		s.ConnConfig.TuyaConfig.TuyaIotCloudAccessCRET = p

	}
}

func COnfigOtherByPassInterface(p string) DialOption {
	return func(s *Config) {
		s.ConnConfig.OtherConfig.ByPassInterface = p

	}
}
func ConfigExtendMi(p bool) DialOption {
	return func(s *Config) {
		s.ExtendMi = p
	}
}

func ConfigDeviceListIntervalTime(p int32) DialOption {
	return func(s *Config) {
		s.DeviceListIntervalTime = p

	}
}
