package common

const (
	// message
	SERVER_ERROR         = "服务器开小差去了，请稍后重试或联系管理员"
	LOGIN_FAIL           = "登录失败"
	NOT_SET_VALIDATOR    = "未设置验证器"
	GET_VALIDATOR_FAIL   = "获取验证器失败"
	NOT_SET_TRANSLATION  = "未设置翻译器"
	GET_TRANSLATION_FAIL = "获取翻译器失败"

	// user
	PHONE_EXIST              = "该手机号已注册"
	PHONE_NOT_EXIST          = "手机号不存在"
	INVALID_PHONE            = "手机号不合法"
	PASSWORS_ERROR           = "登录密码错误"
	NOT_LOGIN                = "用户未登录"
	OLD_PASSWORS_ERROR       = "原密码不正确"
	NEW_PASSWORD_SAME_AS_OLD = "新密码不可与原密码相同"
	UPDATE_PASSWORD_FAILD    = "更新密码失败"
	UPDATE_AVATAR_FAIL       = "更新头像失败"

	SESSION_KEY = "user"

	YES         = "Y"
	NO          = "N"
	TIME_FORMAT = "2006-01-02 15:04:05"

	BASE_SALT = "@#$%^&*0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)
