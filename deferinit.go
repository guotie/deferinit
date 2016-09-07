package deferinit

import (
	"fmt"
	"sync"
)

type fn struct {
	fi  func() // init function
	ff  func() // fini functions
	pos int
}

type gr func(chan struct{}, *sync.WaitGroup)

var (
	_         = fmt.Printf
	fns       = make([]fn, 0)
	routines  = make([]gr, 0)
	exitChans = make([]chan struct{}, 0)
	lock      sync.Mutex
	wg        sync.WaitGroup
)

func InitAll() {
	lock.Lock()
	defer lock.Unlock()

	for _, f := range fns {
		if f.fi != nil {
			f.fi()
		}
	}
}

func FiniAll() {
	lock.Lock()
	defer lock.Unlock()

	for i := len(fns) - 1; i >= 0; i-- {
		if fns[i].ff != nil {
			fns[i].ff()
		}
	}
}

// pos越大, 优先级越高
func AddInit(fi func(), ff func(), pos int) {
	var (
		f     fn
		index int
	)

	lock.Lock()
	defer lock.Unlock()

	s := fn{fi, ff, pos}

	for index = 0; index < len(fns); index++ {
		f = fns[index]
		if f.pos < pos {
			break
		}
	}

	fns = append(fns[0:index], append([]fn{s}, fns[index:]...)...)
}

// routines
func AddRoutine(f gr) {
	routines = append(routines, f)
}

func RunRoutines() {
	exitChans = make([]chan struct{}, len(routines))
	for i, r := range routines {
		exitChans[i] = make(chan struct{})
		wg.Add(1)
		go r(exitChans[i], &wg)
	}
}

func StopRoutines() {
	for _, ch := range exitChans {
		ch <- struct{}{}
	}
	wg.Wait()
}
