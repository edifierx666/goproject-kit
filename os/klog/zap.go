package klog

import (
  "os"
  "strings"
  "time"

  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
)

type EasyZapCfg struct {
  EnableFile    bool   `json:"enable_file" yaml:"enable_file"`
  Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
  Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
  Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // 输出
  Director      string `mapstructure:"director" json:"director"  yaml:"director"`                  // 日志文件夹
  EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级
  StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名
  MaxAge        int    `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                      // 日志留存时间
  ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
  LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
  MaxSize       int    `json:"maxsize" yaml:"maxsize"`
}

func Zap(cfgs ...*EasyZapCfg) *zap.Logger {
  config := &EasyZapCfg{
    EnableFile:    false,
    Level:         "debug",
    Prefix:        "",
    Format:        "console",
    Director:      "./log/log",
    EncodeLevel:   "CapitalLevelEncoder",
    StacktraceKey: "",
    MaxAge:        0,
    ShowLine:      false,
    LogInConsole:  true,
  }

  if len(cfgs) > 0 {
    config = cfgs[0]
  }

  config.LogInConsole = true

  var syncers []zapcore.WriteSyncer

  if config.EnableFile {
    logger := GetFileLogger()

    if config.Director != "" {
      logger.Filename = config.Director
    }

    if config.MaxAge != 0 {
      logger.MaxAge = config.MaxAge
    }

    if config.MaxSize != 0 {
      logger.MaxSize = config.MaxSize
    }

    sync := zapcore.AddSync(GetZapFileWriteSyncer(logger))
    syncers = append(syncers, sync)
  }

  if config.LogInConsole {
    stdoutSync := zapcore.AddSync(os.Stdout)
    syncers = append(syncers, stdoutSync)
  }

  zapEnLevel := zapcore.CapitalLevelEncoder
  if config.EncodeLevel != "" {
    zapEnLevel = ZapEncodeLevel(config.EncodeLevel)
  }

  zapLevel := zapcore.DebugLevel
  if config.Level != "" {
    zapLevel = TransportLevel(config.Level)
  }

  cfg := GetDefEncoderCfg()
  cfg.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
    zapEnLevel(l, encoder)
  }

  if config.Prefix != "" {
    cfg.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
      CustomTimeEncoder(t, encoder, config.Prefix)
    }
  }

  if config.StacktraceKey != "" {
    cfg.StacktraceKey = config.StacktraceKey
  }

  encoder := zapcore.NewConsoleEncoder(cfg)
  if config.Format == "json" {
    encoder = zapcore.NewJSONEncoder(cfg)
  }

  normalCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(syncers...), zapLevel)
  logger := zap.New(normalCore)
  if config.ShowLine {
    logger = logger.WithOptions(zap.AddCaller())
  }

  return logger
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
