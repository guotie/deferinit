package deferinit

import (
	"sync"
)

var (
	fns  []func() = make([]func(), 0)
	fis  []func() = make([]func(), 0)
	lock sync.Mutex
)

func InitAll() {
	lock.Lock()
	defer lock.Unlock()

	for _, f := range fns {
		f()
	}
}

func FiniAll() {
	lock.Lock()
	defer lock.Unlock()

	for i := len(fis) - 1; i >= 0; i-- {
		fis[i]()
	}
}

func AddInit(f func(), fi func()) {
	lock.Lock()
	defer lock.Unlock()

	if f != nil {
		fns = append(fns, f)
	}
	if fi != nil {
		fis = append(fis, fi)
	}
}
