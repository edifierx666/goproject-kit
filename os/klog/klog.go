package klog

import "go.uber.org/zap"

func New(cfgs ...*EasyZapCfg) *zap.Logger {
  return Zap(cfgs...)
}
