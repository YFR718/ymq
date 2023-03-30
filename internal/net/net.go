package net

import (
	"fmt"
	"github.com/YFR718/ymq/internal/system"
	"github.com/YFR718/ymq/pkg/common"
	"net"
)

func Listen(sys system.System) {
	listen, err := net.Listen("tcp", "127.0.0.1:8848")
	if err != nil {
		common.PrintError(err)
		return
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()

		if err != nil {
			common.PrintError(err)
			return
		}
		go func(conn net.Conn) {
			fmt.Println("连接成功", conn.LocalAddr(), conn.RemoteAddr())
			for {
				msg := make([]byte, 128)
				n, err := conn.Read(msg)
				if err != nil {
					common.PrintError(err)
					break
				}
				fmt.Println("get message:", string(msg[:n]))

				if string(msg[:n]) == "Close" {
					sys.Error <- fmt.Errorf("close")
					break
				}

				conn.Write([]byte("get it!"))
			}

		}(conn)

	}
}
