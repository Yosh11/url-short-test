package err

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
func CheckFatal(e error, more ...interface{}) {
	if e == nil {
		return
	}
	if stand == "dev" {
		f, l := detailError()
		if e != nil {
			logrus.WithFields(logrus.Fields{
				"func": f,
				"line": l,
			}).Fatalf("%s | %v", e, more)
		}
	} else {
		logrus.Fatalf("%s | %v", e, more)
	}
}
// CheckError default errors checker
func CheckError(e error, more ...interface{}) {
	if e == nil {
		return
	}
	if stand == "dev" {
		f, l := detailError()
		if e != nil {
			logrus.WithFields(logrus.Fields{
				"func": f,
				"line": l,
			}).Errorf("%s | %v", e, more)
		}
	} else {
		logrus.Errorf("%s | %v", e, more)
	}
}

func Info(r ...interface{}) {
	logrus.Infoln(r...)
}

func detailError() (string, int) {
	pc, _, l, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name(), l
}