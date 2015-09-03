// web server

package  main

import ( 
    "encoding/json"
	"fmt"
	"net"
	"os"
	"bytes"
	"github.com/xuyu/goredis"
	)

const ( 
CONN_HOST = ""
CONN_PORT = "3333"
CONN_TYPE = "tcp"
) 

type response struct {
	ReqId int `json:id`
    DeviceId int `json:id`
    Payload string `json:payload`
}

type DevReqKey struct{
	ReqId int
	DeviceId int
}


type CarData struct {
	ReqId int 
	Payload string
}

func main() {

	listen, error := net.Listen(CONN_TYPE, ":"+CONN_PORT)
	if error != nil {
		fmt.Println("Error in Listening")
		os.Exit(1)
	}
	defer listen.Close()

	fmt.Println("Listing on" + CONN_PORT)

	for  {
		conn, error := listen.Accept()
		if error  != nil {
			fmt.Println("Error Acception failed")
			os.Exit(1)
		 }
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
    storeinRadis(string(buf[:n-1]))
	message := "Hi , I received your message "
	message += string(buf[:n-1])
	conn.Write([]byte(message))
	conn.Close()
}




func storeinRadis(msg string){
	client, _ := goredis.Dial(&goredis.DialConfig{Address: "127.0.0.1:6379"})
	reply, _ := client.ExecuteCommand("LPUSH", "myqueue1", msg)
	fmt.Println(reply)
}



func getFromJson(buf[] byte,n int){
	var res response
	err := json.Unmarshal(buf[:n-1], &res)

	fmt.Println(err)
	fmt.Println(res.Payload)
	fmt.Println(res.ReqId)
	fmt.Println(res.DeviceId)
}