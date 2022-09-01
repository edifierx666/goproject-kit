package klog

import "go.uber.org/zap"

func New(easyZaps ...*EasyZap) *zap.Logger {
  if len(easyZaps) > 0 {
    return Zap(easyZaps[0])
  }
  return Zap(NewEasyZap())
}
