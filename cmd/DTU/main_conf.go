package main

import (
	"DTU/configs/conf_v1"
	"DTU/pkg"
)

// not conf file and  not argument
func FirstInitConf() *conf_v1.Config {
	conf := conf_v1.NewConfig()
	if err := conf_v1.ConfigWriteFile(conf); err != nil {
		panic(err)
	}
	return conf
}

// have conf file and have argument
func ExchangeConf() *conf_v1.Config {
	File := conf_v1.LoadConfig(conf_v1.ConfFile)

	// change default value
	conf := conf_v1.NewConfig(
		conf_v1.ConfigExtendMi(conf_v1.MiExtend),
		conf_v1.ConfigListenPort(conf_v1.ListPort),
		conf_v1.ConfigTuyaAccessId(conf_v1.TuyaAccessId),
		conf_v1.ConfigTuyaAccessCRET(conf_v1.TuyaCret),
	)
	conf_v1.ApplyConfig(conf, File)
	return conf
}

//not conf file but have argument
func WriteChangeConf() *conf_v1.Config {
	// change default value
	conf := conf_v1.NewConfig(
		conf_v1.ConfigExtendMi(conf_v1.MiExtend),
		conf_v1.ConfigListenPort(conf_v1.ListPort),
		conf_v1.ConfigTuyaAccessId(conf_v1.TuyaAccessId),
		conf_v1.ConfigTuyaAccessCRET(conf_v1.TuyaCret),
	)
	if err := conf_v1.ConfigWriteFile(conf); err != nil {
		panic(err)
	}
	return conf
}

// have conf file but not argument
func ReadFileConf() *conf_v1.Config {
	File := conf_v1.LoadConfig(conf_v1.ConfFile)

	// change default value
	conf := conf_v1.NewConfig()
	conf_v1.ApplyConfig(conf, File)
	return conf
}

func NewConf() *conf_v1.Config {
	if pkg.Exists(conf_v1.ConfFile) {
		return ExchangeConf()
	}
	return WriteChangeConf()
}
