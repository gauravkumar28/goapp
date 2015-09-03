// web server

package  main

import ( 
	"fmt"
	//"bytes"
	"gopkg.in/olivere/elastic.v2"
	"os"
	)

func main() {
	response := getFromES()
    write_to_html(response)    

}




func getFromES() [] string{
	res_arry := make([]string, 0, 10)
	client, err := elastic.NewClient()
	  if err != nil {
	    // Handle error
	    panic(err)
	  }

	   //termQuery := elastic.NewTermQuery("ReqId", "2")
	  searchResult, err := client.Search().
	    Index("cardata").   
	    //Query(&termQuery).  
	    From(0).Size(10).   
	    Pretty(true).       
	    Do()                
	  if err != nil {
	    // Handle error
	    panic(err)
	  }
  //fmt.Println(searchResult)
  fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

  for _, hit := range searchResult.Hits.Hits{
      //fmt.Println(hit.Source)
      res_arry = append(res_arry, string(((*(hit.Source)))))
      }
  return res_arry
}

func write_to_html(response[] string){
    f, _ := os.Create("index2.html")
    defer f.Close()
    f.Write([]byte("<html><head><title>Data</title></head><body>"))

    for _, element := range response {
        f.Write([]byte("<div>" + "<h1>" + element + "</h1></div><hr>"))
    }

    f.Write([]byte("</body></html>"))
}