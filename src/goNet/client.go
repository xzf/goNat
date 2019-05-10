package goNet

import (
	"fmt"
	"net"
)

type NatClient struct {
	LocalAddr  string
	RemoteAddr string
	conn       net.Conn
}

func NewNatClient(local string, remote string) *NatClient {
	client := &NatClient{
		LocalAddr:  local,
		RemoteAddr: remote,
	}
	conn, err := net.Dial("tcp4", remote)
	if err != nil {
		fmt.Println("vw6yq5n6ru", err)
		return nil
	}
	_, err = conn.Write([]byte("666"))
	if err != nil {
		fmt.Println("1xossbbd75", err)
		return nil
	}
	return client
}

func (c *NatClient) ReadThread() {
	//req := []string{}
	//for {
	//	c.conn.Read()
	//
	//}
}
