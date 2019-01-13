package model

import (
	"fmt"
	"practice/model/properties"

	p "github.com/magiconair/properties"
)

type FileIO interface {
	ReadFile(fileName string)
	WriteToFile(fileName string)
	SetVale(key, value string)
	GetValue(key string)
}

type FileProperties struct {
	tmp *p.Properties
}

func (b *FileProperties) ReadFile(fileName string) {
	var err error
	b.tmp, err = p.LoadFile(fileName, p.UTF8)
	if err != nil {
		fmt.Println("readfile error:", err)
	}
}

func (b *FileProperties) WriteToFile(fileName string) {
	properties.StoreToFile(b.tmp, "written by FileProperties")
}

func (b *FileProperties) SetVale(key, value string) {
	b.tmp.SetValue(key, value)
}

// GetValue - 範例沒用到
func (b *FileProperties) GetValue(key string) {}
