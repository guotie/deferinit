package deferinit

import (
	"fmt"
	"testing"
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
