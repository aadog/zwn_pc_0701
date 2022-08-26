// 由res2go IDE插件自动生成。
package main

import (
	"github.com/ying32/govcl/vcl"
)

func main() {
	//设置工程缩放
	vcl.Application.SetScaled(true)
	//设置工程窗口标题
	vcl.Application.SetTitle("project1")
	//初始化工程
	vcl.Application.Initialize()
	//设置任务栏显示
	vcl.Application.SetMainFormOnTaskBar(true)
	//创建窗口
	vcl.Application.CreateForm(&Form1)
	//运行工程
	vcl.Application.Run()
}
