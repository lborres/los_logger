package logger

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

type Status struct {
	connected      bool
	lastDisconnect time.Time
	duration       time.Duration
}

// TODO Check if Machine is connected to router
// func pingLocal() {
// }

func pingWeb() bool {
	// TODO Save address to ENV Var
	conn, err := net.DialTimeout("tcp", "8.8.8.8:53", 5*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func (status *Status) checkService() {
	if pingWeb() {
		if !status.connected {
			status.duration = time.Since(status.lastDisconnect)

			msg := fmt.Sprintf("Connection restored after %s", status.duration)
			log.Println(msg)

			status.connected = true
		}
	} else {
		if status.connected {
			status.lastDisconnect = time.Now()

			log.Println("Internet Service Interruption")

			status.connected = false
		}
	}
}

func StartLOSLogger(db *sql.DB) {
	status := Status{
		connected: true,
	}

	NewStorage(db)

	for {
		r := rand.Intn(3000) + 4000

		status.checkService()

		time.Sleep(time.Duration(r) * time.Millisecond)

	}

}
