package main

import (
	"encoding/json"
	"github.com/MIM1993/stream_video/api/dbops"
	"github.com/MIM1993/stream_video/api/defs"
	"github.com/MIM1993/stream_video/api/session"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

const (
	EmptyString = ""
)

func CreatUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//获取数据
	data, _ := ioutil.ReadAll(r.Body)
	var container = &defs.UserCredential{}
	err := json.Unmarshal(data, container)
	if err != nil {
		ResponseErr(w, defs.ErrorHttpBodyParseFailed)
		return
	}
	//校验数据
	if container.UserName == EmptyString || container.PassWord == EmptyString {
		ResponseErr(w, defs.ErrorParamsErr)
		return
	}

	//处理数据
	err = dbops.AddUserCredential(container.UserName, container.PassWord)
	if err != nil {
		ResponseErr(w, defs.ErrorDBErr)
		return
	}

	//添加session，返回 sessionID
	sessionid := session.GenerateNewSessionID(container.UserName)
	//组织返回struct
	req := &defs.SigneUp{
		Success:   true,
		Sessionid: sessionid,
	}

	//序列化后返回
	reqData, err := json.Marshal(req)
	if err != nil {
		ResponseErr(w, defs.ErrorInternalErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(reqData)
}
