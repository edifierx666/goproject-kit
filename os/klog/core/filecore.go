package core

import (
  "github.com/edifierx666/goproject-kit/os/klog/encoder"
  "github.com/edifierx666/goproject-kit/os/klog/writesyncer"
  "go.uber.org/zap/zapcore"
)

type FileCore struct {
  *BaseCore
}

func NewFileCore(bc ...*BaseCore) *FileCore {
  b := &BaseCore{
    encoderCfg: encoder.NewKlogEncoderCfg(),
    level:      zapcore.DebugLevel,
    output:     writesyncer.FileWriteSyncer(),
  }
  if len(bc) >= 0 {
    b = bc[0]
  }

  return &FileCore{
    BaseCore: b,
  }
}
