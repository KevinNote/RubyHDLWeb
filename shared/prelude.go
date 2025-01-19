package shared

import "github.com/KevinZonda/GoX/pkg/iox"

var Prelude []byte

func initPrelude() {
	cfg := GetConfig()
	if cfg.Prelude != "" {
		var err error
		Prelude, err = iox.ReadAllByte(cfg.Prelude)
		if err != nil {
			panic(err)
		}
	}
}
