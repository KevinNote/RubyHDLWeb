package controller

import (
	"github.com/KevinZonda/RubyDHLWeb/controller/ping"
	"github.com/KevinZonda/RubyDHLWeb/controller/ruby"
	"github.com/KevinZonda/RubyDHLWeb/controller/types"
	"github.com/gin-gonic/gin"
)

func Init(r gin.IRouter) {
	register(r, &ping.Controller{}, &ruby.Controller{})
}

func register(r gin.IRouter, cs ...types.IController) {
	for _, c := range cs {
		c.Init(r)
	}
}
