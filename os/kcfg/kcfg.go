package kcfg

import (
  "github.com/fsnotify/fsnotify"
  "github.com/spf13/viper"
)

type Kcfg struct {
  *viper.Viper
  Path       string                   `json:"path" yaml:"path"`
  MergeEnv   bool                     `json:"merge_env" yaml:"merge_env"`
  ConfigType string                   `json:"config_type" yaml:"config_type"`
  WatchMode  bool                     `json:"watch_mode" yaml:"watch_mode"`
  ChangeFn   func(isWatchChange bool) `json:"change_fn" yaml:"change_fn"`
}

func (k *Kcfg) Load() (*Kcfg, error) {
  k.AddConfigPath(".")
  k.SetConfigPath(k.Path)
  k.SetConfigType(k.ConfigType)
  if k.MergeEnv {
    k.Viper.AutomaticEnv()
    k.Viper.AllowEmptyEnv(true)
  }
  if k.WatchMode {
    k.Viper.WatchConfig()
  }
  err := k.Update()
  k.ChangeFn(false)
  return k, err
}

func New() *Kcfg {
  kcfg := &Kcfg{
    Viper:      viper.New(),
    MergeEnv:   false,
    ConfigType: "yaml",
    WatchMode:  true,
    Path:       "./config.yaml",
    ChangeFn:   func(isWatchChange bool) {},
  }
  return kcfg
}

func (k *Kcfg) SetConfigPath(path string) *Kcfg {
  k.Path = path
  k.Viper.SetConfigFile(path)
  return k
}

func (k *Kcfg) SetConfigType(t string) *Kcfg {
  if t == "" {
    t = "yaml"
  }
  k.ConfigType = t
  k.Viper.SetConfigType(t)
  return k
}

func (k *Kcfg) OnConfigChange(fn func(e fsnotify.Event)) *Kcfg {
  k.Viper.OnConfigChange(func(in fsnotify.Event) {
    fn(in)
    k.ChangeFn(true)
  })
  return k
}

func (k *Kcfg) Update() error {
  return k.Viper.ReadInConfig()
}

func (k *Kcfg) ReadAsMap() (error, map[string]interface{}) {
  res := make(map[string]interface{})
  err := k.Unmarshal(&res)
  if err != nil {
    return err, nil
  }
  return err, res
}
