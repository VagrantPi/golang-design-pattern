package model

import "fmt"

type AbstractDisplay interface {
	Open()
	Print()
	Close()
}

type AbstractDisplayTemplate interface {
	Display(AbstractDisplay)
}

type AbstractDisplayTemplateStruct1 struct{}

func (a *AbstractDisplayTemplateStruct1) Display(ad AbstractDisplay) {
	ad.Open()
	ad.Print()
	ad.Close()
}

type AbstractDisplayTemplateStruct2 struct{}

func (a *AbstractDisplayTemplateStruct2) Display(ad AbstractDisplay) {
	ad.Open()
	for i := 0; i < 5; i++ {
		ad.Print()
	}
	ad.Close()
}

type CharDisplate struct {
	Char string
}

func (c CharDisplate) Open() {
	fmt.Printf("<<")
}

func (c CharDisplate) Print() {
	// Char 為 string，不過這個方法只印出第一個字元，所以多了下面的判斷式
	if len(c.Char) != 0 {
		fmt.Printf(c.Char[0:1])
		return
	}
	fmt.Printf("")
}

func (c CharDisplate) Close() {
	fmt.Println(">>")
}

type StringDisplate struct {
	String string
}

func (c StringDisplate) Open() {
	c.PrintLine()
}

func (c StringDisplate) Print() {
	fmt.Println("+ " + c.String + " +")
}

func (c StringDisplate) Close() {
	c.PrintLine()
}

func (c StringDisplate) PrintLine() {
	fmt.Printf("+")
	for i := 0; i < len(c.String)+2; i++ {
		fmt.Printf("-")
	}
	fmt.Println("+")
}
