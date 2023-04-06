package qesygo

import (
	"fmt"
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
