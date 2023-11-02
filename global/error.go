package global

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrors struct {
	BusinessError CustomError
	ValidateError CustomError
	TokenError    CustomError
	ServiceError  CustomError
	GatewayError  CustomError
}

var Errors = CustomErrors{
	BusinessError: CustomError{400, "客户端业务错误"},
	ValidateError: CustomError{422, "请求参数错误"},
	TokenError:    CustomError{401, "登录授权失效"},
	ServiceError:  CustomError{500, "服务出错"},
	GatewayError:  CustomError{50200, "网关错误"},
}
