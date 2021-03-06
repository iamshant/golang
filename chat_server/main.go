package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

type rectangle struct {
	length  float64
	breadth float64
	color   string
}

func main() {

	fmt.Println("Chat server")

	//net = network
	// nl, err := net.Listen("tcp", "127.0.0.1:8888") //ip:port   127.0.0.1:8888
	nl, err := net.Listen("tcp", ":8888") //ip:port   127.0.0.1:8888     //port limit 0-65535
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for {
		conn, err := nl.Accept() //Layer-5 session layer
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		remoteAddr := conn.RemoteAddr().String()
		fmt.Println(remoteAddr) //[::1]:40606

		bs := make([]byte, 1024)
		i, err := conn.Read(bs)
		if err != nil {
			fmt.Println("3")
			fmt.Println(err.Error())
			// os.Exit(1)
		}
		fmt.Println(string(bs[:i]))
		// fmt.Println(i)

		rcvTime := time.Now().Format("2006-01-02 15:04:05")
		msg := fmt.Sprintf("Welcome, we have recieve your message: '%s' at %s.\n", bs[:i], rcvTime)

		conn.Write([]byte(msg))

		// conn.Close()
		// nl.Close()

	}
}
