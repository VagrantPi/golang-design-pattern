package main

import (
	"2-2.object_adapter/model"
)

func main() {
	p := &model.PrintBanner{}
	p.PrintBanner("Hello")
	p.PrintWeek()
	p.PrintStrong()
}
