package logger

import (
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

// Configure takes logging parameters and allows custom parameters.
// file: 		Filepath to where the logging file is saved.
// maxsize: 	Size of the file allowed in megabytes.
// maxbackups: 	Maximum allowed backups of a single log file.
// maxage:		Maximum allowed age of a log file.
// compress:	Compress the logfile?
// timefield: 	zerolog time format
func Configure(file string, maxsize, maxbackups, maxage int, compress bool, timefield string) {
	zerolog.TimeFieldFormat = timefield
	l = zerolog.New(&lumberjack.Logger{
		Filename:   file,
		MaxSize:    maxsize,
		MaxBackups: maxbackups,
		MaxAge:     maxage,
		Compress:   compress,
	}).With().Timestamp().Logger()
}

func StdOutput() {
	l = zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger()
	l.Info().Msg("Testing Stdout")
}

func DefaultOutput() {
	l = zerolog.New(&lumberjack.Logger{
		Filename:   appdir + defdir + os.Args[0] + time.Now().UTC().String() + ".log",
		MaxSize:    128,
		MaxBackups: 3,
		MaxAge:     1,
		Compress:   true,
	}).With().Timestamp().Logger()
}
