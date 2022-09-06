package klog

import (
  "fmt"
  "time"

  "go.uber.org/zap/zapcore"
)

func GetDefEncoderCfg() zapcore.EncoderConfig {
  return zapcore.EncoderConfig{
    MessageKey:     "message",
    LevelKey:       "level",
    TimeKey:        "timeKey",
    NameKey:        "logger",
    CallerKey:      "caller",
    FunctionKey:    "function",
    StacktraceKey:  "stacktrace",
    SkipLineEnding: false,
    LineEnding:     zapcore.DefaultLineEnding,
    // debug info 等级别的输出格式
    EncodeLevel: func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
      encoder.AppendString(fmt.Sprintf("[%v]", l.CapitalString()))
    },
    // time 的输出格式
    EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
      CustomTimeEncoder(t, encoder, "")
    },
    EncodeDuration: zapcore.NanosDurationEncoder,
    EncodeCaller:   zapcore.FullCallerEncoder,
    // logger namspace格式
    EncodeName: func(s string, encoder zapcore.PrimitiveArrayEncoder) {
      encoder.AppendString(fmt.Sprintf("[%s]", s))
    },
    // 每个属性的分隔符
    ConsoleSeparator: " ",
  }
}

func GetConsoleEncoder(config ...zapcore.EncoderConfig) zapcore.Encoder {
  var cfg zapcore.EncoderConfig
  if len(config) == 0 {
    cfg = GetDefEncoderCfg()
  } else {
    cfg = config[0]
  }
  return zapcore.NewConsoleEncoder(cfg)
}

func GetJsonEncoder(config ...zapcore.EncoderConfig) zapcore.Encoder {
  var cfg zapcore.EncoderConfig
  if len(config) == 0 {
    cfg = GetDefEncoderCfg()
  } else {
    cfg = config[0]
  }
  return zapcore.NewJSONEncoder(cfg)
}

func CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder, prefix string) {
  encoder.AppendString(prefix + t.Format("2006/01/02 15:04:05.000"))
}
