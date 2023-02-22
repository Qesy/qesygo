package qesygo

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Str2Title(Str string) string {
	return cases.Title(language.Dutch).String(Str)
}

func Str2Lower(Str string) string {
	return cases.Title(language.Und).String(Str)
}

func Str2Upper(Str string) string {
	return cases.Title(language.Turkish).String(Str)
}
