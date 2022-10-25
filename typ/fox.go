package typ

import (
	"github.com/iEvan-lhr/worker/res"
	"sync"
	"time"
)

type FoxExecutor struct {
	Master []chan struct{}
	DoMap  *sync.Map
	Inner  chan struct{}
	Outer  chan chan struct{}
	i      int
}

func (f *FoxExecutor) Init() {
	for i := 0; i < res.MasterLen; i++ {
		f.Master = append(f.Master, make(chan struct{}))
	}
	f.DoMap = &sync.Map{}
	f.Inner = make(chan struct{})
	f.Outer = make(chan chan struct{})
	go func() {
		for {
			<-f.Inner
			f.Outer <- f.doMap()
		}
	}()
}

func (f *FoxExecutor) DoMaps() chan struct{} {
	f.Inner <- struct{}{}
	return <-f.Outer
}

func (f *FoxExecutor) doMap() chan struct{} {
	if _, ok := f.DoMap.Load(f.i); ok {
		return f.checkNilMissionChan()
	} else {
		f.DoMap.Store(f.i, 0)
		go func(index int) {
			<-f.Master[index]
			f.DoMap.Delete(index)
		}(f.i)
		inx := f.i
		if f.i != res.MasterLen-1 {
			f.i++
		} else {
			f.i = 0
		}
		return f.Master[inx]
	}
}

func (f *FoxExecutor) checkNilMissionChan() chan struct{} {
NEXT:
	for i := 0; i < res.MasterLen; i++ {
		if _, ok := f.DoMap.Load(i); !ok {
			f.DoMap.Store(i, 0)
			f.i = i
			go func(index int) {
				<-f.Master[index]
				f.DoMap.Delete(index)
			}(i)
			return f.Master[i]
		}
	}
	time.Sleep(1 * time.Millisecond)
	goto NEXT
}

//func (f *FoxExecutor) CheckUseChan() {
//	sum := 0
//	for i := 0; i < res.MasterLen; i++ {
//		if _, ok := f.DoMap.Load(i); !ok {
//			sum++
//		}
//	}
//}
