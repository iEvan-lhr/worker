package router

import "github.com/iEvan-lhr/nihility-dust/anything"

type Router struct {
}

func (d *Router) TestRouter(mission chan *anything.Mission, data []interface{}) {
	mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []interface{}{"Router Is ok!"}}
}
