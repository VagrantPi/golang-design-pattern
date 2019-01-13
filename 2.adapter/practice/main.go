package main

import (
	"practice/model"
)

func main() {
	var fp model.FileIO = &model.FileProperties{}
	fp.ReadFile("./input.txt")
	fp.SetVale("year", "2019")
	fp.SetVale("month", "1")
	fp.SetVale("day", "13")
	fp.WriteToFile("newfile.txt")
}
