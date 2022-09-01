package goproject_kit

import (
  "fmt"
  "os"
  "testing"

  "github.com/Netflix/go-env"
  "github.com/edifierx666/goproject-kit/os/kcfg"
  "github.com/edifierx666/goproject-kit/os/klog"
  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
)

func TestName(t *testing.T) {
  res := make(map[string]string)
  envSet, err := env.EnvironToEnvSet(os.Environ())
  s := "ddddddddddddddddddddddddddddddddddddd"
  envSet.Apply(env.ChangeSet{
    "AUTHORS": &s,
  })
  os.Setenv("NIUBI", "NIUBI11111111111111111111111111")
  environ := os.Environ()
  fmt.Println(err, envSet, res, environ)
}
func TestName1(t *testing.T) {
  cfg := klog.GetDefEncoderCfg()
  file, _ := os.OpenFile("./log.log", os.O_CREATE|os.O_RDWR, os.ModePerm)
  sync := zapcore.AddSync(file)
  core := zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), os.Stdout, zap.NewAtomicLevelAt(zap.DebugLevel))
  newCore := zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), sync, zap.NewAtomicLevelAt(zap.WarnLevel))

  tee := zapcore.NewTee(core, newCore)
  logger := zap.New(tee)

  logger.Info("?????", zap.String("asd", "addddddddddd"))
  logger.Named("???///").Info("aaaa")
  logger.Warn("warn")

}
func TestName2(t *testing.T) {
  k := &klog.FileCore{}
  fmt.Println(k)
}

func TestName3(t *testing.T) {
  logger := klog.New(&klog.EasyZapCfg{})
  logger.Info("???????")
}

func TestCfg1(t *testing.T) {
  _, m := kcfg.New().ReadAsMap()
  fmt.Println(m)
}
