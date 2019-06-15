package main

import (
	"io"
	"net"
)

func midServer() {
	var s1, _ = net.Listen("tcp", ":8080")
	defer s1.Close()
	var s2, _ = net.Listen("tcp", ":8081")
	defer s2.Close()
	for {
		var l1, _ = s1.Accept()
		var l2, _ = s2.Accept()
		go io.Copy(l2, l1)
		go io.Copy(l1, l2)
	}
}
