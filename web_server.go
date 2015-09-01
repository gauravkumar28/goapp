// web server

package  main

import ( 
	"fmt"
	"net"
	"os"
	"bytes"
	)

const ( 
CONN_HOST = ""
CONN_PORT = "3333"
CONN_TYPE = "tcp"
)

func main() {
	// listning

	listen, error := net.Listen(CONN_TYPE, ":"+CONN_PORT)
	if error != nil {
		fmt.Println("Error in Listening")
		os.Exit(1)
	}
   // close listner when server down
	defer listen.Close()

	fmt.Println("Listing on" + CONN_PORT)

	for  {
		conn, error := listen.Accept()
		if error  != nil {
			fmt.Println("Error Acception failed")
			os.Exit(1)
		 }
	
	//fmt.Println("Message Received")
	// create new routine to handle connection
        go handleRequest(conn)
   }
}


func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)

    n := bytes.Index(buf, []byte{0})
	if err != nil {
		fmt.Println("Error reading:", err)
	}

	fmt.Println(reqLen)
	message := "Hi , I received your message "
	message += string(buf[:n-1])
    //storeinRadis(message)
	conn.Write([]byte(message))
	conn.Close()
}