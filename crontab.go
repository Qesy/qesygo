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
