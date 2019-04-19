package errors

import (
	"github.com/johnnyeven/libtools/courier/status_error"
)

func init() {
	status_error.StatusErrorCodes.Register("BadRequest", 400200000, "请求参数错误", "", false)
	status_error.StatusErrorCodes.Register("BadAuthToken", 400200001, "Token格式错误，请使用『类型:Token』", "", true)
	status_error.StatusErrorCodes.Register("BadAuthChannal", 400200002, "不支持的Token格式", "", true)
	status_error.StatusErrorCodes.Register("Unauthorized", 401200000, "未授权", "", true)
	status_error.StatusErrorCodes.Register("Forbidden", 403200000, "不允许操作", "", true)
	status_error.StatusErrorCodes.Register("NotFound", 404200000, "未找到", "", false)
	status_error.StatusErrorCodes.Register("Conflict", 409200000, "操作冲突", "", true)
	status_error.StatusErrorCodes.Register("InternalError", 500200000, "内部处理错误", "", false)
	status_error.StatusErrorCodes.Register("UpstreamError", 500200001, "后台服务错误，请稍后再试", "", true)
}
