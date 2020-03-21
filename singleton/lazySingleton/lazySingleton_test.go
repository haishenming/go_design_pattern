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

			fmt.Println("%P", instance)

			w.Done()
		}()
	}

	w.Wait()

}
