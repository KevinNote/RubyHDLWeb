package shared

import "github.com/KevinZonda/GoX/pkg/iox"

var Prelude []byte
var PreludePath string

func initPrelude() {
	cfg := GetConfig()
	if cfg.Prelude != "" {
		var err error
		PreludePath = cfg.Prelude
		Prelude, err = iox.ReadAllByte(cfg.Prelude)
		if err != nil {
			panic(err)
		}
	}
}
