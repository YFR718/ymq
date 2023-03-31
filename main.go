package main

import (
	"github.com/YFR718/ymq/internal/net"
	"github.com/YFR718/ymq/internal/system"
	"github.com/YFR718/ymq/internal/topic"
	"github.com/YFR718/ymq/pkg/common"
)

func main() {
	sys := system.System{Error: make(chan error)}
	topic.InitManager()
	go net.Listen(sys)

	select {
	case err := <-sys.Error:
		common.PrintError(err)
	}
}
