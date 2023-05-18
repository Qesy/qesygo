package qesygo

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Get(url string, para map[string]string) (string, error) {
	paraStr := Http_build_query(para)
	res, err := http.Get(url + "?" + paraStr)
	if err != nil {
		return "", err
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	return string(robots), err
}

func Post(url string, paraInterface interface{}) (string, error) {
	paraStr := ""
	if para, ok := paraInterface.(map[string]string); ok {
		paraStr = Http_build_query(para)
	} else if para, ok := paraInterface.(string); ok {
		paraStr = para
	}

	res, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(paraStr))
	if err != nil {
		return "", err
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	return string(robots), err
}

func PostJson(url string, paraInterface interface{}) (string, error) {
	Json, _ := JsonEncode(paraInterface)
	res, err := http.Post(url, "application/json", strings.NewReader(string(Json)))
	if err != nil {
		return "", err
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	return string(robots), err
}

/*
*  para := map[string]string{"sex":"boy", "age":"18", "name":"老钱"}
 */
func Http_build_query(para map[string]string) string {
	var arr []string
	for key, val := range para {
		arr = append(arr, key+"="+UrlEnCode(val))
	}
	return strings.Join(arr, "&")
}

func UnHttp_build_query(Str string) map[string]string { // 反Http_build_query
	Arr := []string{}
	if Str != "" {
		Arr = strings.Split(Str, "&")
	}
	Result := map[string]string{}
	for _, v := range Arr {
		Tmp := strings.Split(v, "=")
		Result[Tmp[0]] = Tmp[1]
	}
	return Result
}

func UrlEnCode(str string) string {
	return url.QueryEscape(str)
}
