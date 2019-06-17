package handler

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
)

func H(w http.ResponseWriter, r *http.Request) {
	var midUrl, _ = ioutil.ReadAll(r.Body)
	var client, _ = net.Dial("tcp", string(midUrl))
	fmt.Fprintln(w, "Connection established.")
	go func() {
		defer client.Close()
		if client == nil {
			return
		}
		defer client.Close()

		var b [32768]byte
		n, err := client.Read(b[:])
		// fmt.Println(string(b[:n]))
		if err != nil {
			log.Println(err)
			return
		}
		var method, host, address string
		fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
		hostPortURL, err := url.Parse(host)
		if err != nil {
			log.Println(err)
			return
		}

		if hostPortURL.Opaque == "443" { //https访问
			address = hostPortURL.Scheme + ":443"
		} else { //http访问
			if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
				address = hostPortURL.Host + ":80"
			} else {
				address = hostPortURL.Host
			}
		}

		//获得了请求的host和port，就开始拨号吧
		server, err := net.Dial("tcp", address)
		if err != nil {
			log.Println(err)
			return
		}
		if method == "CONNECT" {
			fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
		} else {
			server.Write(b[:n])
		}
		//进行转发
		go io.Copy(server, client)
		io.Copy(client, server)
	}()
}
