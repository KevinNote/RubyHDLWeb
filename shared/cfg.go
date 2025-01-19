package shared

import "encoding/json"

type Config struct {
	Addr  string `json:"addr"`
	Debug bool   `json:"debug"`

	TaskDir string `json:"task_dir"`

	Prelude string `json:"prelude"`

	RcPath string `json:"rc_path"`
	RePath string `json:"re_path"`
}

var cfg *Config

func GetConfig() *Config {
	return cfg
}

func LoadConfig(bs []byte) error {
	return json.Unmarshal(bs, &cfg)
}
