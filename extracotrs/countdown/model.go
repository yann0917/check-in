package countdown

type Day struct {
	Name   string `mapstructure:"name" json:"name" yaml:"name"`
	Date   string `mapstructure:"date" json:"date" yaml:"date"`
	Remark string `mapstructure:"remark" json:"remark" yaml:"remark"`
}
