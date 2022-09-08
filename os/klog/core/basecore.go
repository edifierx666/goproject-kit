package core

import (
  "github.com/edifierx666/goproject-kit/os/klog/encoder"
  "go.uber.org/zap/zapcore"
)

type BaseCore struct {
  encoderCfg *encoder.KlogEncoder
  level      zapcore.LevelEnabler
  output     zapcore.WriteSyncer
  oType      string
}

func (b *BaseCore) Build() zapcore.Core {
  inner := b.encoderCfg.Inner()
  e := zapcore.NewConsoleEncoder(inner)
  if b.oType == "json" {
    e = zapcore.NewJSONEncoder(inner)
  }
  return zapcore.NewCore(e, b.output, b.level)
}

func (b *BaseCore) SetEncoderCfg(encoder *encoder.KlogEncoder) *BaseCore {
  b.encoderCfg = encoder
  return b
}

func (b *BaseCore) SetOutput(ws zapcore.WriteSyncer) *BaseCore {
  b.output = ws
  return b
}

func (b *BaseCore) SetLevel(level zapcore.LevelEnabler) *BaseCore {
  b.level = level
  return b
}

func (b *BaseCore) SetType(t string) *BaseCore {
  b.oType = t
  return b
}

func (b *BaseCore) GetLevel() zapcore.LevelEnabler {
  return b.level
}

func (b *BaseCore) GetEncoderCfg() *encoder.KlogEncoder {
  return b.encoderCfg
}

func (b *BaseCore) GetOutput() zapcore.WriteSyncer {
  return b.output
}

func (b *BaseCore) GetOType() string {
  return b.oType
}
