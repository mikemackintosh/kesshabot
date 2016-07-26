package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/miekg/dns"
)

func startDNS() {
	addr, _ := net.ResolveUDPAddr("udp", ":53")

	conn, err := net.ListenUDP("udp", addr)
	defer conn.Close()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}

	for {
		m := make([]byte, 4096)
		n, _, _, raddr, err := conn.ReadMsgUDP(m, m)
		if err != nil {
			fmt.Println("Error: ", err)
			conn.Write([]byte(`error`))
			continue
		}

		req := new(dns.Msg)
		err = req.Unpack(m[:n])
		if err != nil { // Send a FormatError back
			fmt.Println("Error: ", err)
		}
		conn.Write([]byte(`Your mom`))

		msg := fmt.Sprintf("DNS Query Attempted:\n💻 %s\n%s\n", raddr, req.Question[0].String())
		log.Println(msg)
		_, _, err = twClient.Client.Statuses.Update(msg, nil)
		if err != nil {
			log.Printf("[Error] pubkeyAuth: %s\n", err)
		}
	}
}
