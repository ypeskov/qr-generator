package logger

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"ypeskov/qr-generator/internal/config"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	*log.Logger
}

type ContextHook struct{}

func (hook ContextHook) Levels() []log.Level {
	return log.AllLevels
}

var projectRoot string

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		panic("unable to determine working directory: " + err.Error())
	}
	projectRoot = filepath.ToSlash(cwd) + "/"
}

func (hook ContextHook) Fire(entry *log.Entry) error {
	for i := 4; i < 10; i++ {
		if pc, file, line, ok := runtime.Caller(i); ok {
			funcName := runtime.FuncForPC(pc).Name()
			funcName = path.Base(funcName)

			file = strings.TrimPrefix(file, projectRoot)

			if !isLogrusFile(file) {
				entry.Data["func"] = funcName
				entry.Data["file"] = file
				entry.Data["line"] = line
				break
			}
		}
	}
	return nil
}

func isLogrusFile(file string) bool {
	return strings.Contains(file, "sirupsen/logrus")
}

func New(cfg *config.Config) *Logger {
	l := log.New()

	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Warnf("Invalid log level '%s'. Using 'info' level as default.", cfg.LogLevel)
		level = log.InfoLevel
	}
	l.SetLevel(level)
	log.SetReportCaller(true)
	l.AddHook(ContextHook{})

	return &Logger{l}
}
