package qesygo

import "regexp"

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
