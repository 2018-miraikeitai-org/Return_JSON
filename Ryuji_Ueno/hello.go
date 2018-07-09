package main

import(
	"encoding/json"
	"net/http"
)

type Message struct{
	Result string `json:"Result"`
}

func main(){

	http.HandleFunc("/",index)
	http.ListenAndServe(":5000",nil)
}

func index(w http.ResponseWriter, r *http.Request){
	result := Message{ Result:"Hello World"}
	
	res, err := json.Marshal(result)

	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Write(res)
}

