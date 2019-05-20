package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Configure(file string, maxsize, maxbackups, maxage int32, compress bool, timefield string) {
	zerolog.TimeFieldFormat = timefield
	l = zerolog.New(&lumberjack.Logger{
		Filename:   file,
		MaxSize:    maxsize,
		MaxBackups: maxbackups,
		MaxAge:     maxage,
		Compress:   compress,
	}).With().Timestamp().Logger()
}
