package log4go

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/agtorre/gocolorize"
)

var (
	LogLevel string
	levelMap = map[string]int{
		"debug": 0,
		"info":  1,
		"warn":  2,
		"error": 3,
	}
	levelColorMap = map[string]gocolorize.Colorize{
		"debug": gocolorize.NewColor("magenta"),
		"info":  gocolorize.NewColor("white"),
		"warn":  gocolorize.NewColor("yellow"),
		"error": gocolorize.NewColor("red"),
	}
)

type Logger struct {
	DEBUG, INFO, WARN, ERROR, logger *log.Logger
}

type l4g struct {
	c gocolorize.Colorize
	w io.Writer
}

func (r *l4g) Write(p []byte) (n int, err error) {
	return r.w.Write([]byte(r.c.Paint(string(p))))
}

func NewLogger(level string) *Logger {
	logger := &Logger{}
	var loglog *log.Logger
	var levelint int
	levelint = levelMap[level]
	for k, v := range levelMap {
		if v >= levelint {
			loglog = log.New(&l4g{levelColorMap[k], os.Stdout}, k, log.Ldate|log.Ltime|log.Lshortfile)
		} else {
			loglog = log.New(ioutil.Discard, k, log.Ldate|log.Ltime|log.Lshortfile)
		}
		switch k {
		case "debug":
			logger.DEBUG = loglog
		case "info":
			logger.INFO = loglog
		case "warn":
			logger.WARN = loglog
		case "error":
			logger.ERROR = loglog
		}
	}
	return logger
}
