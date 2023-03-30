package net

import (
	"fmt"
	"net"
	"testing"
)

func Test_TCP_Server(t *testing.T) {
	listen, _ := net.Listen("tcp", "127.0.0.1:8848")
	defer listen.Close()

	conn, _ := listen.Accept() // 建立连接
	//conn.SetDeadline(time.Now().Add(time.Second))

	fmt.Println(conn.LocalAddr(), conn.RemoteAddr())

	msg := make([]byte, 128)
	n, _ := conn.Read(msg)
	fmt.Println(string(msg[:n]))

	conn.Write([]byte("get it!"))
}
