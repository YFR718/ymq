package net

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/YFR718/ymq/internal/system"
	"github.com/YFR718/ymq/internal/topic"
	"github.com/YFR718/ymq/pkg/common"
	"io"
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
				// 读取4字节长度
				// 读取数据包头，通常包头包含了数据包的长度信息
				length := uint32(0)
				err := binary.Read(conn, binary.BigEndian, &length)
				if err != nil {
					if err == io.EOF {
						println("断开连接")
						break
					}
					common.PrintError(err)
					break
				}

				// 读取length长度数据
				// 根据数据包头中的长度信息，从 TCP 连接中读取相应长度的数据
				data := make([]byte, length)
				_, err = conn.Read(data[4:])
				if err != nil {
					common.PrintError(err)
					break
				}
				binary.BigEndian.PutUint32(data[:4], length)

				message, err := common.Unmarshal(data)
				if err != nil {
					common.PrintError(err)
					break
				}

				switch message.Header.Type {
				case common.CREATE_TOPIC:
					t := topic.Topic{}
					_ = json.Unmarshal(message.Body, &t)
					err := topic.TopicManagerInstance.Create(t)
					if err != nil {
						common.PrintError(err)
						break
					}
				case common.SEND_MESSAGE:
					t := topic.Topic{}
					_ = json.Unmarshal(message.Body, &t)
					err := topic.TopicManagerInstance.Send(t)
					if err != nil {
						common.PrintError(err)
						break
					}
				case common.DELETE_TOPIC:
					t := topic.Topic{}
					_ = json.Unmarshal(message.Body, &t)
					topic.TopicManagerInstance.Delete(t)
					if err != nil {
						common.PrintError(err)
						break
					}

				}
				//println("get a message:", string(message.Body))

				// 回复数据

				header := common.Header{Type: common.PONG}
				msg := common.Message{Header: header, Body: []byte("success")}

				s := msg.Marshal()

				write, err := conn.Write(s)
				if err != nil {
					common.PrintError(err)
					println(write)
					return
				}
				println("YES")
			}

		}(conn)

	}
}
