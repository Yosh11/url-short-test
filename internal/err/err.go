package err

import (
	"time"

	"github.com/sirupsen/logrus"
)

// InitLogrus custom logrus
func InitLogrus() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  time.RFC822Z,
	})
}

// CheckFatal default fatal errors checker
func CheckFatal(e error, ifErr ...interface{}) {
	if e != nil {
		// logrus.WithFields(logrus.Fields{ todo upgrade logger to write line & file where we have error (not logrus realise)
		// 	"file": "main",
		// 	"line": 42,
		// }).Fatalln(e)
		logrus.Fatalf("%v => %s", ifErr, e)
	}
}
// CheckError default errors checker
func CheckError(e error, ifErr ...interface{}) {
	if e != nil {
		logrus.Errorf("%v => %s", ifErr, e)
	}
}

func Info(r ...interface{}) {
	logrus.Infoln(r...)
}