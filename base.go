package qesygo

import (
	"log"
	"os"
)

func Log(str string, logLev string) {
	FileName := "./static/log/" + logLev + "_" + Date(0, "20060102") + ".txt"
	logfile, err := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer logfile.Close()
	logger := log.New(logfile, "logger", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(str)

}
