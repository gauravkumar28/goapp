
package  main

import ( 
	"fmt"
	"github.com/xuyu/goredis"

	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	)


func main(){
	client, _ := goredis.Dial(&goredis.DialConfig{Address: "127.0.0.1:6379"})
    for {
	reply, _ := client.RPop("myqueue1")

    if string(reply) !="" {
    	storeinDB(string(reply))
	fmt.Println(string(reply));
    }
  }
}

func storeinDB(data string){
   db, err := sql.Open("mysql", "root"+":"+""+"@/"+"testgo")
   defer db.Close()
    stmtIns, err := db.Prepare("INSERT INTO runningstatuses (response) VALUES(?)")
    if err != nil {
        panic(err.Error())
    }
    stmtIns.Exec(data)
    defer stmtIns.Close()
}