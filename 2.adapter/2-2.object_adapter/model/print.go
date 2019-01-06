package model

import (
	"fmt"

	"2-2.object_adapter/model/banner"
)

// Print class
// golang 沒有 class，所以已 struct 實作，又因為 goalng 沒有 extend，所以這個 struct 沒有辦法被繼承，所以直接被我省略了
// type Print struct{}
// func (b *Print) PrintWeek()   {}
// func (b *Print) PrintStrong() {}

// PrintBanner extend Print
type PrintBanner struct {
	banner banner.Banner
}

func (b *PrintBanner) PrintBanner(s string) {
	fmt.Printf("PrintBanner Address: %p\n", b)
	bb := &banner.Banner{}
	b.banner = *bb.Banner(s)
}

func (b *PrintBanner) PrintWeek() {
	b.banner.ShowWithParen()
}

func (b *PrintBanner) PrintStrong() {
	b.banner.ShowWithAster()
}
