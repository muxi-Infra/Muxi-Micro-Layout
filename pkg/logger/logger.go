package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

func NewZapLogger() *ZapLogger {
	return NewLogger(os.Stderr, InfoLevel)
}

type Logger interface {
	Debug(msg string, fields ...Field)
	Debugf(format string, args ...interface{})
	Info(msg string, fields ...Field)
	Infof(format string, args ...interface{})
	Warn(msg string, fields ...Field)
	Warnf(format string, args ...interface{})
	Error(msg string, fields ...Field)
	Errorf(format string, args ...interface{})
	Panic(msg string, fields ...Field)
	Panicf(format string, args ...interface{})
	Fatal(msg string, fields ...Field)
	Fatalf(format string, args ...interface{})
	SetLevel(level Level)
	Sync() error
}

type ZapLogger struct {
	l  *zap.Logger
	al *zap.AtomicLevel
}

func NewLogger(out io.Writer, level Level) *ZapLogger {
	if out == nil {
		out = os.Stderr
	}

	al := zap.NewAtomicLevelAt(level)
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.RFC3339TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zapcore.AddSync(out),
		al,
	)
	return &ZapLogger{l: zap.New(core), al: &al}
}

func (l *ZapLogger) SetLevel(level Level) {
	if l.al != nil {
		l.al.SetLevel(level)
	}
}

func (l *ZapLogger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

func (l *ZapLogger) Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.l.Debug(msg, nil...)
}

func (l *ZapLogger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func (l *ZapLogger) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.l.Info(msg, nil...)
}

func (l *ZapLogger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *ZapLogger) Warnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.l.Warn(msg, nil...)
}

func (l *ZapLogger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.l.Error(msg, nil...)
}

func (l *ZapLogger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}

func (l *ZapLogger) Panicf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.l.Panic(msg, nil...)
}

func (l *ZapLogger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

func (l *ZapLogger) Fatalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.l.Fatal(msg, nil...)
}

func (l *ZapLogger) Sync() error {
	return l.l.Sync()
}
