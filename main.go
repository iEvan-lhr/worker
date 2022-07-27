package main

import (
	"github.com/iEvan-lhr/nihility-dust/anything"
	"worker/db"
	"worker/engine"
	"worker/router"
)

func main() {
	e := engine.Engine{
		W: anything.Wind{},
	}
	e.Start("9080", []any{&db.Conn{}, &db.WorkerInfo{}}, []any{&router.Router{}})
}
