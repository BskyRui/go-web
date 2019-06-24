package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:1200")
	// connect
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	conn.Write([]byte("timestamp"))
	result, _ := ioutil.ReadAll(conn)
	fmt.Println(string(result))
}
