package biz

import (
	"DTU/configs/conf_v1"
	"DTU/pkg"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	pulsar "github.com/tuya/tuya-pulsar-sdk-go"
	"github.com/tuya/tuya-pulsar-sdk-go/pkg/tylog"
	"github.com/tuya/tuya-pulsar-sdk-go/pkg/tyutils"
	"time"
)

var StopConsumer pulsar.Consumer

type TuyaPulsar interface {
	Parse([]byte) bool
	UidToString() string
	ModeToString() string
}
type TuyaPulsarData struct {
	repo TuyaPulsar
}

func NewTuyaPulsarData(repo TuyaPulsar) *TuyaPulsarData {
	return &TuyaPulsarData{repo: repo}
}

func (p *TuyaPulsarData) Repo() TuyaPulsar {
	return p.repo
}
func StartTuyaPulsar(config *conf_v1.Config_ConnTuya, pulsar2 TuyaPulsar, mail EcMail) {
	pulsar.SetInternalLogLevel(logrus.PanicLevel)
	tylog.SetGlobalLog("sdk", true)
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
	csm.ReceiveAndHandle(context.Background(), &helloHandler{config.TuyaIotCloudAccessCRET[8:24], pulsar2, config, mail})

	time.Sleep(10 * time.Second)
}

type helloHandler struct {
	AesSecret string
	Recv      TuyaPulsar
	Conf      *conf_v1.Config_ConnTuya
	MerMail   EcMail
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

	if !h.Recv.Parse(decode) {
		fmt.Println("解析失败")
		return nil
	}
	fmt.Println("解析后UID:")
	if h.Recv.ModeToString() == "bindUser" {
		fmt.Println("进入")
		resp, err := GetUserInfo(h.Recv.UidToString())
		fmt.Println("获取完成")
		if err != nil {
			panic(err)
		}
		if resp.Result.Email == GetEcEmail(h.Conf, h.MerMail) {
			fmt.Println("进入流程")
			// set GlobalId
			pkg.GlobleId = resp.Result.Uid
			UidSaveFile(resp.Result.Uid, h.Conf)
			StopConsumer.Stop()
		}
	}
	return nil
}
