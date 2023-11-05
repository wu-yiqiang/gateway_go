package initialize

import (
	"gateway_go/utils"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func InitializeValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		_ = v.RegisterValidation("email", utils.ValidateEmail)
		_ = v.RegisterValidation("password", utils.ValidatePassword)
		_ = v.RegisterValidation("mobile", utils.ValidateMobile)
		_ = v.RegisterValidation("valid_service_name", utils.ValidateServiceName)
		_ = v.RegisterValidation("valid_header_transfor", utils.ValidateHeaderTransfor)
		_ = v.RegisterValidation("valid_iplist", utils.ValidateIpList)
		_ = v.RegisterValidation("valid_ipportlist", utils.ValidateIpPortList)
		// 注册自定义 json tag 函数
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
