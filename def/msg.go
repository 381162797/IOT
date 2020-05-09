package def

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "账号密码错误",
	USER_EXIST:                     "用户已存在",
	USER_NOT_EXIST:                 "用户不存在",
	DEVICE_EXIST:                   "设备已存在",
	DEVICE_NOT_EXIST:               "设备不存在",
	TEMPLATE_EXIST:                 "模板已存在",
	TEMPLATE_NOT_EXIST:             "模板不存在",
	TASK_EXIST:                     "任务已存在",
	TASK_NOT_EXIST:                 "任务不存在",
	POINT_EXIST:                    "数据点已存在",
	POINT_NOT_EXIST:                "数据点不存在",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
