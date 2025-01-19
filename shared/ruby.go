package shared

import "github.com/KevinZonda/RubyDHLWeb/lib/RubyDHL"

var Ruby *RubyDHL.RubyDHL

func initRuby() {
	cfg := GetConfig()
	Ruby = RubyDHL.NewRubyDHL(cfg.RcPath, cfg.RePath)
}
