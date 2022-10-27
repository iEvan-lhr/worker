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
	W      anything.Wind
	Origin []string
}

func (e *Engine) Init() {
	e.W.Init()
	f := &typ.FoxExecutor{}
	f.Init()
	anything.SetController(f)
}

func (e *Engine) RegisterRouter() {
	e.W.R.Range(func(key, value interface{}) bool {
		func(name string) {
			http.HandleFunc("/"+tool.EString(name).FirstLowerBackString(), func(writer http.ResponseWriter, request *http.Request) {
				switch len(e.Origin) {
				case 1:
					writer.Header().Set("Access-Control-Allow-Origin", e.Origin[0])
				case 2:
					writer.Header().Set("Access-Control-Allow-Origin", e.Origin[0])
					writer.Header().Set("Access-Control-Allow-Methods", e.Origin[1])
				case 3:
					writer.Header().Set("Access-Control-Allow-Origin", e.Origin[0])
					writer.Header().Set("Access-Control-Allow-Methods", e.Origin[1])
					writer.Header().Set("Access-Control-Allow-Headers", e.Origin[2])
				}
				if len(e.Origin) > 0 {
					writer.Header().Set("Access-Control-Allow-Origin", e.Origin[0])
				}
				key1 := e.W.Schedule(name, []interface{}{writer, request})
				// 出口
				<-e.W.E[key1]
				mission, _ := e.W.A.Load(key1)
				_, _ = fmt.Fprintf(writer, "%s", (mission.([]interface{})[0]).(*anything.Mission).Pursuit)
				delete(e.W.E, key1)
			})
		}(key.(string))
		return true
	})
}

func (e *Engine) Run(addr string) {
	_ = http.ListenAndServe(addr, nil)
}

func (e *Engine) Start(port string, model, routers []interface{}) {
	e.W.Register(model...)
	e.W.Register(routers...)
	e.W.RegisterRouters(routers)
	e.Init()
	e.RegisterRouter()
	log.Println("初始化版本:", time.Now().Format("2006-01-02 15:04:05"))
	e.Run(":" + port)
}
