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
	"os"
	"strconv"
	"strings"

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

func FileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func Max(Para ...int) int {
	Tmp := 0
	for _, v := range Para {
		if Tmp == 0 || v > Tmp {
			Tmp = v
		}
	}
	return Tmp
}

func Min(Para ...int) int {
	Tmp := 0
	for _, v := range Para {
		if Tmp == 0 || v < Tmp {
			Tmp = v
		}
	}
	return Tmp
}
