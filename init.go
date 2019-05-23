package logger

import (
	"os"
	"runtime"

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
	DefaultFileOutput()
}
