package lazysingleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	goTimes := 10

	w := sync.WaitGroup{}
	for i := 0; i < goTimes; i++ {
		w.Add(1)
		go func() {
			instance := GetInstance()

			instance.Add()

			fmt.Println(instance.C())

			w.Done()
		}()
	}

	w.Wait()

}
