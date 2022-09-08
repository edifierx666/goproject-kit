package core

import (
  "os"

  "github.com/edifierx666/goproject-kit/os/klog/encoder"
  "go.uber.org/zap/zapcore"
)

type StdCore struct {
  *BaseCore
}

func NewStdCore(bc ...*BaseCore) *StdCore {
  b := &BaseCore{
    encoderCfg: encoder.NewKlogEncoderCfg(),
    level:      zapcore.DebugLevel,
    output:     os.Stdout,
  }
  if len(bc) >= 1 {
    b = bc[0]
  }
  return &StdCore{
    BaseCore: b,
  }
}
