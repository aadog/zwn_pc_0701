package main

import (
	"fmt"
	"github.com/aadog/tasks-go/backgroundTask"
	"github.com/gocolly/colly/v2"
	"github.com/samber/lo"
	"github.com/ying32/govcl/vcl"
	"net/url"
	"strings"
	"sync"
	"time"
)

// ::private::
type TForm1Fields struct {
}

// 按钮点击
func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	//获取用户输入窗口代码
	code := f.Edit1.Text()
	//获取用户输入开始时间
	date1 := f.DateTimePicker1.Date().Format(`01/02/2006`)
	//获取用户输入结束时间
	date2 := f.DateTimePicker2.Date().Format(`01/02/2006`)
	//获取用户输入 任务间隔 /毫秒为单位(网站有限制，抓取过快会封ip,如果购买大量代理ip不需要间隔)
	d := f.SpinEdit1.Value()

	//按钮设置为不可用状态
	f.Button1.SetEnabled(false)
	//抓取的url地址
	strUrl := fmt.Sprintf(`https://www.accessdata.fda.gov/scripts/cdrh/cfdocs/cfMAUDE/results.cfm?start_search=1&productcode=%s&productproblem=&patientproblem=&devicename=&modelNumber=&reportNumber=&manufacturer=&brandname=&eventtype=&reportdatefrom=%s&reportdateto=%s&pagenum=100000`, code, url.QueryEscape(date1), url.QueryEscape(date2))
	//后端运行任务，保持界面流畅
	backgroundTask.New().SetDeferFn(func() {
		//任务结束后按钮设置为可用
		f.Button1.SetEnabled(true)
	}).SetErrorFn(func(err error) {
		//如果发生错误就报错
		vcl.ShowMessage(err.Error())
	}).Run(func() error {
		var PageMap = sync.Map{}
		var detailList = []string{}

		//建立列表页面页面抓取同步控制器
		c := colly.NewCollector(colly.Async(false))
		//建立详情页抓取同步控制器
		detailC := colly.NewCollector(colly.Async(false))
		//抓取到详情页面后分析详情页数据,是否包含Is this an Adverse Event Report?:Yes,如果有就把他添加到结果里
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
		//每个抓取页面前延迟 毫秒*用户输入的毫秒,用户防ip被封
		detailC.OnRequest(func(request *colly.Request) {
			time.Sleep(time.Millisecond * time.Duration(d))
		})

		//抓取到列表页面后分析里面有几条信息详情页信息,并且任务提交到详情页控制器
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
		//抓取到列表页后看是否需要翻页，如果需要翻页就一页一页抓取数据
		c.OnHTML(`a[href*="results.cfm"]`, func(element *colly.HTMLElement) {
			href := element.Attr(`href`)
			_, ld := PageMap.LoadOrStore(href, "1")
			if !ld {
				element.Request.Visit(href)
			}
		})
		//开始运行列表页控制器
		err := c.Visit(strUrl)
		if err != nil {
			fmt.Println(err)
		}
		//等待任务结束
		c.Wait()
		return nil
	})
}
