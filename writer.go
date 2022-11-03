package xhttp

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data == nil {
		return
	}
	res, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Ok(w http.ResponseWriter, code int, msg string, data interface{}) {
	JSON(w, http.StatusOK, BaseResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func BadRequest(w http.ResponseWriter, code int, msg string, data interface{}) {
	JSON(w, http.StatusBadRequest, BaseResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Unauthorized(w http.ResponseWriter, code int, msg string, data interface{}) {
	JSON(w, http.StatusUnauthorized, BaseResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
