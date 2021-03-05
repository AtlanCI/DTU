package biz

import (
	"DTU/configs/conf_v1"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	pulsar "github.com/tuya/tuya-pulsar-sdk-go"
	"github.com/tuya/tuya-pulsar-sdk-go/pkg/tyutils"
	"time"
)

var StopConsumer pulsar.Consumer

type tuyaPulsar interface {
	Parse([]byte) bool
	UidToString() string
	ModeToString() string
}

func StartTuyaPulsar(config *conf_v1.Config_ConnTuya, pulsar2 tuyaPulsar) {
	topic := pulsar.TopicForAccessID(config.TuyaIotCloudAccessId)

	//create client
	cfg := pulsar.ClientConfig{PulsarAddr: pulsar.PulsarAddrCN}
	c := pulsar.NewClient(cfg)

	//create consume
	csmCfg := pulsar.ConsumerConfig{
		Topic: topic,
		Auth:  pulsar.NewAuthProvider(config.TuyaIotCloudAccessId, config.TuyaIotCloudAccessCRET),
	}
	csm, _ := c.NewConsumer(csmCfg)
	StopConsumer = csm
	//handle message
	csm.ReceiveAndHandle(context.Background(), &helloHandler{config.TuyaIotCloudAccessCRET[8:24], pulsar2, config})

	time.Sleep(10 * time.Second)
}

type helloHandler struct {
	AesSecret string
	Recv      tuyaPulsar
	Conf      *conf_v1.Config_ConnTuya
}

func (h *helloHandler) HandlePayload(ctx context.Context, msg *pulsar.Message, payload []byte) error {

	// let's decode the payload with AES
	m := map[string]interface{}{}
	err := json.Unmarshal(payload, &m)
	if err != nil {
		return nil
	}
	bs := m["data"].(string)
	de, err := base64.StdEncoding.DecodeString(string(bs))
	if err != nil {
		return nil
	}
	// 解密数据
	decode := tyutils.EcbDecrypt(de, []byte(h.AesSecret))
	//fmt.Println(string(decode))
	fmt.Println(string(decode))

	// 数据处理逻辑块

	//var tuyaRe TuyaRecv
	//err = json.Unmarshal(decode, &tuyaRe)
	//if err != nil {
	//	log.Write.Println("*********************解析通道值失败*****************************************************")
	//	return err
	//	//panic(err)
	//}

	//if tuyaRe.BizCode == "bindUser" {
	if !h.Recv.Parse(decode) {
		return nil
	}
	//fmt.Println("解析后UID:", tuyaRe.BizData.UID)
	if h.Recv.ModeToString() == "bindUser" {
		resp, err := GetUserInfo(h.Recv.UidToString())
		if err != nil {
			panic(err)
		}
		if resp.Result.Email == GetEcEmail(h.Conf, *new(EcMail)) {
			UidSaveFile(resp.Result.Uid, h.Conf)
			StopConsumer.Stop()
		}
	}
	//	/
	//	if err != nil {
	//		//fmt.Println("获取信息错误")
	//		log.Write.Println("*********************获取用户信息失败*****************************************************")
	//		return err
	//
	//		//panic(err)
	//	}
	//	//fmt.Println("用户详细",resp)
	//	//log.Write.Println("用户详细",resp)
	//
	//	switch MerchantEmail {
	//	case "":
	//		//if resp.Result.Email == conf.IotEmail {
	//		//	UidSaveFile(resp.Result.Uid)
	//		//	Uid = resp.Result.Uid
	//		//	StopConsumer.Stop()
	//		//	//os.Exit(153)
	//		//}
	//		log.Write.Fatal("用户邮箱为空")
	//		return  nil
	//	default:
	//		if resp.Result.Email == MerchantEmail {
	//			UidSaveFile(resp.Result.Uid)
	//			Uid = resp.Result.Uid
	//			StopConsumer.Stop()
	//			//os.Exit(153)
	//		}
	//	}
	//
	//}

	//tylog.Info("aes decode", tylog.ByteString("decode payload", decode))

	return nil
}
