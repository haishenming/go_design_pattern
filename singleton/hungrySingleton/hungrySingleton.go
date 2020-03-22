package hungrysingleton

import "sync/atomic"

//HungrySingleton 饿汉式，在一开始就实例化
type HungrySingleton struct {
	c int32
}

//C C
func (hs *HungrySingleton) C() int32 {
	return hs.c
}

var instance *HungrySingleton

func newHungrySingleton() *HungrySingleton {
	return &HungrySingleton{}
}

func init() {
	instance = newHungrySingleton()
}

func (hs *HungrySingleton) add() {
	atomic.AddInt32(&hs.c, 1)
}

//GetInstance 获取单例
func GetInstance() *HungrySingleton {
	return instance
}
