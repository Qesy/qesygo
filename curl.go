package qesygo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CurlModel struct {
	Url     string
	Para    map[string]string
	Method  string
	IsHttps bool
	Header  map[string]string
	IsJson  bool
	IsDebug bool
	Body    []byte
}

func (c *CurlModel) SetUrl(Url string) *CurlModel { //设置网址
	c.Url = Url
	return c
}

func (c *CurlModel) SetPara(Para map[string]string) *CurlModel { //设置参数
	c.Para = Para
	return c
}

func (c *CurlModel) SetBody(Body []byte) *CurlModel { //设置参数
	c.Body = Body
	return c
}

func (c *CurlModel) SetMethod(Method string) *CurlModel { //设置参数
	c.Method = Method
	return c
}

func (c *CurlModel) SetIsHttps(IsHttps bool) *CurlModel { //设置参数
	c.IsHttps = IsHttps
	return c
}

func (c *CurlModel) SetHeader(Header map[string]string) *CurlModel { //设置参数
	c.Header = Header
	return c
}

func (c *CurlModel) SetIsJson(IsJson bool) *CurlModel { //设置参数
	c.IsJson = IsJson
	return c
}

func (c *CurlModel) SetIsDebug(IsDebug bool) *CurlModel { //设置参数
	c.IsDebug = IsDebug
	return c
}

func (c *CurlModel) ExecGet() ([]byte, error) { //设置参数
	GetStr := Http_build_query(c.Para)
	resp, err := http.Get(c.Url + "?" + GetStr)
	fmt.Println("Url", c.Url+"?"+GetStr)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c *CurlModel) ExecPost() ([]byte, error) { //设置参数
	Json, _ := JsonEncode(c.Para)
	if c.IsJson && len(c.Body) > 0 {
		Json = c.Body
	}
	req, err := http.NewRequest("POST", c.Url, bytes.NewReader(Json))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	clinet := &http.Client{}
	resp, err := clinet.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)

	// resp, err := http.Post(c.Url, "application/json", bytes.NewReader(Json))
	// if err != nil {
	// 	return []byte{}, err
	// }
	// defer resp.Body.Close()
	// return ioutil.ReadAll(resp.Body)
}
