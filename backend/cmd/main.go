package main

import (
	"log/slog"
	"math/rand"
	"net"
	"os"
	"time"
)

const (
	logFileDir  = 
	logFileName = "los_logger" + ".log"
)

func isConnected() bool {
	conn, err := net.DialTimeout("tcp", "8.8.8.8:53", 5*time.Second)
	if err != nil {
		slog.Debug("OFFLINE")
		return false
	}
	conn.Close()
	slog.Debug("ONLINE")
	return true
}

func main() {
	var connected bool = true
	var lastDisconnect time.Time

	// logHandlerOpts := &slog.HandlerOptions{
	// 	Level: slog.LevelDebug,
	// }
	// logger := slog.New(slog.NewTextHandler(os.Stderr, logHandlerOpts))
	// slog.SetDefault(logger)

	logFile, err := os.OpenFile(logFileDir+logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	logHandlerOpts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	logger := slog.New(slog.NewTextHandler(logFile, logHandlerOpts))
	slog.SetDefault(logger)

	slog.Info("PROGRAM START")

	for {
		r := rand.Intn(2500) + 4500
		if isConnected() {
			if !connected {
				duration := time.Since(lastDisconnect)
				slog.Info("Connection restored after", slog.Duration("duration", duration))
				connected = true
			}
		} else {
			if connected {
				lastDisconnect = time.Now()
				slog.Info("Internet Service Interruption")
				connected = false
			}
		}

		// slog.Warn("Internet Service Interruption")
		// slog.Info("Test")

		time.Sleep(time.Duration(r) * time.Millisecond)
	}

}
