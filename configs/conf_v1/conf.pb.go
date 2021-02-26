// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: conf.proto

package conf_v1

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogPath                string                `protobuf:"bytes,1,opt,name=LogPath,proto3" json:"LogPath,omitempty"`                                //  日志文件位置
	ExtendMi               bool                  `protobuf:"varint,2,opt,name=ExtendMi,proto3" json:"ExtendMi,omitempty"`                             //    启用废弃Mi接口选项
	ConnConfig             *Config_Communication `protobuf:"bytes,3,opt,name=ConnConfig,proto3" json:"ConnConfig,omitempty"`                          // 通信相关配置
	DeviceListIntervalTime int32                 `protobuf:"varint,4,opt,name=DeviceListIntervalTime,proto3" json:"DeviceListIntervalTime,omitempty"` //设备列表生成时间
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetLogPath() string {
	if x != nil {
		return x.LogPath
	}
	return ""
}

func (x *Config) GetExtendMi() bool {
	if x != nil {
		return x.ExtendMi
	}
	return false
}

func (x *Config) GetConnConfig() *Config_Communication {
	if x != nil {
		return x.ConnConfig
	}
	return nil
}

func (x *Config) GetDeviceListIntervalTime() int32 {
	if x != nil {
		return x.DeviceListIntervalTime
	}
	return 0
}

//
//所有通信类接口定义
//all Communication interfaces defintions
type Config_Communication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlleyesConfig *Config_ConnEyes  `protobuf:"bytes,1,opt,name=AlleyesConfig,proto3" json:"AlleyesConfig,omitempty"` //  白目链接信息
	EcConfig      *Config_ConnEc    `protobuf:"bytes,2,opt,name=EcConfig,proto3" json:"EcConfig,omitempty"`           // 边缘计算链接信息
	TuyaConfig    *Config_ConnTuya  `protobuf:"bytes,3,opt,name=TuyaConfig,proto3" json:"TuyaConfig,omitempty"`       //     涂鸦链接信息
	OtherConfig   *Config_ConnOther `protobuf:"bytes,4,opt,name=OtherConfig,proto3" json:"OtherConfig,omitempty"`     //   其他链接信息
	MiConfig      *Config_ConnMi    `protobuf:"bytes,5,opt,name=MiConfig,proto3" json:"MiConfig,omitempty"`           //  小米链接信息
}

func (x *Config_Communication) Reset() {
	*x = Config_Communication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_Communication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_Communication) ProtoMessage() {}

func (x *Config_Communication) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_Communication.ProtoReflect.Descriptor instead.
func (*Config_Communication) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Config_Communication) GetAlleyesConfig() *Config_ConnEyes {
	if x != nil {
		return x.AlleyesConfig
	}
	return nil
}

func (x *Config_Communication) GetEcConfig() *Config_ConnEc {
	if x != nil {
		return x.EcConfig
	}
	return nil
}

func (x *Config_Communication) GetTuyaConfig() *Config_ConnTuya {
	if x != nil {
		return x.TuyaConfig
	}
	return nil
}

func (x *Config_Communication) GetOtherConfig() *Config_ConnOther {
	if x != nil {
		return x.OtherConfig
	}
	return nil
}

func (x *Config_Communication) GetMiConfig() *Config_ConnMi {
	if x != nil {
		return x.MiConfig
	}
	return nil
}

type Config_ConnEyes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IPCStatusGettingInterface string `protobuf:"bytes,1,opt,name=IPCStatusGettingInterface,proto3" json:"IPCStatusGettingInterface,omitempty"` //    获取摄像头状态接口
	EventInfoRecvInterface    string `protobuf:"bytes,2,opt,name=EventInfoRecvInterface,proto3" json:"EventInfoRecvInterface,omitempty"`       //    接受事件信息接口
	GettingIPCIntervalTime    int32  `protobuf:"varint,3,opt,name=GettingIPCIntervalTime,proto3" json:"GettingIPCIntervalTime,omitempty"`      //   获取摄像头状态间隔时间 单位秒
}

func (x *Config_ConnEyes) Reset() {
	*x = Config_ConnEyes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_ConnEyes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_ConnEyes) ProtoMessage() {}

func (x *Config_ConnEyes) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_ConnEyes.ProtoReflect.Descriptor instead.
func (*Config_ConnEyes) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Config_ConnEyes) GetIPCStatusGettingInterface() string {
	if x != nil {
		return x.IPCStatusGettingInterface
	}
	return ""
}

func (x *Config_ConnEyes) GetEventInfoRecvInterface() string {
	if x != nil {
		return x.EventInfoRecvInterface
	}
	return ""
}

func (x *Config_ConnEyes) GetGettingIPCIntervalTime() int32 {
	if x != nil {
		return x.GettingIPCIntervalTime
	}
	return 0
}

type Config_ConnEc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventUpdateInterface  string `protobuf:"bytes,1,opt,name=EventUpdateInterface,proto3" json:"EventUpdateInterface,omitempty"`   //    事件上传接口
	HealthUpdateInterface string `protobuf:"bytes,2,opt,name=HealthUpdateInterface,proto3" json:"HealthUpdateInterface,omitempty"` //    心跳上传接口
}

func (x *Config_ConnEc) Reset() {
	*x = Config_ConnEc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_ConnEc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_ConnEc) ProtoMessage() {}

func (x *Config_ConnEc) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_ConnEc.ProtoReflect.Descriptor instead.
func (*Config_ConnEc) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Config_ConnEc) GetEventUpdateInterface() string {
	if x != nil {
		return x.EventUpdateInterface
	}
	return ""
}

func (x *Config_ConnEc) GetHealthUpdateInterface() string {
	if x != nil {
		return x.HealthUpdateInterface
	}
	return ""
}

type Config_ConnOther struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ByPassInterface string `protobuf:"bytes,1,opt,name=ByPassInterface,proto3" json:"ByPassInterface,omitempty"` //  特殊设备上传接口
}

func (x *Config_ConnOther) Reset() {
	*x = Config_ConnOther{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_ConnOther) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_ConnOther) ProtoMessage() {}

func (x *Config_ConnOther) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_ConnOther.ProtoReflect.Descriptor instead.
func (*Config_ConnOther) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Config_ConnOther) GetByPassInterface() string {
	if x != nil {
		return x.ByPassInterface
	}
	return ""
}

type Config_ConnTuya struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerchantEmailInterface string `protobuf:"bytes,1,opt,name=MerchantEmailInterface,proto3" json:"MerchantEmailInterface,omitempty"` //   获取商家Email接口
	TuyaIotCloudAccessId   string `protobuf:"bytes,2,opt,name=TuyaIotCloudAccessId,proto3" json:"TuyaIotCloudAccessId,omitempty"`     //    tuya 云平台AccessID
	TuyaIotCloudAccessCRET string `protobuf:"bytes,3,opt,name=TuyaIotCloudAccessCRET,proto3" json:"TuyaIotCloudAccessCRET,omitempty"` //   tuya 云平台cret
	TuyaUidCachePath       string `protobuf:"bytes,4,opt,name=TuyaUidCachePath,proto3" json:"TuyaUidCachePath,omitempty"`             // tuya 云平台用户id
}

func (x *Config_ConnTuya) Reset() {
	*x = Config_ConnTuya{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_ConnTuya) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_ConnTuya) ProtoMessage() {}

func (x *Config_ConnTuya) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_ConnTuya.ProtoReflect.Descriptor instead.
func (*Config_ConnTuya) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Config_ConnTuya) GetMerchantEmailInterface() string {
	if x != nil {
		return x.MerchantEmailInterface
	}
	return ""
}

func (x *Config_ConnTuya) GetTuyaIotCloudAccessId() string {
	if x != nil {
		return x.TuyaIotCloudAccessId
	}
	return ""
}

func (x *Config_ConnTuya) GetTuyaIotCloudAccessCRET() string {
	if x != nil {
		return x.TuyaIotCloudAccessCRET
	}
	return ""
}

func (x *Config_ConnTuya) GetTuyaUidCachePath() string {
	if x != nil {
		return x.TuyaUidCachePath
	}
	return ""
}

//
//已经弃用 ！！！！！！
//小米通信接口定义
//如非必然请勿修改此配置
//The definition this is  abandon  !!!!!
//Do not modify this definition    !!!!
type Config_ConnMi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayReturnInterface      string `protobuf:"bytes,1,opt,name=GatewayReturnInterface,proto3" json:"GatewayReturnInterface,omitempty"`           //   网关回传信息接口地址  如192.168.137.1:4321
	GatewayGettingListInterface string `protobuf:"bytes,2,opt,name=GatewayGettingListInterface,proto3" json:"GatewayGettingListInterface,omitempty"` //  获取设备列表接口地址 如192.168.137.1:4567
	GatewayDiscoverTimeOut      int32  `protobuf:"varint,3,opt,name=GatewayDiscoverTimeOut,proto3" json:"GatewayDiscoverTimeOut,omitempty"`          //    自动发现网关的超时时间  单位秒
	GatewayTimeOutCount         string `protobuf:"bytes,4,opt,name=GatewayTimeOutCount,proto3" json:"GatewayTimeOutCount,omitempty"`                 // 自动发现网关的超时重试次数
	GatewayCacheSize            int32  `protobuf:"varint,5,opt,name=GatewayCacheSize,proto3" json:"GatewayCacheSize,omitempty"`                      //  网关通信接口缓存大小 单位字节
	GatewayGroupInterface       string `protobuf:"bytes,6,opt,name=GatewayGroupInterface,proto3" json:"GatewayGroupInterface,omitempty"`             //组播监听网卡名称
}

func (x *Config_ConnMi) Reset() {
	*x = Config_ConnMi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config_ConnMi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config_ConnMi) ProtoMessage() {}

func (x *Config_ConnMi) ProtoReflect() protoreflect.Message {
	mi := &file_conf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config_ConnMi.ProtoReflect.Descriptor instead.
func (*Config_ConnMi) Descriptor() ([]byte, []int) {
	return file_conf_proto_rawDescGZIP(), []int{0, 5}
}

func (x *Config_ConnMi) GetGatewayReturnInterface() string {
	if x != nil {
		return x.GatewayReturnInterface
	}
	return ""
}

func (x *Config_ConnMi) GetGatewayGettingListInterface() string {
	if x != nil {
		return x.GatewayGettingListInterface
	}
	return ""
}

func (x *Config_ConnMi) GetGatewayDiscoverTimeOut() int32 {
	if x != nil {
		return x.GatewayDiscoverTimeOut
	}
	return 0
}

func (x *Config_ConnMi) GetGatewayTimeOutCount() string {
	if x != nil {
		return x.GatewayTimeOutCount
	}
	return ""
}

func (x *Config_ConnMi) GetGatewayCacheSize() int32 {
	if x != nil {
		return x.GatewayCacheSize
	}
	return 0
}

func (x *Config_ConnMi) GetGatewayGroupInterface() string {
	if x != nil {
		return x.GatewayGroupInterface
	}
	return ""
}

var File_conf_proto protoreflect.FileDescriptor

var file_conf_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x0a, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4c, 0x6f, 0x67, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x4d, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x4d, 0x69, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x36, 0x0a, 0x16, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0xae, 0x02,
	0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3e, 0x0a, 0x0d, 0x41, 0x6c, 0x6c, 0x65, 0x79, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x45, 0x79, 0x65, 0x73,
	0x52, 0x0d, 0x41, 0x6c, 0x6c, 0x65, 0x79, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x32, 0x0a, 0x08, 0x45, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x45, 0x63, 0x52, 0x08, 0x45, 0x63, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x38, 0x0a, 0x0a, 0x54, 0x75, 0x79, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x54, 0x75, 0x79,
	0x61, 0x52, 0x0a, 0x54, 0x75, 0x79, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a,
	0x0b, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x0b, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x0a, 0x08, 0x4d, 0x69,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x4d, 0x69, 0x52, 0x08, 0x4d, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0xb8,
	0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x6e, 0x45, 0x79, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x19, 0x49,
	0x50, 0x43, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x47, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19,
	0x49, 0x50, 0x43, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x47, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x63, 0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x63, 0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x36, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x50, 0x43, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x16, 0x47, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x50, 0x43, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x72, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x6e, 0x45, 0x63, 0x12, 0x32, 0x0a, 0x14, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x1a, 0x35, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x6e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x42, 0x79,
	0x50, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x42, 0x79, 0x50, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x1a, 0xda, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x6e, 0x54, 0x75, 0x79,
	0x61, 0x12, 0x36, 0x0a, 0x16, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x16, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x54, 0x75, 0x79,
	0x61, 0x49, 0x6f, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x54, 0x75, 0x79, 0x61, 0x49, 0x6f, 0x74,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x36, 0x0a,
	0x16, 0x54, 0x75, 0x79, 0x61, 0x49, 0x6f, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x43, 0x52, 0x45, 0x54, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x54,
	0x75, 0x79, 0x61, 0x49, 0x6f, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x43, 0x52, 0x45, 0x54, 0x12, 0x2a, 0x0a, 0x10, 0x54, 0x75, 0x79, 0x61, 0x55, 0x69, 0x64,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x54, 0x75, 0x79, 0x61, 0x55, 0x69, 0x64, 0x43, 0x61, 0x63, 0x68, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x1a, 0xce, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x69, 0x12, 0x36, 0x0a, 0x16,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x1b, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x47,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1b, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x47, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x30,
	0x0a, 0x13, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2a, 0x0a, 0x10, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x34, 0x0a, 0x15,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_proto_rawDescOnce sync.Once
	file_conf_proto_rawDescData = file_conf_proto_rawDesc
)

func file_conf_proto_rawDescGZIP() []byte {
	file_conf_proto_rawDescOnce.Do(func() {
		file_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_proto_rawDescData)
	})
	return file_conf_proto_rawDescData
}

var file_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_conf_proto_goTypes = []interface{}{
	(*Config)(nil),               // 0: conf.v1.Config
	(*Config_Communication)(nil), // 1: conf.v1.Config.Communication
	(*Config_ConnEyes)(nil),      // 2: conf.v1.Config.ConnEyes
	(*Config_ConnEc)(nil),        // 3: conf.v1.Config.ConnEc
	(*Config_ConnOther)(nil),     // 4: conf.v1.Config.ConnOther
	(*Config_ConnTuya)(nil),      // 5: conf.v1.Config.ConnTuya
	(*Config_ConnMi)(nil),        // 6: conf.v1.Config.ConnMi
}
var file_conf_proto_depIdxs = []int32{
	1, // 0: conf.v1.Config.ConnConfig:type_name -> conf.v1.Config.Communication
	2, // 1: conf.v1.Config.Communication.AlleyesConfig:type_name -> conf.v1.Config.ConnEyes
	3, // 2: conf.v1.Config.Communication.EcConfig:type_name -> conf.v1.Config.ConnEc
	5, // 3: conf.v1.Config.Communication.TuyaConfig:type_name -> conf.v1.Config.ConnTuya
	4, // 4: conf.v1.Config.Communication.OtherConfig:type_name -> conf.v1.Config.ConnOther
	6, // 5: conf.v1.Config.Communication.MiConfig:type_name -> conf.v1.Config.ConnMi
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_conf_proto_init() }
func file_conf_proto_init() {
	if File_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_Communication); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_ConnEyes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_ConnEc); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_ConnOther); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_ConnTuya); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_conf_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config_ConnMi); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_proto_goTypes,
		DependencyIndexes: file_conf_proto_depIdxs,
		MessageInfos:      file_conf_proto_msgTypes,
	}.Build()
	File_conf_proto = out.File
	file_conf_proto_rawDesc = nil
	file_conf_proto_goTypes = nil
	file_conf_proto_depIdxs = nil
}
