package conf_v1

import (
	"flag"
	"github.com/golang/protobuf/jsonpb"
	"io"
	"os"
)

var MiExtend bool
var ListPort string
var TuyaAccessId string
var TuyaCret string
var ConfFile string

func init() {
	flag.BoolVar(&MiExtend, "m", false, "open / close mi active")
	flag.StringVar(&ListPort, "p", ":9901", "set listen port")
	flag.StringVar(&TuyaAccessId, "a", "5jh3wwnsa9wq5xd", "Tuya Iot Cloud Access ID")
	flag.StringVar(&TuyaCret, "c", "bda301def1f95c75", "Tuya Iot Cloud Access CRET")
	flag.StringVar(&ConfFile, "f", "/data/ha/conf/config.json", "DTU Config File path ")
}

// init Config object
// optional select  parameter pass options func set
func NewConfig(options ...func(*Config)) *Config {
	// configurations default value
	Conf := Config{
		LogPath:                "/var/log/v2.log",
		DeviceListIntervalTime: 10,
		ExtendMi:               false,
		ListenPort:             ":9901",
		ConnConfig: &Config_Communication{
			EcConfig: &Config_ConnEc{
				EventUpdateInterface:  "http://127.0.0.1:7001/gateway/violation",
				HealthUpdateInterface: "http://127.0.0.1:7001/gateway/status",
				DeviceListInterface:   "/ha/device/list",
			},
			AlleyesConfig: &Config_ConnEyes{
				EventInfoRecvInterface:    "/v1/wntime/agent/violations/receive",
				GettingIPCIntervalTime:    10,
				IPCStatusGettingInterface: "http://127.0.0.1:9008/v1/alleyes/syncbot/BoxStatus/get",
			},
			TuyaConfig: &Config_ConnTuya{
				MerchantEmailInterface: "http://127.0.0.1:7001/gateway/getRestaurant",
				TuyaIotCloudAccessId:   "5jh3ww4hls",
				TuyaIotCloudAccessCRET: "bdf1f95c75",
				TuyaUidCachePath:       "/data/ha/conf/UID",
			},
			OtherConfig: &Config_ConnOther{
				ByPassInterface: "http://127.0.0.1:9001/recevice/violations",
			},
		},
	}
	// run options , set field value
	for _, option := range options {
		option(&Conf)
	}

	return &Conf
}

//Load conf file
//return a byte type slice
func LoadConfig(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

//write config file
//write config file from c
func ConfigWriteFile(c *Config) error {
	m := jsonpb.Marshaler{Indent: " ", EmitDefaults: true}
	file, err := os.OpenFile(ConfFile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	return m.Marshal(file, c)
}

// applyConfig
//apply c from reader
func ApplyConfig(c *Config, reader io.Reader) {
	err := jsonpb.Unmarshal(reader, c)
	if err != nil {
		panic(err)
	}
	//TuyaSDK.SetEnv("https://openapi.tuyacn.com", c.ConnConfig.TuyaConfig.TuyaIotCloudAccessId, c.ConnConfig.TuyaConfig.TuyaIotCloudAccessCRET)
}
