package quark

type Message struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type CodeMap struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	// 通用错误
	StatusOk                     = 200   // 执行成功
	StatusError                  = 10001 // 自定义错误类型
	StatusParamError             = 10002 // 参数错误，比较抽象的状态码
	StatusPhoneFormatError       = 10003 // 手机号格式错误
	StatusEmailFormatError       = 10004 // 邮箱格式错误
	StatusSMSCodeError           = 10005 // 短信验证码错误
	StatusIDCardFormatError      = 10006 // 身份证格式错误
	StatusParamCannotBeEmpty     = 10007 // 参数不可为空
	StatusAccountOrPasswordError = 10008 // 账号或密码错误
	StatusPasswordMismatch       = 10009 // 确认密码不一致
	StatusImageValidationError   = 10010 // 图片大小、尺寸或文件类型不符合要求
	StatusImageUploadFailed      = 10011 // 图片上传失败

	// 接口鉴权相关
	StatusCommonParamMissing       = 10101 // 公共参数不能为空
	StatusRequestTimeExpired       = 10102 // 请求时间已过期
	StatusSignatureError           = 10103 // 签名错误
	StatusAppIDInvalid             = 10104 // appid 无效
	StatusInterfaceValidationError = 10105 // 接口验证服务错误
	StatusRemoteServerError        = 10106 // 远程服务器内部错误
	StatusNoAccessPermission       = 10107 // 无接口访问权限

	// 项目业务类
	StatusUserNotFound            = 20001 // 用户不存在
	StatusNoDataFound             = 20002 // 没查到数据
	StatusFrequentSMSRequests     = 20003 // 手机短信验证码请求过于频繁
	StatusUserDeactivated         = 20004 // 用户已注销
	StatusUserDataSyncError       = 20005 // 用户数据同步异常
	StatusOauthInvalidClient      = 20006 // Oauth2.0 非法客户端
	StatusOauthRedirectURIInvalid = 20007 // Oauth2.0 无效的 redirect_uri
	StatusInternalAPICallError    = 20008 // 内部服务调用异常
	StatusRedisRoutingInfoError   = 20009 // Redis 路由信息异常
	StatusRedisProjectInfoError   = 20010 // Redis 项目信息异常

	// 数据库类
	StatusDBOperationFailed   = 30001 // 数据库操作失败
	StatusDBWriteFailed       = 30002 // 数据库写入失败
	StatusDBUpdateFailed      = 30003 // 数据库更新失败
	StatusDBDeleteFailed      = 30004 // 数据库删除失败
	StatusDBQueryFailed       = 30005 // 数据库查询失败
	StatusRedisOperationError = 30006 // Redis 操作异常

	// 会话类
	StatusSessionExpired      = 40001 // 会话过期（登录状态过期）
	StatusInvalidAuthCode     = 40002 // 无效的授权码（Oauth2.0 中的 code）或临时凭证
	StatusInvalidAccessToken  = 40003 // 无效的访问令牌（Oauth2.0 中的 token）或最终凭证
	StatusInvalidEmailCode    = 40004 // 无效的邮箱验证码
	StatusMismatchedClient    = 40005 // 不匹配的客户端
	StatusInvalidRefreshToken = 40006 // Oauth2.0 无效的 refresh token
	StatusTokenFormatError    = 40007 // 错误的访问令牌（格式错误）
	StatusInvalidScope        = 40008 // Oauth2.0 无效的 scope，授权范围失效
	StatusHTTPAuthFailed      = 40009 // HTTP Basic Auth 失败

	// 权限类
	StatusNotLoggedIn        = 50001 // 未登录（SSO 统一授权）
	StatusUserAccountError   = 50002 // 用户账号信息异常
	StatusPhoneNotVerified   = 50003 // 手机号码未认证
	StatusPermissionExceeded = 50004 // 超出普通用户权限
	StatusInvisibleUser      = 50005 // 不可见用户非法登陆

	// 远程调用类
	StatusRemoteCallFailed    = 70001 // 远程调用失败（如易盾、邮件发送等）
	StatusBaiduAPICallError   = 70002 // 调用百度接口异常
	StatusTencentAPICallError = 70003 // 调用腾讯接口异常

	// 加解密类
	StatusInvalidSignature           = 80001 // 无效的签名
	StatusSignatureVerificationError = 80002 // 验证签名时发生错误
	StatusEncryptionFailed           = 80003 // 加密失败
	StatusDecryptionFailed           = 80004 // 解密失败
	StatusInvalidEncryptionKey       = 80005 // 无效的加密 key（对称或非对称）

	// 其他类
	StatusProgramError         = 90001 // 程序错误
	StatusUUIDGenerationFailed = 90002 // UUID 生成失败
)

var CodeMaps = []*CodeMap{
	// 通用错误
	{
		Code: StatusOk,
		Msg:  "ok",
	},
	{
		Code: StatusError,
		Msg:  "Internal Server Error",
	},
	{
		Code: StatusParamError,
		Msg:  "参数错误",
	},
	{
		Code: StatusPhoneFormatError,
		Msg:  "手机号格式错误",
	},
	{
		Code: StatusEmailFormatError,
		Msg:  "邮箱格式错误",
	},
	{
		Code: StatusSMSCodeError,
		Msg:  "短信验证码错误",
	},
	{
		Code: StatusIDCardFormatError,
		Msg:  "身份证格式错误",
	},
	{
		Code: StatusParamCannotBeEmpty,
		Msg:  "参数不可为空",
	},
	{
		Code: StatusAccountOrPasswordError,
		Msg:  "账号或密码错误",
	},
	{
		Code: StatusPasswordMismatch,
		Msg:  "确认密码不一致",
	},
	{
		Code: StatusImageValidationError,
		Msg:  "图片大小、尺寸或文件类型不符合要求",
	},
	{
		Code: StatusImageUploadFailed,
		Msg:  "图片上传失败",
	},

	// 接口鉴权相关
	{
		Code: StatusCommonParamMissing,
		Msg:  "公共参数不能为空",
	},
	{
		Code: StatusRequestTimeExpired,
		Msg:  "请求时间已过期",
	},
	{
		Code: StatusSignatureError,
		Msg:  "签名错误",
	},
	{
		Code: StatusAppIDInvalid,
		Msg:  "appid 无效",
	},
	{
		Code: StatusInterfaceValidationError,
		Msg:  "接口验证服务错误",
	},
	{
		Code: StatusRemoteServerError,
		Msg:  "远程服务器内部错误",
	},
	{
		Code: StatusNoAccessPermission,
		Msg:  "无接口访问权限",
	},

	// 项目业务类
	{
		Code: StatusUserNotFound,
		Msg:  "用户不存在",
	},
	{
		Code: StatusNoDataFound,
		Msg:  "没有找到数据",
	},
	{
		Code: StatusFrequentSMSRequests,
		Msg:  "手机短信验证码请求过于频繁",
	},
	{
		Code: StatusUserDeactivated,
		Msg:  "用户已注销",
	},
	{
		Code: StatusUserDataSyncError,
		Msg:  "用户数据同步异常",
	},
	{
		Code: StatusOauthInvalidClient,
		Msg:  "Oauth2.0 非法客户端",
	},
	{
		Code: StatusOauthRedirectURIInvalid,
		Msg:  "Oauth2.0 无效的 redirect_uri",
	},
	{
		Code: StatusInternalAPICallError,
		Msg:  "内部服务调用异常",
	},
	{
		Code: StatusRedisRoutingInfoError,
		Msg:  "Redis 路由信息异常",
	},
	{
		Code: StatusRedisProjectInfoError,
		Msg:  "Redis 项目信息异常",
	},

	// 数据库类
	{
		Code: StatusDBOperationFailed,
		Msg:  "数据库操作失败",
	},
	{
		Code: StatusDBWriteFailed,
		Msg:  "数据库写入失败",
	},
	{
		Code: StatusDBUpdateFailed,
		Msg:  "数据库更新失败",
	},
	{
		Code: StatusDBDeleteFailed,
		Msg:  "数据库删除失败",
	},
	{
		Code: StatusDBQueryFailed,
		Msg:  "数据库查询失败",
	},
	{
		Code: StatusRedisOperationError,
		Msg:  "Redis 操作异常",
	},

	// 会话类
	{
		Code: StatusSessionExpired,
		Msg:  "会话过期",
	},
	{
		Code: StatusInvalidAuthCode,
		Msg:  "无效的授权码或临时凭证",
	},
	{
		Code: StatusInvalidAccessToken,
		Msg:  "无效的访问令牌或最终凭证",
	},
	{
		Code: StatusInvalidEmailCode,
		Msg:  "无效的邮箱验证码",
	},
	{
		Code: StatusMismatchedClient,
		Msg:  "不匹配的客户端",
	},
	{
		Code: StatusInvalidRefreshToken,
		Msg:  "Oauth2.0 无效的 refresh token",
	},
	{
		Code: StatusTokenFormatError,
		Msg:  "错误的访问令牌格式",
	},
	{
		Code: StatusInvalidScope,
		Msg:  "Oauth2.0 无效的 scope，授权范围失效",
	},
	{
		Code: StatusHTTPAuthFailed,
		Msg:  "HTTP Basic Auth 失败",
	},

	// 权限类
	{
		Code: StatusNotLoggedIn,
		Msg:  "未登录",
	},
	{
		Code: StatusUserAccountError,
		Msg:  "用户账号信息异常",
	},
	{
		Code: StatusPhoneNotVerified,
		Msg:  "手机号码未认证",
	},
	{
		Code: StatusPermissionExceeded,
		Msg:  "超出普通用户权限",
	},
	{
		Code: StatusInvisibleUser,
		Msg:  "不可见用户非法登陆",
	},

	// 远程调用类
	{
		Code: StatusRemoteCallFailed,
		Msg:  "远程调用失败",
	},
	{
		Code: StatusBaiduAPICallError,
		Msg:  "调用百度接口异常",
	},
	{
		Code: StatusTencentAPICallError,
		Msg:  "调用腾讯接口异常",
	},

	// 加解密类
	{
		Code: StatusInvalidSignature,
		Msg:  "无效的签名",
	},
	{
		Code: StatusSignatureVerificationError,
		Msg:  "验证签名时发生错误",
	},
	{
		Code: StatusEncryptionFailed,
		Msg:  "加密失败",
	},
	{
		Code: StatusDecryptionFailed,
		Msg:  "解密失败",
	},
	{
		Code: StatusInvalidEncryptionKey,
		Msg:  "无效的加密 key",
	},

	// 其他类
	{
		Code: StatusProgramError,
		Msg:  "程序错误",
	},
	{
		Code: StatusUUIDGenerationFailed,
		Msg:  "UUID 生成失败",
	},
}

// 根据code获取错误信息
func GetMsgByCode(code int) string {
	for _, v := range CodeMaps {
		if v.Code == code {
			return v.Msg
		}
	}
	return ""
}

// 返回错误信息，Error("内部服务调用异常") | Error("错误", map[string]interface{}{"title":"标题"}) | Error(10001, "错误", map[string]interface{}{"title":"标题"})
func Error(params ...interface{}) *Message {
	var (
		code = StatusError
		msg  = ""
		data interface{}
	)
	if len(params) == 1 {
		msg = params[0].(string)
	}
	if len(params) == 2 {
		msg = params[0].(string)
		data = params[1]
	}
	if len(params) == 3 {
		code = params[0].(int)
		msg = params[1].(string)
		if msg == "" {
			msg = GetMsgByCode(StatusError)
		}
		data = params[2]
	}

	return &Message{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// 返回正确信息
func Success(message string, data interface{}) *Message {
	return &Message{
		Code: StatusOk,
		Msg:  message,
		Data: data,
	}
}
