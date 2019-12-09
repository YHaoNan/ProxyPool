package getter

import (
	"HTTProxyPool/checker"
	"HTTProxyPool/model"
	"HTTProxyPool/scheduler"
	"HTTProxyPool/utils"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

type kxDailiGetter struct {

}


func (g kxDailiGetter) Run(result chan <- *model.Proxy){
	doc, err := utils.Get("http://www.kxdaili.com/dailiip.html")
	if err != "" {
		return
	}
	proxylist := make([]*model.Proxy,0)

	doc.Find(".hot-product-content tbody tr").Each(func(i int, selection *goquery.Selection) {
		var ip ,pos , isp string
		var port int
		var proxyType model.ProxyType = model.HTTP

		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			text := selection.Text()
			text = utils.TrimSpaceCharacter(text)
			switch i {
			case 0:
				ip = text
			case 1:
				port, _ = strconv.Atoi(text)
			case 5:
				pos = text
			case 3:
				if text == "HTTP" {
					proxyType = model.HTTP
				}else if text == "HTTPS" {
					proxyType = model.HTTPS
				}else if text == "HTTP,HTTPS" {
					proxyType = model.BOTH
				}
			}
		})

		proxy := &model.Proxy{
			ip,port,pos,isp,proxyType,
		}
		proxylist = append(proxylist, proxy)
	})
	checker.CheckAllAsyncAndPushToResult(result,proxylist)
}

func init(){
	var getter kxDailiGetter
	scheduler.Registe(getter)
}