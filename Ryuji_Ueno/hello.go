package main

import(
	"encoding/json"
	"net/http"
	"log"
	"fmt"
)

type Message struct{
	Result string
}

func main(){

	message := Message{"Hello World"}
	bdata := EncodingJSON(message)
        http.ListenAndServe("5000",nil)
	handleRequests()
}

func EncodingJSON(mes Message)[]byte{
	bdata,err := json.Marshal(mes)
	if err != nil{
		fmt.Println(err)
		return nil
	}
	return bdata
}

func homePage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
   	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests(){
	http.HandleFunc("/",EncodingJSON(Message))
	log.Fatal(http.ListenAndServe(":5000",nil))
}

