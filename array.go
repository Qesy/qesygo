package qesygo

import (
	"math/rand"
	"strconv"
)

/*
people := []Person{
        {"Bob", 31},
        {"John", 42},
        {"Michael", 17},
        {"Jenny", 26},
    }

    fmt.Println(people)
    // There are two ways to sort a slice. First, one can define
    // a set of methods for the slice type, as with ByAge, and
    // call sort.Sort. In this first example we use that technique.
    sort.Sort(ByAge(people))
    fmt.Println(people)
*/

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func Array_merge(arr ...[]string) []string { // 合并字符串数组
	var retArr []string
	for _, n := range arr {
		retArr = append(retArr, n...)
	}
	return retArr
}

func Array_merge_int(arr ...[]int) []int { // 合并int数组
	var retArr []int
	for _, n := range arr {
		retArr = append(retArr, n...)
	}
	return retArr
}

func InArray(Arr interface{}, str string) bool { // 数组内是否包含
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
	return rand.Perm(num)
}

func Array_Diff(Arr1 []int, Arr2 []int) []int { // 获取不在B数组内的A的值
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

func Array_Diff_String(Arr1 []string, Arr2 []string) []string { // 获取不在B数组内的A的值
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

func Array_Mixed(Arr1 []int, Arr2 []int) []int { // 获取在B数组内A数组的值
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

func Array_Mixed_String(Arr1 []string, Arr2 []string) []string { // 获取在B数组内A数组的值
	BigArr, SmallArr := Arr1, Arr2
	if len(Arr2) > len(Arr1) {
		BigArr = Arr2
		SmallArr = Arr1
	}
	TempArr := []string{}
	for _, v := range BigArr {
		if InArray(SmallArr, v) {
			TempArr = append(TempArr, v)
		}
	}
	return TempArr
}

func Array_column(Arr []map[string]string, Str string) []string { // 获取[]map[string]string{}中指定key的值
	NewArr := []string{}
	for _, v := range Arr {
		NewArr = append(NewArr, v[Str])
	}
	return NewArr
}

func Array_column_index(Arr []map[string]string, Str string, Index string) map[string]string { // 获取[]map[string]string{}中指定key的值带索引
	NewArr := map[string]string{}
	for _, v := range Arr {
		NewArr[v[Index]] = v[Str]
	}
	return NewArr
}

func Array_unique(Arr []string) []string { // 过滤数组中重复的值
	NewArr := []string{}
	for _, v := range Arr {
		if !InArray(NewArr, v) {
			NewArr = append(NewArr, v)
		}
	}
	return NewArr
}

func Keys(Arr map[string]string) []string { // 获取数组中的Key
	Keys := []string{}
	for k := range Arr {
		Keys = append(Keys, k)
	}
	return Keys
}

func IntsToStrs(intArr []int) []string { // 把数值数组转为字符串数组
	strArr := make([]string, len(intArr))
	for i, val := range intArr {
		strArr[i] = strconv.Itoa(val)
	}
	return strArr
}

func StrsToInts(StrArr []string) []int { // 把字符串数组转为数值数组
	intArr := make([]int, len(StrArr))
	for i, val := range StrArr {
		intArr[i], _ = strconv.Atoi(val)
	}
	return intArr
}
