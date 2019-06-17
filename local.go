package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

var MID_SERVER = ""
var MID_SERVER_LOCAL_SIDE = ":8080"
var MID_SERVER_REMOTE_SIDE = ":8081"
var SERVER = "http://localhost:8082"

func handle(l net.Conn) {
	defer l.Close()
	var l1, _ = net.Dial("tcp", MID_SERVER+MID_SERVER_LOCAL_SIDE)
	var req, _ = http.NewRequest("GET", SERVER, bytes.NewBufferString(MID_SERVER+MID_SERVER_REMOTE_SIDE))
	var client = &http.Client{}
	var res, _ = client.Do(req)
	defer res.Body.Close()
	var b, _ = ioutil.ReadAll(res.Body)
	fmt.Println(string(b))
	go io.Copy(l1, l)
	io.Copy(l, l1)
}
func localServer() {
	var s, _ = net.Listen("tcp", ":1080")
	defer s.Close()
	for {
		var l, _ = s.Accept()
		go handle(l)
	}
}
