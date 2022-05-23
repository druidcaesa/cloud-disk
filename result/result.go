package result

import "cloud-disk/internal/types"

// OK 成功返回结果
func OK(msg string, data ...interface{}) types.Result {
	resp := types.Result{
		Code: 200,
	}
	if msg == "" {
		resp.Message = "操作成功"
	}
	resp.Data = data[0]
	return resp
}

// ERROR 失败返回结果
func ERROR(msg string) types.Result {
	return types.Result{
		Code:    500,
		Message: msg,
	}
}
