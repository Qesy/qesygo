package QesyGo

import (
	"log"
	"os"
	"runtime/debug"
)

func Log(str string, logLev string) {
	FileName := Date(Time("Second"), "20060102")
	logfile, err := os.OpenFile(logLev+"_"+FileName+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		os.Exit(-1)
	}
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(str)
	debug.PrintStack()
}
