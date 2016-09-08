package QesyGo

func Array_merge(arr ...[]string) []string {
	var retArr []string
	for _, n := range arr {
		for _, v := range n {
			retArr = append(retArr, v)
		}
	}
	return retArr
}
