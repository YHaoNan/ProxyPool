package getter

import (
	"HTTProxyPool/checker"
	"HTTProxyPool/model"
	"HTTProxyPool/scheduler"
	"HTTProxyPool/utils"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

type _89FreeProxyGetter struct {

}

func (g _89FreeProxyGetter) Run(result chan <- *model.Proxy){

	doc, err := utils.Get("http://www.kxdaili.com/dailiip.html")
	if err != "" {
		return
	}

	proxylist := make([]*model.Proxy,0)

	doc.Find(".layui-table tbody tr").Each(func(i int, selection *goquery.Selection) {
		var ip ,pos , isp string
		var port int


		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			text := selection.Text()
			text = utils.TrimSpaceCharacter(text)

			switch i {
			case 0:
				ip = text
			case 1:
				port, _ = strconv.Atoi(text)
			case 2:
				pos = text
			case 3:
				isp = text
			}
		})

		proxy := &model.Proxy{
			ip,port,pos,isp,model.BOTH,
		}

		proxylist = append(proxylist, proxy)
	})

	checker.CheckAllAsyncAndPushToResult(result,proxylist)
}

func init(){
	var getter _89FreeProxyGetter
	scheduler.Registe(getter)
}
