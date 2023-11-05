package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// ValidateMobile 校验手机号
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	if mobile == "" {
		return true
	}
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}

// ValidateEmail 校验邮箱
func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	if email == "" {
		return true
	}
	ok, _ := regexp.MatchString(`^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`, email)
	if !ok {
		return false
	}
	return true
}

// ValidatePassword 密码强度校验
func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if password == "" {
		return true
	}
	ok, _ := regexp.MatchString(`^[\w_-]{8,16}$`, password)
	if !ok {
		return false
	}
	return true
}

// ValidateServiceName 服务名校验
func ValidateServiceName(fl validator.FieldLevel) bool {
	return false
}

// ValidateHeaderTransfor 头字段校验
func ValidateHeaderTransfor(fl validator.FieldLevel) bool {
	return false
}

// ValidateIpList IP列表校验
func ValidateIpList(fl validator.FieldLevel) bool {
	return false
}

// ValidateIpPortList IP端口列表校验
func ValidateIpPortList(fl validator.FieldLevel) bool {

	return false
}
