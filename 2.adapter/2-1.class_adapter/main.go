package main

import (
	"2-1.class_adapter/model"
)

func main() {
	pb := &model.PrintBanner{}
	pb.PrintBanner("Hello")
	var p model.Print = pb
	// p (規格 A) 透過 PrintBanner(adapter) 使用 Banner(規格 B)的方法
	p.PrintWeek()
	p.PrintStrong()
}
