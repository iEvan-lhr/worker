package typ

import (
	"github.com/iEvan-lhr/worker/res"
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
		go func(index int) {
			m := f.DoMaps()
			time.Sleep(time.Duration(rand.Intn(2)) * time.Millisecond)
			m <- struct{}{}
		}(i)
		if i == 15 {
			i = 0
		}
		sum++
		if sum == 9000 {
			break
		}
	}
	time.Sleep(2 * time.Second)
}
