package main

import (
	"fmt"

	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/GoX/pkg/panicx"
	"github.com/KevinZonda/RubyDHLWeb/controller"
	"github.com/KevinZonda/RubyDHLWeb/shared"
)

func initCfg() {
	bs, err := iox.ReadAllByte("config.json")
	panicx.NotNilErr(err)
	panicx.NotNilErr(shared.LoadConfig(bs))
}

func main() {
	fmt.Println("Loading Config...")
	initCfg()

	fmt.Println("Initialising Shared...")
	shared.Init()

	fmt.Println("Initialising Controller...")
	controller.Init(shared.Engine)

	shared.RunGin()
}
