package router

import (
	"fmt"
	"github.com/iEvan-lhr/nihility-dust/anything"
	"gorm.io/gorm"
)

type Router struct {
}

func (d *Router) TestRouter(mission chan *anything.Mission, data []interface{}) {
	temp := <-anything.DoChanTemp(mission, []interface{}{"GetConn", "NULL"})
	db := temp.Pursuit[0].(*gorm.DB)
	fmt.Println(db.Name())
	mission <- &anything.Mission{Name: anything.ExitFunction, Pursuit: []interface{}{db.Name()}}
}
