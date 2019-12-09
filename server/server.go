package server

import (
	"HTTProxyPool/model"
	"encoding/json"
	"fmt"
	"net/http"
)

var results []model.Proxy

func getProxies(w http.ResponseWriter,r *http.Request){
	if data , err := json.Marshal(results); err!=nil {
		fmt.Fprintf(w,"{\"error\":\"%s\"}",err.Error())
	}else{
		fmt.Fprint(w,string(data))
	}
}
func AppendResult(proxy model.Proxy){
	results = append(results, proxy)
}
func StartServer() error{
	http.HandleFunc("/", getProxies)
	err := http.ListenAndServe(":4000",nil)
	return err
}