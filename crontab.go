package QesyGo

import (
	"time"
)

func Crontab(f func()) {
	Println("Server Cron Run...... ")
	timer1 := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer1.C:
			f()
		}
	}
}

func CrontabMicro(f func()) {
	Println("Server CrontabMicro Run...... ")
	timer1 := time.NewTicker(1 * time.Millisecond)
	for {
		select {
		case <-timer1.C:
			f()
		}
	}
}
