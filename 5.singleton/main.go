package main

import (
	"fmt"

	"5.singleton/model"
)

func main() {
	fmt.Println("start.")
	instanc1 := model.GetInstance()
	instanc1.Count()
	instanc1.Count()
	instanc2 := model.GetInstance()

	fmt.Println("instanc1 count is:", instanc1.Get())
	fmt.Println("instanc2 count is:", instanc2.Get())
	fmt.Println("End.")
}
