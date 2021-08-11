package config

type Server struct {
	JueJin       JueJin       `mapstructure:"jue-jin" json:"jue_jin" yaml:"jue-jin"`
	TieBa        TieBa        `mapstructure:"tie-ba" json:"tie_ba" yaml:"tie-ba"`
	Notification Notification `mapstructure:"notification" json:"notification" yaml:"notification"`
}

type JueJin struct {
	Signature string `mapstructure:"signature" json:"signature" yaml:"signature"`
	Cookie    string `mapstructure:"cookie" json:"cookie" yaml:"cookie"`
}

type TieBa struct {
	Signature string `mapstructure:"signature" json:"signature" yaml:"signature"`
	Cookie    string `mapstructure:"cookie" json:"cookie" yaml:"cookie"`
}

type Notification struct {
	PushPlus   PushPlus   `mapstructure:"push-plus" json:"push_plus" yaml:"push-plus"`
	ServerChan ServerChan `mapstructure:"server-chan" json:"server_chan" yaml:"server-chan"`
}

// PushPlus Push+ 推送设置
type PushPlus struct {
	Token string `mapstructure:"token" json:"token" yaml:"token"`
}

// ServerChan Server酱推送设置
// FIXME: 方糖服务号推送可能被腾讯弃用
type ServerChan struct {
	SendKey string `mapstructure:"send-key" json:"send_key" yaml:"send-key"`
}
