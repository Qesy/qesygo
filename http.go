package QesyGo

import(
	"net/http"
	"io/ioutil"
	"strings"	
	//"fmt"
	//"encoding/json"
)

func Get(url string, para map[string]string) (string, error) {
	paraStr := Http_build_query(para)
	res, err := http.Get(url+"?"+paraStr)
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

func Post(url string, para map[string]string) (string, error) {
	paraStr := Http_build_query(para)
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

/*
*  para := map[string]string{"sex":"boy", "age":"18", "name":"老钱"}
*/
func Http_build_query(para map[string]string) string{
	var arr []string
	for key, val := range para{
		arr = append(arr, key +"="+ val);
	}
	return strings.Join(arr, "&")
}