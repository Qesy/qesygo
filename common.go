package QesyGo

import(
	"math/rand"
	"fmt"
	"io/ioutil"
	"encoding/json"
	//"strings"
)

type RandWeiht struct {
    Name string
    Weight int
}

type RandWeihtArr []RandWeiht


func Substr(str string, start int, end int) string { 
	var endNum int
	s := []byte(str)
	if end > 0{
		endNum = start+ end
	}else{
		endNum = len(str) + end
	}
	return string(s[start:endNum])
}

func Rate(num int) bool{
	rand := rand.Intn(100)+1
	if(rand <= num){
		return true
	}else{
		return false
	}
}

/*
* RandWeihtArr := &lib.RandWeihtArr{{"user1",8}, {"user2",1},{"user3",1}}
* who := RandWeihtArr.RandWeight()
*/
func (arr *RandWeihtArr) RandWeight() string{
	var all int
	for _, v := range *arr{
		all += v.Weight
	}
	plusNum := 0
	tempArr := make(map[string][2]int)
	for _, v := range *arr{
		plusNum += v.Weight
		tempArr[v.Name] = [2]int{plusNum-v.Weight, plusNum}
	}
	randNum := rand.Intn(all)+1
	var ret string
	for k, v := range tempArr{
		if(randNum > v[0] && randNum <= v[1]){
			ret = k
			break
		}
	}
	return ret
}

func ReadFile(str string) ([]byte, error){
	return ioutil.ReadFile(str)
}

func JsonEncode(arr interface{}) ([]byte, error) {
	return json.Marshal(arr)
}

func JsonDecode(str []byte, jsonArr interface{}) error {
	err := json.Unmarshal(str, jsonArr)
    return err

}

func Printf(robots string){
	fmt.Printf("%s", robots)
}