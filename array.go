package qesygo

import (
	"math/rand"
	"strconv"
	"time"
)

func Array_merge(arr ...[]string) []string {
	var retArr []string
	for _, n := range arr {
		for _, v := range n {
			retArr = append(retArr, v)
		}
	}
	return retArr
}

func Array_merge_int(arr ...[]int) []int {
	var retArr []int
	for _, n := range arr {
		for _, v := range n {
			retArr = append(retArr, v)
		}
	}
	return retArr
}

func InArray(Arr interface{}, str string) bool {
	if strArr, ok := Arr.([]string); ok {
		for _, v := range strArr {
			if v == str {
				return true
			}
		}
		return false
	}
	if mapArr, ok := Arr.(map[string]string); ok {
		for _, v := range mapArr {
			if v == str {
				return true
			}
		}
		return false
	}
	if intArr, ok := Arr.([]int); ok {
		strInt := StrToInt(str)
		for _, v := range intArr {
			if v == strInt {
				return true
			}
		}
		return false
	}
	if intArr, ok := Arr.([]int32); ok {
		strInt32 := StrToInt32(str)
		for _, v := range intArr {
			if v == strInt32 {
				return true
			}
		}
		return false
	}
	return false

}

func Array_Rand(num int) []int {
	rand.Seed(time.Now().UnixNano())
	return rand.Perm(num)
}

func Array_Diff(Arr1 []int, Arr2 []int) []int {
	BigArr, SmallArr := Arr1, Arr2
	if len(Arr2) > len(Arr1) {
		BigArr = Arr2
		SmallArr = Arr1
	}
	TempArr := []int{}
	for _, v := range BigArr {
		if !InArray(SmallArr, strconv.Itoa(v)) {
			TempArr = append(TempArr, v)
		}
	}
	return TempArr
}

func Array_Diff_String(Arr1 []string, Arr2 []string) []string {
	BigArr, SmallArr := Arr1, Arr2
	if len(Arr2) > len(Arr1) {
		BigArr = Arr2
		SmallArr = Arr1
	}
	TempArr := []string{}
	for _, v := range BigArr {
		if !InArray(SmallArr, v) {
			TempArr = append(TempArr, v)
		}
	}
	return TempArr
}

func Array_Mixed(Arr1 []int, Arr2 []int) []int {
	BigArr, SmallArr := Arr1, Arr2
	if len(Arr2) > len(Arr1) {
		BigArr = Arr2
		SmallArr = Arr1
	}
	TempArr := []int{}
	for _, v := range BigArr {
		if InArray(SmallArr, strconv.Itoa(v)) {
			TempArr = append(TempArr, v)
		}
	}
	return TempArr
}

func Array_column(Arr []map[string]string, Str string) []string {
	NewArr := []string{}
	for _, v := range Arr {
		NewArr = append(NewArr, v[Str])
	}
	return NewArr
}

func Array_column(Arr []map[string]string, Str string, Index string) []string {
	NewArr := map[string]string{}
	for _, v := range Arr {
		NewArr[v[Index]] = v[Str]
	}
	return NewArr
}
