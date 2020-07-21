package defs

//创建用戶接口请求数据结构
type UserCredential struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

//session信息
type SessionInfo struct {
	UserName string
	TTL      int64
}

//登录成功返回信息
type SigneUp struct {
	Success   bool
	Sessionid string
}
