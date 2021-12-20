package log

import (
	"os"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

var stand string

// InitLogrus custom logrus
func InitLogrus() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  time.RFC822Z,
	})
	stand = os.Getenv("STAND")
	if stand == "" {
		logrus.Fatalln("Don`t find stand env")
	}
}

// CheckFatal default fatal errors checker
func CheckFatal(err error, more ...interface{}) {
	if err == nil {
		return
	}
	if stand == "dev" {
		f, l := detailError()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"func": f,
				"line": l,
			}).Fatalf("%s | %v", err, more)
		}
	} else {
		logrus.Fatalf("%s | %v", err, more)
	}
}

// CheckError default errors checker
func CheckError(err error, more ...interface{}) {
	if err == nil {
		return
	}
	if stand == "dev" {
		f, l := detailError()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"func": f,
				"line": l,
			}).Errorf("%s | %v", err, more)
		}
	} else {
		logrus.Errorf("%s | %v", err, more)
	}
}

func Info(r ...interface{}) {
	logrus.Infoln(r...)
}

func Fata(err error, more ...interface{}) {
	logrus.Fatalln(err, more)
}

func detailError() (string, int) {
	pc, _, l, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name(), l
}
