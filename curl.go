package qesygo

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

type CurlModel struct {
	Url      string
	Para     map[string]string
	Method   string
	IsHttps  bool
	Header   map[string]string
	IsJson   bool
	IsDebug  bool
	Body     []byte
	IsCert   bool //是否验证证书
	CertFile string
	KeyFile  string
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

func (c *CurlModel) SetCert(certFile string, keyFile string) *CurlModel { //设置实用证书
	c.CertFile = certFile
	c.KeyFile = keyFile
	c.IsCert = true
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
	return io.ReadAll(resp.Body)
}

func (c *CurlModel) ExecPost() ([]byte, error) { //设置参数
	b := []byte{}
	if c.IsJson {
		b, _ = JsonEncode(c.Para)
	} else {
		b = []byte(Http_build_query(c.Para))
	}

	if len(c.Body) > 0 {
		b = c.Body
	}
	req, err := http.NewRequest("POST", c.Url, bytes.NewReader(b))
	if c.IsJson {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	clinet := &http.Client{}
	if c.IsCert {
		var cliCrt tls.Certificate
		cliCrt, err = tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
		if err != nil {
			return []byte{}, err
		}
		clinet = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{Certificates: []tls.Certificate{cliCrt}},
			},
		}
	}
	if c.IsDebug {
		fmt.Println("DEBUG", string(b))
	}

	resp, err := clinet.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)

	// resp, err := http.Post(c.Url, "application/json", bytes.NewReader(Json))
	// if err != nil {
	// 	return []byte{}, err
	// }
	// defer resp.Body.Close()
	// return ioutil.ReadAll(resp.Body)
}
