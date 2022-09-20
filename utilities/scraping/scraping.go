package scraping

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func Scrap() {

	resp, err := http.Get("https://free.proxy-sale.com/")
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//spr := fmt.Sprintf("%s", doc)
	//fmt.Println(spr)
	links := make([]string, 0)
	//<span class="your-ip__numbers">188.130.172.26</span>
	//body > div:nth-child(6) > section.hide-ip-table
	if doc.Type == html.ElementNode && doc.Data == "href" {
		for _, v := range doc.Attr {

			if v.Key == "//free.proxy-sale.com/my-ip/" {
				links = append(links, v.Val)
				//fmt.Println(doc)

			}
		}
	}
	fmt.Println(links)
}
