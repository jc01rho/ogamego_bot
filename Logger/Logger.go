package Logger

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
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
		log.SetLevel(log.TraceLevel)

		logMultiWriter = io.MultiWriter(os.Stdout)
	} else {
		log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
		log.SetLevel(log.DebugLevel)

		logMultiWriter = io.MultiWriter(os.Stdout, lumberjackLogrotate)
	}

	log.SetOutput(logMultiWriter)

	log.WithFields(log.Fields{
		"Runtime Version": runtime.Version(),
		"Number of CPUs":  runtime.NumCPU(),
		"Arch":            runtime.GOARCH,
	}).Info("Application Initializing")

}

func CurrentFileNameAndLine() []interface{} {
	//value interface{}
	_, b, c, _ := runtime.Caller(1)
	arrays := []interface{}{}
	arrays = append(arrays, b)
	arrays = append(arrays, c)
	return arrays
}
