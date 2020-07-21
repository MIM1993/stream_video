package session

import (
	"github.com/MIM1993/stream_video/api/dbops"
	"github.com/MIM1993/stream_video/api/defs"
	"github.com/MIM1993/stream_video/api/utils"
	"sync"
	"time"
)

//创建本地缓存
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

//计算时间
func nowInMilli() int64 {
	return time.Now().UnixNano() / 1e6
}

//创建sessionID
func GenerateNewSessionID(userName string) string {
	uuid := utils.NewUUID(userName)
	//session过期时间
	ttl := nowInMilli() + 30*60*1000

	ss := &defs.SessionInfo{
		UserName: userName,
		TTL:      ttl,
	}

	//放入缓存
	sessionMap.Store(uuid, ss)
	//存入数据库
	dbops.InsertSession(userName, ttl, uuid)

	return uuid
}
