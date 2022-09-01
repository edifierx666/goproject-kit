package klog

import (
  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
  "gopkg.in/natefinch/lumberjack.v2"
)

type FileCore struct {
  Encoder     zapcore.Encoder
  EnableLevel zapcore.LevelEnabler
  FileLogger  *lumberjack.Logger
}

func NewFileCore() *FileCore {
  return (&FileCore{}).WithDefaultEncoder().WithDefaultEnableLevel().WithDefaultFileLogger()
}

func (fc *FileCore) WithDefaultEncoder() *FileCore {
  fc.Encoder = GetConsoleEncoder(GetDefEncoderCfg())
  return fc
}

func (fc *FileCore) WithDefaultEnableLevel() *FileCore {
  fc.EnableLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
  return fc
}

func (fc *FileCore) WithDefaultFileLogger() *FileCore {
  fc.FileLogger = GetFileLogger()
  return fc
}

func (fc *FileCore) SetEncoder(enc zapcore.Encoder) *FileCore {
  fc.Encoder = enc
  return fc
}

func (fc *FileCore) SetFileLogger(lg *lumberjack.Logger) *FileCore {
  fc.FileLogger = lg
  return fc
}

func (fc *FileCore) SetEnableLevel(enab zapcore.LevelEnabler) *FileCore {
  fc.EnableLevel = enab
  return fc
}

func (fc *FileCore) Build() zapcore.Core {
  return ZapFileCore(fc.Encoder, fc.EnableLevel, fc.FileLogger)
}

func ZapFileCore(enc zapcore.Encoder, enab zapcore.LevelEnabler, fileLogger ...*lumberjack.Logger) zapcore.Core {
  return zapcore.NewCore(enc, GetZapFileWriteSyncer(fileLogger...), enab)
}

func GetZapFileWriteSyncer(fileLogger ...*lumberjack.Logger) zapcore.WriteSyncer {
  var filelog *lumberjack.Logger
  if len(fileLogger) == 0 {
    filelog = GetFileLogger()
  } else {
    filelog = fileLogger[0]
  }
  return zapcore.AddSync(filelog)
}

func GetFileLogger() *lumberjack.Logger {
  return &lumberjack.Logger{
    Filename:   "./log",
    MaxSize:    500,
    MaxAge:     90,
    MaxBackups: 3,
    Compress:   false,
  }
}
