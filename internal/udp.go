package main

import (
	"log"
	"net"
	"time"
)

func main()  {
	addr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9000}
	udp, err := net.DialUDP(addr.Network(), addr, addr)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < 100; i++ {
		go func() {
			n, err := udp.Write([]byte("abcd"))
			if err != nil {
				log.Println(err)
			}
			log.Println("发送", n)
		}()
	}

	time.Sleep(10 * time.Second)

}
