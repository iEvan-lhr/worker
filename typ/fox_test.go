package typ

import (
	"github.com/iEvan-lhr/worker/res"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	res.MasterLen = 16
	f := &FoxExecutor{}
	sum := 0
	f.Init()
	for i := 0; i < 16; i++ {
		maps := f.DoMaps()
		go func(m chan struct{}, index int) {
			time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
			m <- struct{}{}
		}(maps, i)
		if i == 15 {
			i = 0
		}
		sum++
		if sum == 900 {
			break
		}
	}
	time.Sleep(2 * time.Second)
	log.Println(f.Counts)
	count := 0
	for _, v := range f.Counts {
		count += v
	}
	log.Println("All count", count)
}
