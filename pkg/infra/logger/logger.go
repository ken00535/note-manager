package logger

import (
	"os"
	"path"
	"sync"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

type hook struct{}

var (
	log     = logrus.New()
	fileLog = logrus.New()
	once    sync.Once
)

func initLogger() {
	// set file log
	logPath, exist := os.LookupEnv("LOG_PATH")
	if !exist {
		logPath = "log"
	}
	file, err := os.OpenFile(path.Join(logPath, "log.log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	fileLog.SetLevel(logrus.TraceLevel)
	formatterForFile := runtime.Formatter{
		ChildFormatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%    \tfile=%file% line=%line%\n",
		},
		Line: true,
		File: true,
	}
	fileLog.SetOutput(file)
	fileLog.SetFormatter(&formatterForFile)
	// set main log
	log.SetLevel(logrus.InfoLevel)
	formatterConsole := runtime.Formatter{
		ChildFormatter: &logrus.TextFormatter{
			TimestampFormat: "2006-01-02T15:04:05.999-07:00",
			ForceColors:     true,
		},
		Line: true,
		File: true,
	}
	log.SetFormatter(&formatterConsole)
	log.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	log.AddHook(&hook{})
}

// Levels is used to satisfy the interface logrus.Hook
func (m *hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire is used to satisfy the interface logrus.Hook
func (m *hook) Fire(entry *logrus.Entry) error {
	if entry.Level > logrus.FatalLevel {
		fileLog.Log(entry.Level, entry.Message)
	} else if entry.Level == logrus.PanicLevel {
		defer func() {
			recover()
		}()
		fileLog.Log(entry.Level, entry.Message)
	}
	return nil
}

// New a logger
func New() Logger {
	once.Do(initLogger)
	return log
}
