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

			fmt.Println("%P", instance)

			w.Done()
		}()
	}

	w.Wait()

}

func BenchmarkLazySingleton_Add(b *testing.B) {

	w := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		w.Add(1)
		go func() {
			instance := GetInstance()

			instance.Add()

			w.Done()
		}()
	}

	w.Wait()

}

func BenchmarkLazySingleton_AddWithMutex(b *testing.B) {

	w := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		w.Add(1)
		go func() {
			instance := GetInstance()

			instance.AddWithMutex()

			w.Done()
		}()
	}

	w.Wait()

}

func BenchmarkLazySingleton_AddWithRWMutex(b *testing.B) {
	w := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		w.Add(1)
		go func() {
			instance := GetInstance()

			instance.AddWithRWMutex()

			w.Done()
		}()
	}

	w.Wait()

}
