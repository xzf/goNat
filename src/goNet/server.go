package goNet

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

//255,255,255,255,
//4+8+real data
//addr+data len+real data

const DefaultAddr = ":8888"

type NatServer struct {
	lister    net.Listener
	natClient net.Conn
	clientMap map[string]net.Conn
}

func NewNatServer(addr string) *NatServer {
	lister, err := net.Listen("tcp4", addr)
	if err != nil {
		fmt.Println("oflpmzsvxk", err)
		return nil
	}
	return &NatServer{
		lister:    lister,
		clientMap: map[string]net.Conn{},
	}
}

func (s *NatServer) ListenThread() {
	for {
		conn, err := s.lister.Accept()
		if err != nil {
			fmt.Println("nnzpp77suw", err)
			continue
		}
		id := conn.RemoteAddr().String()
		if s.clientMap == nil {
			s.clientMap = map[string]net.Conn{}
		}
		s.clientMap[id] = conn
		go func(con net.Conn) {
			req := make([]byte, 2048000)
			for {
				n, err := con.Read(req)
				if err != nil {
					fmt.Println("4e3hboer09", err)
					continue
				}
				if s.natClient.RemoteAddr().String() == "" &&
					n == 3 && string(req[:3]) == "666" {
					s.natClient = con
				}
				_, err = s.natClient.Write(req[:n])
				if err != nil {
					fmt.Println("ip0gnwiluo", err)
					continue
				}
				time.Sleep(time.Millisecond * 10)
			}
		}(conn)
	}
}

func (s *NatServer) RespThread() {
	for {
		for s.natClient.RemoteAddr().String() == "" {
			time.Sleep(time.Millisecond * 500)
		}
		lengthByte := make([]byte, 8)
		_, err := s.natClient.Read(lengthByte)
		if err != nil {
			fmt.Println("hv6rs20u5v", err)
			continue
		}
		length, err := strconv.ParseInt(string(lengthByte), 10, 64)
		if err != nil {
			fmt.Println("ddvk2dpkej", err)
			continue
		}
		dataByte := make([]byte, length)
		n, err := s.natClient.Read(dataByte)
		if err != nil {
			fmt.Println("hv6rs20u5v", err)
			continue
		}
		for n < int(length) {
			length -= int64(n)
			dataByteTmp := make([]byte, length)
			n, err = s.natClient.Read(dataByteTmp)
			if err != nil {
				fmt.Println("el5ro88rrp", err)
				break
			}
			dataByte = append(dataByte, dataByteTmp...)
		}
	}
}
