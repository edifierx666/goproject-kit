package inner

import (
  "gopkg.in/natefinch/lumberjack.v2"
)

type BaseFileLogger struct {
  logger *lumberjack.Logger
}

func NewFileLogger(encoders ...*lumberjack.Logger) *BaseFileLogger {
  var inner *lumberjack.Logger
  if len(encoders) >= 1 {
    inner = encoders[0]
  } else {
    inner = DefFileLogger()
  }

  return &BaseFileLogger{logger: inner}
}

func DefFileLogger() *lumberjack.Logger {
  return &lumberjack.Logger{
    Filename:   "./log",
    MaxSize:    100,
    MaxAge:     100,
    MaxBackups: 3,
    Compress:   false,
  }
}

func (b *BaseFileLogger) Filename(filename string) *BaseFileLogger {
  b.logger.Filename = filename
  return b
}

func (b *BaseFileLogger) MaxSize(maxSize int) *BaseFileLogger {
  b.logger.MaxSize = maxSize
  return b
}

func (b *BaseFileLogger) MaxAge(maxAge int) *BaseFileLogger {
  b.logger.MaxAge = maxAge
  return b
}

func (b *BaseFileLogger) MaxBackups(maxBackups int) *BaseFileLogger {
  b.logger.MaxBackups = maxBackups
  return b
}

func (b *BaseFileLogger) LocalTime(localTime bool) *BaseFileLogger {
  b.logger.LocalTime = localTime
  return b
}

func (b *BaseFileLogger) Compress(compress bool) *BaseFileLogger {
  b.logger.Compress = compress
  return b
}

func (b *BaseFileLogger) Build() *lumberjack.Logger {
  return b.logger
}
