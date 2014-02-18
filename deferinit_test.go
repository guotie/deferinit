package deferinit

import (
	"fmt"
	"testing"
)

func TestDeferInit(t *testing.T) {
	AddInit(func() {
		fmt.Println("1")
	}, nil)
	AddInit(func() {
		fmt.Println("2")
	}, nil)
	AddInit(func() {
		fmt.Println("3")
	}, func() {
		fmt.Println("-3")
	})

	InitAll()
	FiniAll()
}
