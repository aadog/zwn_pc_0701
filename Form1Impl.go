package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/samber/lo"
	"github.com/ying32/govcl/vcl"
	"net/url"
	"strings"
	"sync"
	"test.com/a/vcladv"
	"time"
)

// ::private::
type TForm1Fields struct {
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	code := f.Edit1.Text()
	date1 := f.DateTimePicker1.Date().Format(`01/02/2006`)
	date2 := f.DateTimePicker2.Date().Format(`01/02/2006`)
	d := f.SpinEdit1.Value()

	f.Button1.SetEnabled(false)
	strUrl := fmt.Sprintf(`https://www.accessdata.fda.gov/scripts/cdrh/cfdocs/cfMAUDE/results.cfm?start_search=1&productcode=%s&productproblem=&patientproblem=&devicename=&modelNumber=&reportNumber=&manufacturer=&brandname=&eventtype=&reportdatefrom=%s&reportdateto=%s&pagenum=100000`, code, url.QueryEscape(date1), url.QueryEscape(date2))
	vcladv.NewBackgroundTask().SetDeferFn(func() {
		f.Button1.SetEnabled(true)
	}).SetErrorFn(func(err error) {
		vcl.ShowMessage(err.Error())
	}).Run(func() error {
		var PageMap = sync.Map{}
		var detailList = []string{}

		c := colly.NewCollector(colly.Async(false))
		detailC := colly.NewCollector(colly.Async(false))
		detailC.OnHTML(`th:contains("Is this an Adverse Event Report?")`, func(element *colly.HTMLElement) {
			r := strings.TrimSpace(element.DOM.Next().Text())
			if r == "Yes" {
				vcl.ThreadSync(func() {
					f.Memo1.Lines().Append(element.Request.URL.String())
				})
			} else {
				//fmt.Printf("%s:%s\n",element.Request.URL.String(),"No")
			}
		})
		detailC.OnRequest(func(request *colly.Request) {
			time.Sleep(time.Millisecond * time.Duration(d))
		})

		c.OnHTML(`a[href*="detail.cfm"]`, func(element *colly.HTMLElement) {
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
				detailC.Wait()
			}
		})
		c.OnHTML(`a[href*="results.cfm"]`, func(element *colly.HTMLElement) {
			href := element.Attr(`href`)
			_, ld := PageMap.LoadOrStore(href, "1")
			if !ld {
				element.Request.Visit(href)
			}
		})
		err := c.Visit(strUrl)
		if err != nil {
			fmt.Println(err)
		}
		c.Wait()
		return nil
	})
}
