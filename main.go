package main

import (
	"github.com/iEvan-lhr/nihility-dust/anything"
	"github.com/iEvan-lhr/worker/db"
	"github.com/iEvan-lhr/worker/engine"
	"github.com/iEvan-lhr/worker/model"
	"github.com/iEvan-lhr/worker/router"
)

func main() {
	e := engine.Engine{
		W: anything.Wind{},
	}
	e.Start("9080", []any{&db.Conn{}, &db.WorkerInfo{}, &model.User{}}, []any{&router.Router{}})
}
