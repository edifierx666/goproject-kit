package encoder

import (
  "fmt"
  "io"
  "time"

  "go.uber.org/zap/zapcore"
)

type KlogEncoder struct {
  inner zapcore.EncoderConfig
}

func (e *KlogEncoder) Inner() zapcore.EncoderConfig {
  return e.inner
}

func NewKlogEncoderCfg(encoders ...zapcore.EncoderConfig) *KlogEncoder {
  var inner zapcore.EncoderConfig
  if len(encoders) >= 1 {
    inner = encoders[0]
  } else {
    inner = DefEncoderCfg()
  }

  return &KlogEncoder{inner: inner}
}

func (e *KlogEncoder) MessageKey(messageKey string) *KlogEncoder {
  e.inner.MessageKey = messageKey
  return e
}

func (e *KlogEncoder) LevelKey(levelKey string) *KlogEncoder {
  e.inner.LevelKey = levelKey
  return e
}

func (e *KlogEncoder) TimeKey(timeKey string) *KlogEncoder {
  e.inner.TimeKey = timeKey
  return e
}

func (e *KlogEncoder) NameKey(nameKey string) *KlogEncoder {
  e.inner.NameKey = nameKey
  return e
}

func (e *KlogEncoder) CallerKey(callerKey string) *KlogEncoder {
  e.inner.CallerKey = callerKey
  return e
}

func (e *KlogEncoder) FunctionKey(functionKey string) *KlogEncoder {
  e.inner.FunctionKey = functionKey
  return e
}

func (e *KlogEncoder) StacktraceKey(stacktraceKey string) *KlogEncoder {
  e.inner.StacktraceKey = stacktraceKey
  return e
}

func (e *KlogEncoder) SkipLineEnding(skipLineEnding bool) *KlogEncoder {
  e.inner.SkipLineEnding = skipLineEnding
  return e
}

func (e *KlogEncoder) LineEnding(lineEnding string) *KlogEncoder {
  e.inner.LineEnding = lineEnding
  return e
}

func (e *KlogEncoder) EncodeLevel(encodeLevel zapcore.LevelEncoder) *KlogEncoder {
  e.inner.EncodeLevel = encodeLevel
  return e
}

func (e *KlogEncoder) EncodeTime(encodeTime zapcore.TimeEncoder) *KlogEncoder {
  e.inner.EncodeTime = encodeTime
  return e
}

func (e *KlogEncoder) EncodeDuration(encodeDuration zapcore.DurationEncoder) *KlogEncoder {
  e.inner.EncodeDuration = encodeDuration
  return e
}

func (e *KlogEncoder) EncodeCaller(encodeCaller zapcore.CallerEncoder) *KlogEncoder {
  e.inner.EncodeCaller = encodeCaller
  return e
}

func (e *KlogEncoder) EncodeName(encodeName zapcore.NameEncoder) *KlogEncoder {
  e.inner.EncodeName = encodeName
  return e
}

func (e *KlogEncoder) NewReflectedEncoder(newReflectedEncoder func(writer io.Writer) zapcore.ReflectedEncoder) *KlogEncoder {
  e.inner.NewReflectedEncoder = newReflectedEncoder
  return e
}

func (e *KlogEncoder) ConsoleSeparator(consoleSeparator string) *KlogEncoder {
  e.inner.ConsoleSeparator = consoleSeparator
  return e
}

func (e *KlogEncoder) Build() (*KlogEncoder, zapcore.EncoderConfig) {
  return e, e.inner
}

func (e *KlogEncoder) BuildConsoleEncoder() zapcore.Encoder {
  return zapcore.NewConsoleEncoder(e.inner)
}

func (e *KlogEncoder) BuildJsonEncoder() zapcore.Encoder {
  return zapcore.NewJSONEncoder(e.inner)
}

func DefEncoderCfg() zapcore.EncoderConfig {
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
      DefaultTimeEncoder(t, encoder, "")
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

func DefaultTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder, prefix string) {
  encoder.AppendString(prefix + t.Format("2006-01-02 15:04:05.000"))
}
