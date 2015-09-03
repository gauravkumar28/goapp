package  main

import ( 
	"fmt"
    "os"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	)


func main(){
    response := readfromDB()
    fmt.Println(response)
    write_to_html(response)    
}

func readfromDB() []string {
    res_arry := make([]string, 0, 10)
    db, err := sql.Open("mysql", "root"+":"+""+"@/"+"testgo")
    defer db.Close()
    stmtOuts, err := db.Prepare("SELECT response FROM runningstatuses")
    if err != nil {
        panic(err.Error())
    }
    defer stmtOuts.Close()
    rows, err := stmtOuts.Query()
    defer stmtOuts.Close()
    for rows.Next() {
        var response string
        rows.Scan(&response)
        res_arry = append(res_arry, response)
    }
    return res_arry
}

func write_to_html(response[] string){
    f, _ := os.Create("index.html")
    defer f.Close()
    f.Write([]byte("<html><head><title>Data</title></head><body>"))

    for _, element := range response {
        f.Write([]byte("<div>" + "<h1>" + element + "</h1></div><hr>"))
    }

    f.Write([]byte("</body></html>"))
}