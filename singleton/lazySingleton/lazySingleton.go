package lazysingleton

import (
	"sync"
	"sync/atomic"
)

//LazySingleton 懒汉式单例模式
type LazySingleton struct {
	c int32
}

//C C
func (ls *LazySingleton) C() int32 {
	return ls.c
}

var instance *LazySingleton
var once sync.Once

func newLazySingleton() *LazySingleton {
	return &LazySingleton{}
}

// GetInstance 获取instance 懒加载
func GetInstance() *LazySingleton {

	// 这种写法不是线程安全的
	//if instance == nil {
	//	instance = newLazySingleton()
	//}

	// golang中可以使用sync.Once保证构造函数只执行一次，并且是懒加载。
	once.Do(func() {
		instance = newLazySingleton()
	})

	return instance
}

//Add 加一
func (ls *LazySingleton) Add() {

	atomic.AddInt32(&ls.c, 1)
}
