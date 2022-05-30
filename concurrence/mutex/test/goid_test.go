package test

import (
	"fmt"
	"go_learning/concurrence/mutex"
	"sync"
	"testing"
)

func TestGoId(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		fmt.Println(mutex.GoId())
		wg.Done()
	}()

	go func() {
		fmt.Println(mutex.GoId())
		wg.Done()
	}()

	wg.Wait()
}
