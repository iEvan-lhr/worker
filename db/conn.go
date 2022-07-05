package db

import (
	"database/sql"
	"github.com/iEvan-lhr/nihility-dust/anything"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Conn struct {
}

func (c *Conn) GetConn(mission chan *anything.Mission, data []interface{}) {
	dbc, err := sql.Open("mysql", "root:Luhaoran0!@tcp(106.12.170.224:3306)/evan?parseTime=true")
	if err != nil {
		panic(err)
	}
	err = dbc.Ping()
	if err != nil {
		panic(err)
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: dbc,
	}), &gorm.Config{})
	mission <- &anything.Mission{Name: anything.RM, Pursuit: []interface{}{gormDB}}
}
