package checker

import (
	"HTTProxyPool/model"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var testSite = "http://www.baidu.com"

func CheckAllAsyncAndPushToResult(result chan <- *model.Proxy ,proxylist []*model.Proxy){
	var wg sync.WaitGroup
	wg.Add(len(proxylist))

	for _,proxy := range proxylist{
		go func(p *model.Proxy) {
			canUse := Check(p)
			if canUse {
				result <- p
			}else{
				log.Println("Unuseable proxy : "+p.Address)
			}
			wg.Done()
		}(proxy)
	}

	wg.Wait()
}

func Check(proxy *model.Proxy) bool{
	transport := &http.Transport{Proxy: func(request *http.Request) (url *url.URL, err error) {
		var schema string
		switch proxy.Type {
		case model.HTTP:
			schema = "http"
		case model.HTTPS:
			schema = "https"
		case model.BOTH:
			schema = "http"
		}
		return url.Parse(fmt.Sprintf("%s://%s:%d",schema,proxy.Address,proxy.Port))
	}}

	client := &http.Client{Timeout: 10 * time.Second,Transport:transport}
	_, err := client.Get(testSite)
	if err != nil {
		return false
	}
	return true
}