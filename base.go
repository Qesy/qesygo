package qesygo

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Log() {
	logFile, err := os.OpenFile("./static/log/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)
	log.SetOutput(logFile)
	fmt.Println("Log Start Success !")
}

func LogSave() { // 保存日志
	Src := "./static/log/error.log"
	copy(Src, "./static/log/error_"+Date(0, "20060102")+".log")
	os.WriteFile(Src, []byte{}, 0666)
}

func copy(src, dst string) (int64, error) {
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
