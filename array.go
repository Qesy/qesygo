package qesygo

import (
	"math/rand"
	"strconv"
	"time"
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

func Array_merge(arr ...[]string) []string {
	var retArr []string
	for _, n := range arr {
		retArr = append(retArr, n...)
	}
	return retArr
}

func Array_merge_int(arr ...[]int) []int {
	var retArr []int
	for _, n := range arr {
		retArr = append(retArr, n...)
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

func Array_column_index(Arr []map[string]string, Str string, Index string) map[string]string {
	NewArr := map[string]string{}
	for _, v := range Arr {
		NewArr[v[Index]] = v[Str]
	}
	return NewArr
}

func Array_unique(Arr []string) []string {
	NewArr := []string{}
	for _, v := range Arr {
		if !InArray(NewArr, v) {
			NewArr = append(NewArr, v)
		}
	}
	return NewArr
}

func Keys(Arr map[string]string) []string {
	Keys := []string{}
	for k, _ := range Arr {
		Keys = append(Keys, k)
	}
	return Keys
}

func IntsToStrs(intArr []int) []string {
	strArr := make([]string, len(intArr))
	for i, val := range intArr {
		strArr[i] = strconv.Itoa(val)
	}
	return strArr
}

func StrsToInts(StrArr []string) []int {
	intArr := make([]int, len(StrArr))
	for i, val := range StrArr {
		intArr[i], _ = strconv.Atoi(val)
	}
	return intArr
}
