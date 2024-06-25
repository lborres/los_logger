package logger

import (
	"io"
	"log"
	"os"
)

type LogWriter struct {
	writer io.Writer
	isOpen bool
}

func (l *LogWriter) CloseWriter() error {
	if !l.isOpen {
		return nil
	}
	l.isOpen = false
	return l.writer.(*os.File).Close()
}

func InitLogWriter(fileloc string) *LogWriter {
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	logFile, err := os.OpenFile(fileloc, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file with err: %s", err)
	}

	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	return &LogWriter{
		writer: mw,
		isOpen: true,
	}
}
