package main

import (
	"fmt"

	"3.template/model"
)

func main() {
	var data1 model.AbstractDisplay = &model.CharDisplate{"H"}
	var data2 model.AbstractDisplay = &model.StringDisplate{"Hello, world!"}
	var data3 model.AbstractDisplay = &model.StringDisplate{"Template Pattern!"}
	var tmpl1 model.AbstractDisplayTemplate = &model.AbstractDisplayTemplateStruct1{}
	var tmpl2 model.AbstractDisplayTemplate = &model.AbstractDisplayTemplateStruct2{}
	fmt.Println("template 1")
	tmpl1.Display(data1)
	tmpl1.Display(data2)
	tmpl1.Display(data3)
	fmt.Println("\n==============================\n")
	fmt.Println("template 2")
	tmpl2.Display(data1)
	tmpl2.Display(data2)
	tmpl2.Display(data3)
}
