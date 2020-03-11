package Logger

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

var IsDevelopment = true

func InitLogger() {

	lumberjackLogrotate := &lumberjack.Logger{
		Filename:   "log/ogamebot.log",
		MaxSize:    5,  // Max megabytes before log is rotated
		MaxBackups: 10, // Max number of old log files to keep
		MaxAge:     7,  // Max number of days to retain log files
		Compress:   false,
	}

	var logMultiWriter io.Writer
	if IsDevelopment {
		log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", ForceColors: true})

		logMultiWriter = io.MultiWriter(os.Stdout)
	} else {
		log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", ForceColors: true})

		logMultiWriter = io.MultiWriter(os.Stdout, lumberjackLogrotate)
	}

	log.SetOutput(logMultiWriter)

	log.WithFields(log.Fields{
		"Runtime Version": runtime.Version(),
		"Number of CPUs":  runtime.NumCPU(),
		"Arch":            runtime.GOARCH,
	}).Info("Application Initializing")

}

func LogFileData() (string, int) {
	_, f, l, _ := runtime.Caller(2)
	return filepath.Base(f), l
}
