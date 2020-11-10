package main

import (
	"fmt"
	"sync"
	"unsafe"
)

//单例设计模式
type Singleton struct{}
type ESingleton struct{}

//懒汉模式，用的时候才初始化
var singleton *Singleton

//饿汉模式，一开始就初始化
var eSingleton *ESingleton = new(ESingleton)

// 确保执行一次
var once sync.Once

// 懒汉模式
func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("once.Do")
		singleton = &Singleton{}
	})
	return singleton
}

func GetInstanceE() *ESingleton {
	return eSingleton
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			obj := GetInstance()
			fmt.Printf("add1:%x \n", unsafe.Pointer(obj))
			fmt.Printf("add2:%p \n", obj)
			wg.Done()
		}()
	}
	wg.Wait()
}
