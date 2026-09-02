package response

import "context"

// SuccessResponse 成功响应
type SuccessResponse struct {
	Code int    `json:"code"`           // 状态码
	Msg  string `json:"msg"`            // 提示信息
	Data any    `json:"data,omitempty"` // 响应数据
}

// Ok 创建成功响应
func Ok(_ context.Context, data any) any {
	return &SuccessResponse{
		Code: 0,
		Msg:  "ok",
		Data: data,
	}
}
