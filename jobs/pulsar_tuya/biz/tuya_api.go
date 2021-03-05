package biz

import (
	"DTU/configs/conf_v1"
	"DTU/pkg"
	"github.com/tuya/tuya-pulsar-sdk-go/pkg/tyutils"
	"os"
)

type EcMail interface {
	Mail([]byte) string
	ToString() string
}

func UidSaveFile(Uid string, config *conf_v1.Config_ConnTuya) {
	if !tyutils.Exists(config.TuyaUidCachePath) {
		file, err := os.Create(config.TuyaUidCachePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		_, err = file.Write([]byte(Uid))
		if err != nil {
			panic(err)
		}
	}
}

func GetEcEmail(config *conf_v1.Config_ConnTuya, mail EcMail) string {
	_, body, err := pkg.HttpPOST(config.MerchantEmailInterface, nil, 3, nil, "")
	if err != nil {
		panic(err)
	}
	if mail.Mail(body) == "" {
		return mail.ToString()
	}
	//UidSaveFile(mail.ToString(),config)
	return ""
}
