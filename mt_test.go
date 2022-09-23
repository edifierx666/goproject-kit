package goproject_kit

import (
  "fmt"
  "os"
  "testing"

  "github.com/Netflix/go-env"
  "github.com/edifierx666/goproject-kit/os/kcfg"
  "github.com/edifierx666/goproject-kit/utils/kcast"
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
  toBool := kcast.ToBool("true", false)
  fmt.Println(toBool)
}
func TestName2(t *testing.T) {
}

func TestName3(t *testing.T) {

}

func TestCfg1(t *testing.T) {
  _, m := kcfg.New().ReadAsMap()
  fmt.Println(m)
}
