package shared

import (
	"time"

	"github.com/KevinZonda/RubyDHLWeb/lib/RubyHDL"
)

var Ruby *RubyHDL.RubyHDL

func initRuby() {
	cfg := GetConfig()
	Ruby = RubyHDL.NewRubyHDL(cfg.RcPath, cfg.RePath, time.Second*time.Duration(cfg.Timeout))
}
