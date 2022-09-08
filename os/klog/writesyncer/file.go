package writesyncer

import (
  "github.com/edifierx666/goproject-kit/os/klog/inner"
  "go.uber.org/zap/zapcore"
  "gopkg.in/natefinch/lumberjack.v2"
)

func FileWriteSyncer(fileLogger ...*lumberjack.Logger) zapcore.WriteSyncer {
  var filelog *lumberjack.Logger
  if len(fileLogger) == 0 {
    filelog = inner.NewFileLogger().Build()
  } else {
    filelog = fileLogger[0]
  }
  return zapcore.AddSync(filelog)
}
