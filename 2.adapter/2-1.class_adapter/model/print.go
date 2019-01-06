package model

import (
	"fmt"

	"2-1.class_adapter/model/banner"
)

type Print interface {
	PrintWeek()
	PrintStrong()
}

// PrintBanner 繼承 Banner 並實作 Print
// 因為 golang 中並沒有繼承，所以這邊使用 Embedding struct 來實作
type PrintBanner struct {
	banner.Banner
}

func (b *PrintBanner) PrintBanner(s string) {
	fmt.Printf("PrintBanner Address: %p\n", b)
	b.Banner.Banner(s)
}

func (b *PrintBanner) PrintWeek() {
	b.Banner.ShowWithParen()
}

func (b *PrintBanner) PrintStrong() {
	b.Banner.ShowWithAster()
}
