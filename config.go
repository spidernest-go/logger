package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

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
