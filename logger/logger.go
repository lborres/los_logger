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

func (status *Status) writeLog(msg string) {
	// Store data in DB

	// Log message
	log.Println(msg)

	// ? How do I write to log file?
	// logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err!= nil {
	// 	log.Fatalf("Failed to open log file: %v\n", err)
	// }
	// defer logFile.Close()
	// Create a new logger instance with io.MultiWriter
	// logger := log.New(io.MultiWriter(logFile, os.Stdout), "", log.LstdFlags|log.Lshortfile)

	// // Use the logger
	// logger.Println("This message goes to both the log file and stdout.")
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
			status.writeLog(msg)
			status.connected = true
		}
	} else {
		if status.connected {
			status.lastDisconnect = time.Now()
			status.writeLog("Internet Service Interruption")
			status.connected = false
		}
	}
}

func StartLogger(chExit <-chan struct{}, db *sql.DB) {
	status := Status{
		connected: true,
	}

	NewStorage(db)

	for {
		select {
		case <-chExit:
			log.Println("Shutting Down Logger")
			return

		default:
			r := rand.Intn(3000) + 4000

			status.checkService()

			time.Sleep(time.Duration(r) * time.Millisecond)
		}
	}
}
