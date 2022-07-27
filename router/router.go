package router

import (
	"github.com/iEvan-lhr/nihility-dust/anything"
)

type Router struct {
}

func (r *Router) TestRouter(mission chan *anything.Mission, data []any) {
	temp := <-anything.DoChanN("GetWorkerInfoList", nil)
	//log.Println(temp)
	mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []any{temp.Pursuit}}
}

func (r *Router) Index(mission chan *anything.Mission, data []any) {
	mission <- &anything.Mission{Name: "HtmlResource", Pursuit: []any{"index"}}
}
