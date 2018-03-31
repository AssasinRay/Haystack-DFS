package main

import (
	"fmt"
	"net"
	"encoding/gob"
)

type Request struct{
	Command string
	Url     string
	Photo   []byte
}


func  handleProxyConnection(proxyConn net.Conn ){
	dec := gob.NewDecoder(proxyConn)
	for{
		clientReq:= &Request{}
		dec.Decode(clientReq)
		fmt.Println(clientReq)
	}
}

func main() {
	fmt.Println("open Directory Server!")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("failed to listen!")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("failed to accept!")
			continue
		}
		fmt.Println("accept a proxy")
		go  handleProxyConnection(conn)
	}
}