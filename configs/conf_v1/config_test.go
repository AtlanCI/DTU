package conf_v1

import (
	"fmt"
	"testing"
)

func TestApplyConfig(t *testing.T) {
	file := LoadConfig("/data/ha/conf/ha.conf")
	a := &Config{}
	ApplyConfig(a,file)
	fmt.Print(a)
}



func TestConfigWriteFile(t *testing.T) {
	if err := ConfigWriteFile(NewConfig());err != nil {
		panic(err)
	}
}

func TestNewConfig2(t *testing.T) {
	config := NewConfig(ConfigExtendMi(MiExtend),ConfigTuyaAccessId(TuyaAccessId),ConfigTuyaAccessCRET(TuyaCret),ConfigListenPort(ListPort))
	fmt.Print(config)
}