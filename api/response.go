package main

import (
	"encoding/json"
	"net/http"
	"github.com/MIM1993/stream_video/api/defs"
)

//返回错误函数
func ResponseErr(w http.ResponseWriter, err defs.ErrResponse) {
	w.WriteHeader(err.HttpCode)
	errData, _ := json.Marshal(err.Error)
	w.Write(errData)
}

//返回普通错误
func ResponseNormalErr(w http.ResponseWriter, httpCode int, repMsg []byte) {
	w.WriteHeader(httpCode)
	w.Write(repMsg)
}
