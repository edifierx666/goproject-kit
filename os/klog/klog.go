package klog

import (
  "strings"

  "github.com/edifierx666/goproject-kit/os/klog/core"
  "github.com/edifierx666/goproject-kit/os/klog/encoder"
  "github.com/edifierx666/goproject-kit/os/klog/inner"
  "github.com/edifierx666/goproject-kit/os/klog/writesyncer"
  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
)

type Logger struct {
  *zap.Logger
  encoderCfg *encoder.KlogEncoder
  loggerCfg  *LoggerCfg
}

func New() *Logger {
  return &Logger{
    Logger:     nil,
    encoderCfg: encoder.NewKlogEncoderCfg(),
    loggerCfg:  NewLoggerCfg(),
  }
}
func (l *Logger) SetEncoderCfg(cfg *encoder.KlogEncoder) *Logger {
  l.encoderCfg = cfg
  return l
}

func (l *Logger) SetLoggerCfg(cfg *LoggerCfg) *Logger {
  l.loggerCfg = cfg
  return l
}

func (l *Logger) EncoderCfg() *encoder.KlogEncoder {
  return l.encoderCfg
}

func (l *Logger) New() *Logger {
  return NewLogger(l.encoderCfg, l.loggerCfg)
}

func (l *Logger) Build() *Logger {
  return NewLogger(l.encoderCfg, l.loggerCfg)
}

func (l *Logger) LoggerCfg() *LoggerCfg {
  return l.loggerCfg
}

func (l *Logger) WithOutput(w zapcore.WriteSyncer, encoderCfg ...*encoder.KlogEncoder) *zap.Logger {
  bc := &core.BaseCore{}
  cfg := l.encoderCfg
  if len(encoderCfg) >= 1 {
    cfg = encoderCfg[0]
  }
  bc.SetEncoderCfg(cfg)
  bc.SetType(l.loggerCfg.Type)
  bc.SetLevel(TransportLevel(l.loggerCfg.Level))
  bc.SetOutput(w)
  b := bc.Build()
  wrapCore := zap.WrapCore(func(z zapcore.Core) zapcore.Core {
    return zapcore.NewTee(z, b)
  })
  return l.WithOptions(wrapCore)
}

func NewLogger(encoderCfg *encoder.KlogEncoder, cfg ...*LoggerCfg) *Logger {
  loggerCfg := NewLoggerCfg(cfg...)

  var cores []zapcore.Core
  if encoderCfg == nil {
    encoderCfg = encoder.NewKlogEncoderCfg()
  }

  if loggerCfg.StacktraceKey != "" {
    encoderCfg.StacktraceKey(loggerCfg.StacktraceKey)
  }

  if loggerCfg.EncodeLevel != "" {
    encoderCfg.EncodeLevel(ZapEncodeLevel(loggerCfg.EncodeLevel))
  }
  if loggerCfg.LogInConsole {
    cores = append(cores, CreateConsoleCoreWithLoggerCfg(encoderCfg, loggerCfg))
  }

  if loggerCfg.LogInFile {
    cores = append(cores, CreateFileCoreWithLoggerCfg(encoderCfg, loggerCfg))
  }

  if len(cores) == 0 {
    cores = append(cores, zapcore.NewNopCore())
  }
  logger := zap.New(zapcore.NewTee(cores...))
  if loggerCfg.ShowLine {
    logger.WithOptions(zap.AddCaller())
  }
  return &Logger{
    Logger:     logger,
    encoderCfg: encoderCfg,
    loggerCfg:  loggerCfg,
  }
}
func CreateFileCoreWithLoggerCfg(encoderCfg *encoder.KlogEncoder, cfg *LoggerCfg) zapcore.Core {
  fileCore := core.NewFileCore()
  fileCore.SetType(cfg.Type)
  fileCore.SetEncoderCfg(encoderCfg)
  fileCore.SetLevel(TransportLevel(cfg.Level))
  logger := inner.DefFileLogger()
  if cfg.MaxSize != 0 {
    logger.MaxSize = cfg.MaxSize
  }
  if cfg.MaxAge != 0 {
    logger.MaxAge = cfg.MaxAge
  }

  if cfg.MaxBackups != 0 {
    logger.MaxBackups = cfg.MaxBackups
  }
  if cfg.Compress {
    logger.Compress = cfg.Compress
  }
  if cfg.Filename != "" {
    logger.Filename = cfg.Filename
  }
  fileCore.SetOutput(writesyncer.FileWriteSyncer(logger))
  return fileCore.Build()
}
func CreateConsoleCoreWithLoggerCfg(encoderCfg *encoder.KlogEncoder, cfg *LoggerCfg) zapcore.Core {
  stdCore := core.NewStdCore()
  stdCore.SetType(cfg.Type)
  stdCore.SetEncoderCfg(encoderCfg)
  stdCore.SetLevel(TransportLevel(cfg.Level))
  return stdCore.Build()
}
func ZapEncodeLevel(levelEncoder string) zapcore.LevelEncoder {
  switch {
  case levelEncoder == "LowercaseLevelEncoder": // 小写编码器(默认)
    return zapcore.LowercaseLevelEncoder
  case levelEncoder == "LowercaseColorLevelEncoder": // 小写编码器带颜色
    return zapcore.LowercaseColorLevelEncoder
  case levelEncoder == "CapitalLevelEncoder": // 大写编码器
    return zapcore.CapitalLevelEncoder
  case levelEncoder == "CapitalColorLevelEncoder": // 大写编码器带颜色
    return zapcore.CapitalColorLevelEncoder
  default:
    return zapcore.LowercaseLevelEncoder
  }
}
func TransportLevel(l string) zapcore.Level {
  l = strings.ToLower(l)
  switch l {
  case "debug":
    return zapcore.DebugLevel
  case "info":
    return zapcore.InfoLevel
  case "warn":
    return zapcore.WarnLevel
  case "error":
    return zapcore.WarnLevel
  case "dpanic":
    return zapcore.DPanicLevel
  case "panic":
    return zapcore.PanicLevel
  case "fatal":
    return zapcore.FatalLevel
  default:
    return zapcore.DebugLevel
  }
}
