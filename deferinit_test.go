package deferinit

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDeferInit(t *testing.T) {
	AddInit(func() {
		fmt.Println("1")
	}, nil, 1)
	AddInit(func() {
		fmt.Println("3")
	}, nil, 3)
	AddInit(nil, func() {
		fmt.Println("-3")
	}, -3)
	AddInit(func() {
		fmt.Println("4")
	}, func() {
		fmt.Println("-4")
	}, 4)

	InitAll()
	FiniAll()
}

func TestRoutines(t *testing.T) {
	AddRoutine(func(ch chan struct{}, wg *sync.WaitGroup) {
		fmt.Println("routine 1 start")
		select {
		case <-ch:
		}
		fmt.Println("routine 1 exit")
		wg.Done()
	})

	AddRoutine(func(ch chan struct{}, wg *sync.WaitGroup) {
		fmt.Println("routine 2 start")
		select {
		case <-ch:
		}
		fmt.Println("routine 2 exit")
		wg.Done()
	})

	AddRoutine(func(ch chan struct{}, wg *sync.WaitGroup) {
		fmt.Println("routine 3 start")
		select {
		case <-ch:
		}
		fmt.Println("routine 3 exit")
		wg.Done()
	})

	RunRoutines()
	time.Sleep(1000)
	StopRoutines()
}
