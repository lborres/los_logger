package logger

import (
	"log"
	"math/rand"
	"time"
)

func checkService() {
	log.Println("Checking Service")
}

func InitLogger(chExit <-chan struct{}, logfile string) {
	for {
		select {
		case <-chExit:
			log.Println("Logger Shutting Down")
			return

		default:
			r := rand.Intn(2500) + 4500

			checkService()

			log.Printf("Sleep Duration is %dms\n", r)
			time.Sleep(time.Duration(r) * time.Millisecond)
		}
	}
}
