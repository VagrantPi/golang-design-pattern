package main

import (
	"2.adapter/model"
)

func main() {
	pb := &model.PrintBanner{}
	pb.PrintBanner("Hello")
	// p (規格 A) 透過 PrintBanner(adapter) 使用 Banner(規格 B)的方法
	pb.PrintWeek()
	pb.PrintStrong()
}
