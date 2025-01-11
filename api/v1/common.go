package v1

import (
	"encoding/json"
	"errors"
	"test_mysql/serializer"
)

func ErrorResponse(err error) serializer.Response {
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return serializer.Response{
			Status: 400,
			Data:   "JSON 类型不匹配",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 400,
		Msg:    "请求参数错误",
		Error:  err.Error(),
	}
}
