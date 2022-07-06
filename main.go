package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/samber/lo"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

var PageMap = sync.Map{}
var detailList = []string{}

func main() {

	ch := make(chan os.Signal, 1)
	go func() {
		signal.Notify(ch, os.Kill, os.Interrupt)
	}()
	defer func() {
		<-ch
	}()

	var ul string
	bul, err := os.ReadFile(`地址.txt`)
	if err != nil {
		panic(err)
	}
	ul = strings.TrimSpace(string(bul))
	fmt.Println("开始工作，等待")

	c := colly.NewCollector(colly.Async(false))
	detailC := colly.NewCollector(colly.Async(false))
	detailC.OnHTML(`th:contains("Is this an Adverse Event Report?")`, func(element *colly.HTMLElement) {
		r := strings.TrimSpace(element.DOM.Next().Text())
		if r == "Yes" {
			fmt.Printf("%s\n", element.Request.URL.String())
		} else {
			//fmt.Printf("%s:%s\n",element.Request.URL.String(),"No")
		}
	})
	detailC.OnRequest(func(request *colly.Request) {
		time.Sleep(time.Millisecond * 100)
	})

	c.OnHTML(`a[href*="detail.cfm?"]`, func(element *colly.HTMLElement) {
		href := element.Attr(`href`)

		_, b := lo.Find(detailList, func(s string) bool {
			return href == s
		})
		if !b {
			detailList = append(detailList, href)
			err := detailC.Visit(element.Request.AbsoluteURL(href))
			if err != nil {
				panic(err)
			}
		}
	})
	c.OnHTML(`a[href*="results.cfm?"]`, func(element *colly.HTMLElement) {
		href := element.Attr(`href`)
		_, ld := PageMap.LoadOrStore(href, "1")
		if !ld {
			element.Request.Visit(href)
		}
	})
	err = c.Visit(ul)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("一共搜索到记录:%d条\n", len(detailList))
}
