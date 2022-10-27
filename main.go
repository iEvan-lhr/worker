package main

import (
	"github.com/iEvan-lhr/nihility-dust/anything"
	"github.com/iEvan-lhr/worker/engine"
)

func main() {
	e := engine.Engine{
		W: anything.Wind{},
	}
	e.Start("9080", []interface{}{}, []interface{}{})
}
