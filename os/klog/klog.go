package klog

import (
  "strings"

  "github.com/edifierx666/goproject-kit/os/klog/core"
  "github.com/edifierx666/goproject-kit/os/klog/encoder"
  "github.com/edifierx666/goproject-kit/os/klog/writesyncer"
  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
  "gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
  zap zap.Logger
}

func NewLogger(cfg ...*LoggerCfg) *zap.Logger {
  loggerCfg := NewLoggerCfg(cfg...)

  var cores []zapcore.Core
  encoderCfg := encoder.NewKlogEncoderCfg()

  if loggerCfg.StacktraceKey != "" {
    encoderCfg.StacktraceKey(loggerCfg.StacktraceKey)
  }

  if loggerCfg.EncodeLevel != "" {
    encoderCfg.EncodeLevel(ZapEncodeLevel(loggerCfg.EncodeLevel))
  }
  if loggerCfg.LogInConsole {
    cores = append(cores, createConsoleCoreWithLoggerCfg(encoderCfg, loggerCfg))
  }

  if loggerCfg.LogInFile {
    cores = append(cores, createFileCoreWithLoggerCfg(encoderCfg, loggerCfg))
  }

  if len(cores) == 0 {
    cores = append(cores, zapcore.NewNopCore())
  }

  logger := zap.New(zapcore.NewTee(cores...))
  if loggerCfg.ShowLine {
    logger.WithOptions(zap.AddCaller())
  }
  return logger
}

func createFileCoreWithLoggerCfg(encoderCfg *encoder.KlogEncoder, cfg *LoggerCfg) zapcore.Core {
  fileCore := core.NewFileCore()
  fileCore.SetType(cfg.Type)
  fileCore.SetEncoderCfg(encoderCfg)
  fileCore.SetLevel(TransportLevel(cfg.Level))
  logger := lumberjack.Logger{
    Filename:   "./log",
    MaxSize:    100,
    MaxAge:     100,
    MaxBackups: 3,
    Compress:   false,
  }
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
  fileCore.SetOutput(writesyncer.FileWriteSyncer(&logger))
  return fileCore.Build()
}
func createConsoleCoreWithLoggerCfg(encoderCfg *encoder.KlogEncoder, cfg *LoggerCfg) zapcore.Core {
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

// TransportLevel 根据字符串转化为 zapcore.Level
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
