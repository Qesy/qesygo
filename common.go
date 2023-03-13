package qesygo

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/axgle/mahonia"
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

func Rand(Min int, Max int) int {
	tempNum := Max - Min
	if tempNum <= 0 {
		return Min
	}
	randSeek := time.Now().UnixNano()
	rand.Seed(randSeek)
	time.Sleep(time.Nanosecond * 1)
	return Min + rand.Intn(tempNum)
}

func Rate(num int) bool {
	rand := Rand(1, 101)
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
	randNum := Rand(0, all) + 1
	var ret string
	for k, v := range tempArr {
		if randNum > v[0] && randNum <= v[1] {
			ret = k
			break
		}
	}
	return ret
}

func VeriPara(Req map[string]string, Para []string) bool {
	for _, val := range Para {
		if Str, ok := Req[val]; !ok || Str == "" {
			fmt.Println("VeriPara", val)
			return false
		}
	}
	return true
}

func ReadFile(str string) ([]byte, error) {
	return ioutil.ReadFile(str)
}

func JsonEncode(arr interface{}) ([]byte, error) {
	return json.Marshal(arr)
}

func JsonDecode(str []byte, jsonArr interface{}) error {
	strNew := string(str)
	if strNew == "null" || strNew == "" {
		return nil
	}
	err := json.Unmarshal(str, jsonArr)
	return err

}

func Printf(format string, a ...interface{}) {
	fmt.Printf(format, a)
}

func Fprintf(w http.ResponseWriter, str string) {
	fmt.Fprintf(w, str)
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Die(v interface{}) {
	log.Fatal(v)
}

func Implode(arr []string, sep string) string {
	return strings.Join(arr, sep)
}

func Explode(str string, sep string) []string {
	if str == "" {
		return []string{}
	}
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
	case "Microsecond":
		t = now.UnixNano() / 1000
	case "Millisecond":
		t = now.UnixNano() / 1000000
	case "Second":
		t = now.UnixNano() / 1000000000
	}
	return t
}

func TimeStr(str string) string {
	t := Time(str)
	return strconv.FormatInt(t, 10)
}

func TimeInt(str string) int {
	t := Time(str)
	ret, _ := Int64ToInt(t)
	return ret
}

//-- format : "2006-01-02 03:04:05 PM" --
/*
@timestamp 传0 ，即现在时间

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
	if timestamp == 0 {
		timestamp = Time("Second")
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

func DateTimeGet() int { //获取当天0点时间戳
	t := time.Now()
	timeStr := t.Format("2006-01-02")
	t, _ = time.ParseInLocation("2006-01-02", timeStr, time.Now().Location())
	return int(t.Unix())
}

// -- "01/02/2006", "02/08/2015" --
func StrToTimeByDate(format string, input string) int64 {
	tm2, _ := time.ParseInLocation(format, input, time.Now().Location())
	return tm2.Unix()
}

func StrToTime(format string, input string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(format, input, loc)
	return theTime.Unix()
}

// 转码 ConvertToString( string, "gbk", "utf-8")
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result

}

func Int64ToInt(num int64) (int, error) {
	str := strconv.FormatInt(num, 10)
	return strconv.Atoi(str)
}

func IntToStr(a interface{}) string {
	return fmt.Sprintf("%d", a)
}

func StrToInt(str string) int {
	ret, _ := strconv.Atoi(str)
	return ret
}

func StrToInt32(str string) int32 {
	return int32(StrToInt(str))
}

func StrToInt64(str string) int64 {
	ret, _ := strconv.ParseInt(str, 10, 64)
	return ret
}

func Unset(arr []string, str string) []string {
	newArr := []string{}
	for _, v := range arr {
		if v != str && v != "" {
			newArr = append(newArr, v)
		}
	}
	return newArr
}

// -- kind 0:纯数字，1：小写，2：大写，3：数字+大小写字幕 --
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = Rand(1, 3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + Rand(1, scope))
	}
	return result
}

func Wordwrap(Str string, width int, breakStr string) string {
	breakStrByte := []byte(breakStr)
	strByte := []byte(Str)
	var tempByte []byte
	for k, v := range strByte {
		if (k+1)%width == 0 && k != len(strByte)-1 {
			tempByte = append(tempByte, v)
			tempByte = append(tempByte, breakStrByte...)
		} else {
			tempByte = append(tempByte, v)
		}
	}
	return string(tempByte)
}

func Base64Encode(json string) string {
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	return b64.EncodeToString([]byte(json))
}

func Base64Decode(json string) ([]byte, error) {
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	return b64.DecodeString(json)
}

func Abs(a int) (ret int) {
	ret = (a ^ a>>31) - a>>31
	return
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func SHA1(str string) string {
	s := sha1.Sum([]byte(str))
	strsha1 := hex.EncodeToString(s[:])
	return strsha1
}

func IntToBytes(x int32) []byte {
	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.LittleEndian, x)
	return b_buf.Bytes()
}

func BytesToInt(b []byte) int32 {
	var x int32
	b_buf := bytes.NewBuffer(b)
	binary.Read(b_buf, binary.LittleEndian, &x)
	return x
}
