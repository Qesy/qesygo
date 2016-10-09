package QesyGo

import (
	"math/rand"
	"strconv"
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

func InArray(Arr interface{}, str string) bool {
	strArr, ok := Arr.([]string)
	if ok {
		for _, v := range strArr {
			if v == str {
				return true
			}
		}
		return false
	}
	mapArr, ok := Arr.(map[string]string)
	if ok {
		for _, v := range mapArr {
			if v == str {
				return true
			}
		}
		return false
	}
	intArr, ok := Arr.([]int)
	if ok {
		for _, v := range intArr {
			if strconv.Itoa(v) == str {
				return true
			}
		}
		return false
	}
	return false

}

func Array_Rand(num int) []int {
	return rand.Perm(num)
}
