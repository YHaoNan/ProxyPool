package main

import (
	_ "HTTProxyPool/getter"
	"HTTProxyPool/scheduler"
	"HTTProxyPool/server"
)

func main(){
	go scheduler.Run()
	server.StartServer()
}
