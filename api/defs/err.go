package defs

import "net/http"

type Err struct {
	Err_code int    `json:"err_code"`
	Err_msg  string `json:"err_msg"`
}

type ErrResponse struct {
	HttpCode int
	Error    Err
}

var (
	ErrorHttpBodyParseFailed = ErrResponse{
		HttpCode: http.StatusBadRequest,
		Error:    Err{Err_code: 001, Err_msg: "requset body is not correct"},
	}

	ErrorNotAuthUser = ErrResponse{
		HttpCode: http.StatusUnauthorized,
		Error:    Err{Err_code: 002, Err_msg: "User authentication failed"},
	}

	ErrorDBErr = ErrResponse{
		HttpCode: http.StatusInternalServerError,
		Error:    Err{Err_code: 003, Err_msg: "DB operation err"},
	}

	ErrorInternalErr = ErrResponse{
		HttpCode: http.StatusInternalServerError,
		Error:    Err{Err_code: 004, Err_msg: "Internal server err"},
	}

	ErrorParamsErr = ErrResponse{
		HttpCode: http.StatusBadRequest,
		Error:    Err{Err_code: 005, Err_msg: "Request params err"},
	}
)
