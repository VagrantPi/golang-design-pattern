package banner

import "fmt"

type Banner struct {
	bannerString string
}

func (b *Banner) Banner(s string) {
	b.bannerString = s
}

func (b *Banner) ShowWithParen() {
	fmt.Printf("(%v)\n", b.bannerString)
}

func (b *Banner) ShowWithAster() {
	fmt.Printf("*%v*\n", b.bannerString)
}
