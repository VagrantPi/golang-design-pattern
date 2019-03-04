package model

import (
	"fmt"
	"sync"
)

type singletonOP interface {
	Count()
	Get() int
}

type singleton struct {
	count int
}

func (s *singleton) Count() {
	s.count++
}

func (s *singleton) Get() int {
	return s.count
}

func (s *singleton) singleton() {
	fmt.Println("create a instance.")
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
		instance.singleton()
	})
	return instance
}
