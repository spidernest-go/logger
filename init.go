package logger

import (
	"os"
	"runtime"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

var (
	l      zerolog.Logger
	appdir string
	defdir string
)

func init() {
	// Find the platform-specific directory
	switch runtime.GOOS {
	case "windows":
		appdir = os.Getenv("APPDATA")
		defdir = "\\default\\"
	default:
		appdir = "/var/log"
		defdir = "/default/"
	}

	// Allocate a new logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	l = zerolog.New(&lumberjack.Logger{
		Filename:   appdir + defdir + os.Args[0] + time.Now().UTC().String() + ".log",
		MaxSize:    128,
		MaxBackups: 3,
		MaxAge:     1,
		Compress:   true,
	}).With().Timestamp().Logger()
}
