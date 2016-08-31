package QesyGo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type RandWeiht struct {
	Name   string
	Weight int
}

type RandWeihtArr []RandWeiht

func Substr(str string, start int, end int) string {
	var endNum int
	s := []byte(str)
	if end > 0 {
		endNum = start + end
	} else {
		endNum = len(str) + end
	}
	return string(s[start:endNum])
}

func Rate(num int) bool {
	rand := rand.Intn(100) + 1
	if rand <= num {
		return true
	} else {
		return false
	}
}

/*
* RandWeihtArr := &lib.RandWeihtArr{{"user1",8}, {"user2",1},{"user3",1}}
* who := RandWeihtArr.RandWeight()
 */
func (arr *RandWeihtArr) RandWeight() string {
	var all int
	for _, v := range *arr {
		all += v.Weight
	}
	plusNum := 0
	tempArr := make(map[string][2]int)
	for _, v := range *arr {
		plusNum += v.Weight
		tempArr[v.Name] = [2]int{plusNum - v.Weight, plusNum}
	}
	randNum := rand.Intn(all) + 1
	var ret string
	for k, v := range tempArr {
		if randNum > v[0] && randNum <= v[1] {
			ret = k
			break
		}
	}
	return ret
}

func ReadFile(str string) ([]byte, error) {
	return ioutil.ReadFile(str)
}

func JsonEncode(arr interface{}) ([]byte, error) {
	return json.Marshal(arr)
}

func JsonDecode(str []byte, jsonArr interface{}) error {
	err := json.Unmarshal(str, jsonArr)
	return err

}

func Printf(format string, a ...interface{}) {
	fmt.Printf(format, a)
}

func Fprintf(w http.ResponseWriter, str string) {
	fmt.Fprintf(w, str)
}

func Die(v interface{}) {
	log.Fatal(v)
}

func Implode(arr []string, sep string) string {
	return strings.Join(arr, sep)
}

func Explode(str string, sep string) []string {
	return strings.Split(str, sep)
}

func Err(str string) error {
	return errors.New(str)
}

func Println(str ...interface{}) {
	fmt.Println(str)
}

func Time(str string) int64 {
	now := time.Now()
	t := now.UnixNano()
	switch str {
	case "Millisecond":
		t = now.UnixNano() / 1000
	case "Microsecond":
		t = now.UnixNano() / 1000000
	case "Second":
		t = now.UnixNano() / 1000000000
	}
	return t
}

//-- format : "2006-01-02 03:04:05 PM" --
/*
月份 1,01,Jan,January
日　 2,02,_2
时　 3,03,15,PM,pm,AM,am
分　 4,04
秒　 5,05
年　 06,2006
周几 Mon,Monday
时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
时区字母缩写 MST
*/
func Date(timestamp int64, format string) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

//-- "01/02/2006", "02/08/2015" --
func StrToTime(format string, input string) int64 {
	tm2, _ := time.Parse(format, input)
	return tm2.Unix()
}
