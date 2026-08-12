package qesygo

import (
	"fmt"
	"io"
	"os"
)

func LogSave(Path string) { // 保存日志
	Yestoday := DateTimeGet() - 86400
	SavePath := Path + "_" + Date(Yestoday, "20060102") + ".log"
	if _, Err := os.Stat(SavePath); Err == nil { //文件存在则不处理
		return
	}
	txtCopy(Path+".log", Path+"_"+Date(Yestoday, "20060102")+".log")
	os.WriteFile(Path+".log", []byte{}, 0666)
}

func txtCopy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
