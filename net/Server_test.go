package net

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {

	s := NewServer("[liu li games V0.3]")

	s.Serve()
}

func TestClient(t *testing.T) {
	fmt.Println("client start...")
	time.Sleep(1 + time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:9527")
	if err != nil {
		fmt.Println("client start err", err)
		return
	}

	for {
		_, err := conn.Write([]byte("Hello LiuLiGames V0.2 "))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err", err)
			return
		}

		fmt.Printf("server call back: %s ,cnt = %d \n", buf, cnt)

		time.Sleep(1 * time.Second)
	}
}
