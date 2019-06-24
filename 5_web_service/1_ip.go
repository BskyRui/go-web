package main

import (
	"fmt"
	"net"
)

func main() {
	var (
		name string = "192.168.1.1"
	)
	addr := net.ParseIP(name)
	fmt.Println(addr, addr.String())
}
