package view_object

import (
	"errors"
	"git.zx-tech.net/app/ait-go-app/common/business_error"
	"git.zx-tech.net/ljhua/zerolog/log"
)

type CommonResp struct {
	Data    interface{} `json:"data"`    // 请求是否成功
	Code    int         `json:"no"`      // 错误码，成功时为0,预期之外错误为-1，业务错误为1
	Message string      `json:"message"` // 附加信息
}

func Resp(data interface{}) CommonResp {

	if err, ok := data.(error); ok {
		log.Err(err)
		var busErr *business_error.BusinessError
		if errors.As(err, &busErr) {
			return CommonResp{
				Data:    struct{}{},
				Code:    1,
				Message: err.Error(),
			}
		}
		return CommonResp{
			Data:    struct{}{},
			Code:    -1,
			Message: err.Error(),
		}
	}
	if data == nil {
		data = struct{}{}
	}
	return CommonResp{
		Data:    data,
		Code:    0,
		Message: "",
	}
}
