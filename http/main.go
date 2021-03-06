package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Println("Hello Shant")

	//net = network
	// nl, err := net.Listen("tcp", "127.0.0.1:8888") //ip:port   127.0.0.1:8888
	nl, err := net.Listen("tcp", ":8080") //ip:port   127.0.0.1:8888     //port limit 0-65535
	if err != nil {
		fmt.Println("1")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for {
		conn, err := nl.Accept() //Layer-5 session layer
		if err != nil {
			fmt.Println("2")
			fmt.Println(err.Error())
			// os.Exit(1)
		}

		// remoteAddr := conn.RemoteAddr().String()
		// fmt.Println(remoteAddr) //[::1]:40606

		bs := make([]byte, 1024)

		i, err := conn.Read(bs)
		if err != nil {
			fmt.Println("3")
			fmt.Println(err.Error())
			// os.Exit(1)
		}

		fmt.Println(i)

		// conn.Write([]byte("welcome we have recieve your message \n"))

		reqstr := string(bs)
		fmt.Println(reqstr)

		body := `<!DOCTYPE html><html><head><title>Page Title</title></head><body><h1>My First Heading</h1><p>My first paragraph.</p></body></html>`

		resp := "HTTP/1.1 200 OK\r\n" +
			"Content-Length: %d\r\n" +
			"Content-Type: test/html\r\n" +
			"\r\n%s"

		msg := fmt.Sprintf(resp, len(body), body)
		fullMsg := resp + body
		fmt.Println(msg)
		conn.Write([]byte(fullMsg))

		// conn.Close()
		// nl.Close()

	}
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Test")
// 	fmt.Println("Connect Successful")
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":8080", nil)
// }
