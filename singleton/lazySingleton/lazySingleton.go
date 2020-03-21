package lazysingleton

import "sync"

//LazySingleton 懒汉式单例模式
type LazySingleton struct {
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
