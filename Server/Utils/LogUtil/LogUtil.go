package LogUtil

import (
	"github.com/sirupsen/logrus"
	"os"
)

type LogrusErrorWriter struct{}

func New() *logrus.Entry {

	return logrus.NewEntry(&logrus.Logger{
		Out:          os.Stdout,
		Hooks:        nil,
		Formatter:    &logrus.JSONFormatter{},
		ReportCaller: false,
		Level: func() logrus.Level {
			logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
			if err != nil {
				logLevel = logrus.InfoLevel
			}

			return logLevel
		}(),
		ExitFunc: os.Exit,
	})
}
func (w LogrusErrorWriter) Write(p []byte) (n int, err error) {
	logrus.Errorf("%s", string(p))
	return len(p), nil
}
