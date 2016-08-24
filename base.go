package QesyGo

func GetConf(fileName string) map[string]map[string]string {
    str, err := ReadFile(fileName)
    if(err != nil){
        Die(err)
    }    
    confRs := map[string]map[string]string{}
    err = JsonDecode(str, &confRs)
    if(err != nil){
        Printf("config reload error: %s", err)
    }
    return confRs
}