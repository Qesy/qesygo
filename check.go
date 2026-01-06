package qesygo

import (
	"regexp"
	"unicode"
)

func CheckMobile(phone string) bool { // CheckMobile 检验手机号 (第一位必为1的十一位数字)
	regRuler := "^1\\d{10}$"
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(phone)
}

func CheckIdCard(card string) bool { // CheckIdCard 检验身份证
	//18位身份证 ^(\d{17})([0-9]|X)$
	// 匹配规则
	// (^\d{15}$) 15位身份证
	// (^\d{18}$) 18位身份证
	// (^\d{17}(\d|X|x)$) 18位身份证 最后一位为X的用户
	regRuler := "(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)"
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(card)
}

// CheckPasswordComplexity 校验密码复杂度
// 规则：
// 1. 长度不少于 8 位
// 2. 必须包含：大写字母、小写字母、数字、特殊字符
func CheckPasswordComplexity(password string) bool {
	if len(password) < 8 {
		return false
	}
	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasDigit && hasSpecial
}
