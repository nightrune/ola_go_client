package main

import (
	"./ola_proto"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
	}
}

func Something() {
	fmt.Printf("Here there something\n")
}

func readForever(conn *net.TCPConn) {
	var buf [512]byte

	for {
		_, err := conn.Read(buf[0:])
		if err != nil {
			nerr, ok := err.(net.Error)
			if ok == true {
				if nerr.Timeout() {
					fmt.Printf("Read timed out\n")
				} else {
					fmt.Printf("Connection was closed during read..\n")
					return
				}
			} else {
				checkError(err)
				return
			}
		} else {
			fmt.Printf("Data: %s\n", buf)
		}
	}
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp4", "localhost:9010")
	checkError(err)
	fmt.Printf("Address: %s\n", addr)
	conn, err := net.DialTCP("tcp", nil, addr)
	checkError(err)
	conn.SetNoDelay(true)
	go readForever(conn)
	x := &ola_proto.PluginListRequest{}
	data, err := proto.Marshal(x)
	checkError(err)
	fmt.Printf("Test data: %s\n", data)
	_, err = conn.Write(data)
	checkError(err)
	time.Sleep(3000 * time.Millisecond)
	conn.Close()
	time.Sleep(3000 * time.Millisecond)
}
