package res

import "runtime"

var MasterLen int

func ReSetMasterLen(i int) {
	MasterLen = i
}

func init() {
	if runtime.NumCPU() < 8 {
		MasterLen = 16
	} else {
		MasterLen = runtime.NumCPU() * 2
	}
}
