package engine

import (
	"fmt"
	"github.com/iEvan-lhr/nihility-dust/wind"
	tool "github.com/iEvan-lhr/string"
	"net/http"
)

type Engine struct {
	W wind.Wind
}

func (e *Engine) Init() {
	e.W.Init()
}

func (e *Engine) RegisterRouter() {
	for s := range e.W.M {
		func(name string) {
			http.HandleFunc("/"+tool.EString(name).FirstLowerBackString(), func(writer http.ResponseWriter, request *http.Request) {
				key := e.W.Schedule(name, writer, request)
				// 出口
				<-e.W.E[key]
				mission, _ := e.W.A.Load(key)
				_, _ = fmt.Fprintf(writer, "%s", mission.([]interface{})[0])
				delete(e.W.E, key)
			})
		}(s)
	}
}

func (e *Engine) Run(addr string) {
	_ = http.ListenAndServe(addr, nil)
}
