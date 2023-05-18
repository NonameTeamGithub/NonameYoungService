package logger

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var logger zerolog.Logger
var once sync.Once

func GetLogger() zerolog.Logger {
	once.Do(func() {
		c := color.New(color.FgRed)
		logger = zerolog.New(
			zerolog.ConsoleWriter{
				Out:        os.Stderr,
				TimeFormat: time.UTC.String(),
				FormatCaller: func(i interface{}) string {
					return "|" + filepath.Base(fmt.Sprintf("%s|", i))
				},
				FormatErrFieldName: func(i interface{}) string {
					return c.Sprintf("|" + strings.ToUpper(fmt.Sprintf("[%s] -> ", i)))
				},
			}).Level(zerolog.InfoLevel).With().Timestamp().Caller().Logger()
	})
	return logger
}
