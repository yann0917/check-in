package config

import "time"

type Server struct {
	JueJin        JueJin        `mapstructure:"jue-jin" json:"jue_jin" yaml:"jue-jin"`
	TieBa         TieBa         `mapstructure:"tie-ba" json:"tie_ba" yaml:"tie-ba"`
	Bilibili      Bilibili      `mapstructure:"bilibili" json:"bilibili" yaml:"bilibili"`
	Notification  Notification  `mapstructure:"notification" json:"notification" yaml:"notification"`
	CountdownDays CountdownDays `mapstructure:"countdown-days" json:"countdown_days" yaml:"countdown-days"`
}

type JueJin struct {
	Signature string `mapstructure:"signature" json:"signature" yaml:"signature"`
	Cookie    string `mapstructure:"cookie" json:"cookie" yaml:"cookie"`
}

type TieBa struct {
	Signature string `mapstructure:"signature" json:"signature" yaml:"signature"`
	Cookie    string `mapstructure:"cookie" json:"cookie" yaml:"cookie"`
}

type Bilibili struct {
	Cookie string `mapstructure:"cookie" json:"cookie" yaml:"cookie"`
}

type Notification struct {
	PushPlus   PushPlus   `mapstructure:"push-plus" json:"push_plus" yaml:"push-plus"`
	ServerChan ServerChan `mapstructure:"server-chan" json:"server_chan" yaml:"server-chan"`
}

// PushPlus Push+ 推送设置
type PushPlus struct {
	Token    string `mapstructure:"token" json:"token" yaml:"token"`
	Template string `mapstructure:"template" json:"template" yaml:"template"`
	Channel  string `mapstructure:"channel" json:"channel" yaml:"channel"`
	Webhook  string `mapstructure:"webhook" json:"webhook" yaml:"webhook"`
	Callback string `mapstructure:"callback" json:"callback" yaml:"callback"`
	Topic    string `mapstructure:"topic" json:"topic" yaml:"topic"` // 群组编码，不填仅发送给自己；channel为webhook时无效
}

// ServerChan Server酱推送设置
// FIXME: 方糖服务号推送可能被腾讯弃用
type ServerChan struct {
	SendKey string `mapstructure:"send-key" json:"send_key" yaml:"send-key"`
}

type CountdownDays struct {
	List []Day `mapstructure:"list" json:"list" yaml:"list"`
}

type Day struct {
	Name   string    `mapstructure:"name" json:"name" yaml:"name"`
	Date   time.Time `mapstructure:"date" json:"date" yaml:"date"`
	Remark string    `mapstructure:"remark" json:"remark" yaml:"remark"`
}
