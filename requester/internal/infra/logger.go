package infra

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger interface {
	Info(msg string)
	Debug(msg string)
	Error(msg string)
}

type ZerologLogger struct{}

func NewLogger() Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	return &ZerologLogger{}
}

func (l *ZerologLogger) Info(msg string) {
	log.Info().Msg(msg)
}

func (l *ZerologLogger) Debug(msg string) {
	log.Debug().Msg(msg)
}

func (l *ZerologLogger) Error(msg string) {
	log.Error().Msg(msg)
}
