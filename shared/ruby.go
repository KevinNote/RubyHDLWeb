package shared

import (
	"github.com/KevinZonda/RubyDHLWeb/lib/RubyDHL"
	"time"
)

var Ruby *RubyDHL.RubyDHL

func initRuby() {
	cfg := GetConfig()
	Ruby = RubyDHL.NewRubyDHL(cfg.RcPath, cfg.RePath, time.Second*time.Duration(cfg.Timeout))
}
