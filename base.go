package QesyGo

import (
	"log"
	"os"
)

func Log(str string) {
	FileName := Date(Time("Second"), "20060102")
	logfile, err := os.OpenFile("log_"+FileName+".log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		os.Exit(-1)
	}
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(str)
}
