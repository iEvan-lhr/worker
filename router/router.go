package router

import (
	"github.com/iEvan-lhr/nihility-dust/anything"
)

type Router struct {
}

func (r *Router) TestRouter(mission chan *anything.Mission, data []any) {
	temp := <-anything.DoChanN("GetWorkerInfoList", nil)
	mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []any{temp.Pursuit}}
}

func (r *Router) Index(mission chan *anything.Mission, data []any) {
	mission <- &anything.Mission{Name: "HtmlResource", Pursuit: []any{"index"}}
}

func (r *Router) GetUser(mission chan *anything.Mission, data []any) {
	name := ""
	switch len(data) {
	case 1:
		name = "AuthLogin"
	case 2:
		name = "PasswordLogin"
	case 3:
		name = "AdminLogin"
	default:
		mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []any{"ERROR TYPE is not supported"}}
	}
	mission <- &anything.Mission{Name: name, Pursuit: data}
}
