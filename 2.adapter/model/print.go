package model

import (
	"2.adapter/model/banner"
)

// PrintBanner - object adapter
// 因為 golang 中並沒有繼承，所以這邊使用 composition(embedding struct) 來實作
type PrintBanner struct {
	banner.Banner
}

func (b *PrintBanner) PrintBanner(s string) {
	b.Banner.Banner(s)
}

func (b *PrintBanner) PrintWeek() {
	b.Banner.ShowWithParen()
}

func (b *PrintBanner) PrintStrong() {
	b.Banner.ShowWithAster()
}
