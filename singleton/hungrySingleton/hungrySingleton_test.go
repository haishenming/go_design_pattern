package hungrysingleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	goTimes := 10

	wg := sync.WaitGroup{}
	for i := 0; i < goTimes; i++ {
		wg.Add(1)
		go func() {
			instance := GetInstance()
			instance.add()
			fmt.Println(instance.C())
			wg.Done()
		}()

	}

	wg.Wait()

}
