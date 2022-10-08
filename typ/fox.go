package typ

import (
	"github.com/iEvan-lhr/worker/res"
	"sync"
	"time"
)

type FoxExecutor struct {
	Master []chan struct{}
	DoMap  *sync.Map
	Counts []int
	i      int
}

func (f *FoxExecutor) Init() {
	for i := 0; i < res.MasterLen; i++ {
		f.Master = append(f.Master, make(chan struct{}))
	}
	f.Counts = make([]int, 16)
	f.DoMap = &sync.Map{}
}

func (f *FoxExecutor) DoMaps() chan struct{} {
	if _, ok := f.DoMap.Load(f.i); ok {
		return f.checkNilMissionChan()
	} else {
		f.DoMap.Store(f.i, 0)
		f.Counts[f.i]++
		go func(index int) {
			<-f.Master[index]
			f.DoMap.Delete(index)
		}(f.i)
		if f.i != res.MasterLen-1 {
			f.i++
		} else {
			f.i = 0
		}
	}
	return f.Master[f.i]
}

func (f *FoxExecutor) checkNilMissionChan() chan struct{} {
NEXT:
	for i := 0; i < res.MasterLen; i++ {
		if _, ok := f.DoMap.Load(i); !ok {
			f.DoMap.Store(i, 0)
			f.i = i
			f.Counts[f.i]++
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
