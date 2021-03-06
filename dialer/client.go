package main

import (
	"fmt"
	"net"
	"os"
)

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

func main() {

	fmt.Println("Client")

	//net = network
	// nl, err := net.Listen("tcp", "127.0.0.1:8888") //ip:port   127.0.0.1:8888
	nl, err := net.Dial("tcp", ":8888") //ip:port   127.0.0.1:8888     //port limit 0-65535
	if err != nil {
		fmt.Println("1")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	nl.Write([]byte("hello"))

	bs := make([]byte, 1024)
	i, _ := nl.Read(bs)
	fmt.Println(string(bs[:i]))
	nl.Close()

}
