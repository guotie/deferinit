package deferinit

import (
	"sync"
)

type fn struct {
	fi  func() // init function
	ff  func() // fini functions
	pos int
}

type gr func(chan struct{})

var (
	fns       = make([]fn, 0)
	routines  = make([]gr, 0)
	exitChans = make([]chan struct{}, 0)
	lock      sync.Mutex
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

func AddInit(fi func(), ff func(), pos int) {
	lock.Lock()
	defer lock.Unlock()

	s := fn{fi, ff, pos}

	var (
		f     fn
		index int
	)
	for index, f = range fns {
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
		go r(exitChans[i])
	}
}

func StopRoutines() {
	for _, ch := range exitChans {
		ch <- struct{}{}
		_ = <-ch
	}
}
