package main

import (
	"fmt"
	"net"
	"os"
	"encoding/gob"
	"io/ioutil"
)

type Request struct{
	Command string
	Url     string
	Photo   []byte
}

func main(){
	fmt.Println("Welcome to use Haystack!")
	conn, err := net.Dial("tcp","localhost:8080")
	if err != nil{
		fmt.Println("failed to connect proxy")
		os.Exit(3)
	}
	encoder := gob.NewEncoder(conn)
	for{
		fmt.Println("type your command fecth/upload/delete")
		var input string
		fmt.Scanln(&input)
		var url string
		switch input {
		case "fetch":
			fmt.Println("type your photo url")
			fmt.Scanln(&url)
			clientReq := &Request{"fetch",url,nil}
			encoder.Encode(clientReq)
		case "upload":
			fmt.Println("type your photo url")
			fmt.Scanln(&url)
			fmt.Println("type your photo path with name")
			var photo_addr string
			fmt.Scanln(&photo_addr)
			photoData,err := ioutil.ReadFile(photo_addr)
			if err != nil {
				fmt.Println("read photo failed")
			}
			clientReq := &Request{"upload",url,photoData}
			encoder.Encode(clientReq)
		case "delete":
			fmt.Println("type your photo url")
			fmt.Scanln(&url)
			clientReq := &Request{"delete",url,nil}
			encoder.Encode(clientReq)
		case "exit":
			fmt.Println("exit the haystack")
			conn.Close()
			os.Exit(0)
		default:
			fmt.Println("invalid command!")
		}
	}
}