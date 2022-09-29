package engine

import (
	"fmt"
	"github.com/iEvan-lhr/nihility-dust/anything"
	tool "github.com/iEvan-lhr/string"
	"github.com/iEvan-lhr/worker/typ"
	"log"
	"net/http"
	"time"
)

type Engine struct {
	W anything.Wind
}

func (e *Engine) Init() {
	e.W.Init()
	f := &typ.FoxExecutor{}
	f.Init()
	e.W.SetController(f)
}

func (e *Engine) RegisterRouter() {
	e.W.R.Range(func(key, value any) bool {
		func(name string) {
			http.HandleFunc("/"+tool.EString(name).FirstLowerBackString(), func(writer http.ResponseWriter, request *http.Request) {
				key := e.W.Schedule(name, []any{writer, request})
				// 出口
				<-e.W.E[key]
				mission, _ := e.W.A.Load(key)
				_, _ = fmt.Fprintf(writer, "%s", mission.([]any)[0])
				delete(e.W.E, key)
			})
		}(key.(string))
		return true
	})
}

func (e *Engine) Run(addr string) {
	_ = http.ListenAndServe(addr, nil)
}

func (e *Engine) Start(port string, model, routers []any) {
	e.W.Register(model...)
	e.W.Register(routers...)
	e.W.RegisterRouters(routers)
	e.Init()
	e.RegisterRouter()
	log.Println("初始化版本:", time.Now().Format("2006-01-02 15:04:05"))
	e.Run(":" + port)
}
