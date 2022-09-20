package kcfg

import (
  "github.com/fsnotify/fsnotify"
  "github.com/spf13/viper"
)

type Kcfg struct {
  viper      *viper.Viper
  Path       string `json:"path" yaml:"path"`
  MergeEnv   bool   `json:"merge_env" yaml:"merge_env"`
  ConfigType string `json:"config_type" yaml:"config_type"`
  WatchMode  bool   `json:"watch_mode" yaml:"watch_mode"`
}

func (k *Kcfg) Viper() *viper.Viper {
  return k.viper
}

func (k *Kcfg) Build() *Kcfg {
  k.SetConfigPath(k.Path)
  k.SetConfigType(k.ConfigType)
  if k.MergeEnv {
    k.viper.AutomaticEnv()
  }
  if k.WatchMode {
    k.viper.WatchConfig()
  }
  return k
}

func New() *Kcfg {
  kcfg := &Kcfg{
    viper:      viper.New(),
    MergeEnv:   false,
    ConfigType: "yaml",
    WatchMode:  true,
    Path:       "./config.yaml",
  }
  return kcfg
}

func (k *Kcfg) SetConfigPath(path string) *Kcfg {
  k.Path = path
  k.viper.SetConfigFile(path)
  return k
}

func (k *Kcfg) SetConfigType(t string) *Kcfg {
  if t == "" {
    t = "yaml"
  }
  k.ConfigType = t
  k.viper.SetConfigType(t)
  return k
}

func (k *Kcfg) OnConfigChange(fn func(e fsnotify.Event)) *Kcfg {
  k.viper.OnConfigChange(fn)
  return k
}

func (k *Kcfg) Update() *Kcfg {
  err := k.viper.ReadInConfig()
  if err != nil {
    panic(err)
  }
  return k
}

func (k *Kcfg) Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error {
  err := k.viper.Unmarshal(rawVal, opts...)
  return err
}

func (k *Kcfg) ReadAsMap() (error, map[string]interface{}) {
  res := make(map[string]interface{})
  err := k.Unmarshal(&res)
  if err != nil {
    return err, nil
  }
  return err, res
}
