package scheduler

import (
	"HTTProxyPool/model"
	"HTTProxyPool/server"
	"sync"
)

var getterList = make([]Getter,0)

func Registe(g Getter){
	getterList = append(getterList, g)
}

func Run(){

	var wg sync.WaitGroup

	wg.Add(len(getterList))
	result := make(chan *model.Proxy)
	for _, g := range getterList{
		go func(getter Getter) {
			getter.Run(result)
			wg.Done()
		}(g)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	Display(result)
}

func Display(result chan *model.Proxy) {
	for r := range result{
		server.AppendResult(*r)
	}
}