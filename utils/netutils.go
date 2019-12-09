package utils

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func Get(url string)  (*goquery.Document, string){
	resp, err := http.Get(url)
	if err != nil {
		return nil , "Request faild"
	}else if resp.StatusCode != 200 {
		return nil, "Response code is not 200"
	}

	doc , err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, "Response is not vaild!"
	}

	defer resp.Body.Close()
	return doc,""
}
