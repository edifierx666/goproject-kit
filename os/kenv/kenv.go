package kenv

import (
  "log"
  "os"

  "github.com/Netflix/go-env"
)

var EnvSet env.EnvSet
var OsEnv []string

func init() {
  Update()
}

// Update 在修改了env后调用会更新EnvSet和OsEnv 比如os.setenv 或一些外部设置后
func Update() {
  OsEnv = os.Environ()
  environ, err := env.UnmarshalFromEnviron(OsEnv)
  if err != nil {
    log.Printf("kenv.UnmarshalFromEnviron err %v\r\n", err)
  }
  EnvSet = environ
}

func Has(key string) bool {
  _, ok := EnvSet[key]
  return ok
}

func Get(key string) string {
  v, ok := EnvSet[key]
  if !ok {
    return ""
  }
  return v
}
