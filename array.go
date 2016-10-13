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
