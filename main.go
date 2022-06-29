package main

import (
	"github.com/iEvan-lhr/nihility-dust/wind"
	"log"
	"time"
	"worker/engine"
	"worker/router"
)

func main() {
	e := engine.Engine{
		W: wind.Wind{},
	}
	e.W.Register(&router.Router{})
	e.Init()
	e.RegisterRouter()
	log.Println("初始化版本:", time.Now().Format("2006-01-02 15:04:05"))
	e.Run(":9098")
}
