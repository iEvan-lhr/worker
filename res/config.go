package res

import "runtime"

var MasterLen int

func ReSetMasterLen(i int) {
	MasterLen = i
}

func init() {
	MasterLen = runtime.NumCPU() * 2
}
