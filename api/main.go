package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreatUser)

	return router
}

func main() {
	//注册路由
	r := RegisterHandlers()

	//监听本地8889端口
	err := http.ListenAndServe("127.0.0.1:8889", r)
	if err != nil {
		log.Printf("ListenAndServe err :%v\n", err)
	}
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
}
