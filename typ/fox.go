package typ

import (
	"github.com/iEvan-lhr/worker/res"
	"sync"
	"time"
)

type FoxExecutor struct {
	Master []chan struct{}
	DoMap  *sync.Map
	i      int
}

func (f *FoxExecutor) Init() {
	for i := 0; i < res.MasterLen; i++ {
		f.Master = append(f.Master, make(chan struct{}))
	}
	f.DoMap = &sync.Map{}
}

func (f *FoxExecutor) InitByUser(num int) {
	res.MasterLen = num
	for i := 0; i < res.MasterLen; i++ {
		f.Master = append(f.Master, make(chan struct{}))
	}
	f.DoMap = &sync.Map{}
}

func (f *FoxExecutor) DoMaps() chan struct{} {
	return f.doMap()
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
	time.Sleep(10 * time.Millisecond)
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
