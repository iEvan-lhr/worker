package db

import (
	"github.com/iEvan-lhr/nihility-dust/anything"
	"github.com/iEvan-lhr/worker/typ"
	"gorm.io/gorm"
)

type WorkerInfo struct {
}

func (w WorkerInfo) GetWorkerInfoList(mission chan *anything.Mission, data []any) {
	temp := <-anything.DoChanN("GetConn", nil)
	db := temp.Pursuit[0].(*gorm.DB)
	work := typ.WorkerInfo{}
	db.Find(&work, "id = ?", 1)
	mission <- &anything.Mission{Name: anything.IM, Pursuit: []any{work}}
}
