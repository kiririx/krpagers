package main

import (
	_ "github.com/go-redis/redis/v9"
	_ "github.com/kiririx/amasugi"
	"github.com/kiririx/krpagers/conf"
	"github.com/kiririx/krpagers/router"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 设置端口号启动
	router.SetupRouter(conf.Ginner)
	if err := conf.Ginner.Run(":8080"); err != nil {
		panic(err)
	}
}
