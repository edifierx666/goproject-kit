package klog

type LoggerCfg struct {
  LogInFile     bool   `json:"log_in_file" yaml:"log_in_file"`
  LogInConsole  bool   `json:"log_in_console" yaml:"log_in_console"`
  Filename      string `json:"filename" yaml:"filename"`
  MaxSize       int    `json:"max_size" yaml:"max_size"`
  MaxAge        int    `json:"max_age" yaml:"max_age"`
  MaxBackups    int    `json:"max_backups" yaml:"max_backups"`
  Compress      bool   `json:"compress" yaml:"compress"`
  ShowLine      bool   `json:"show_line" yaml:"show_line"`
  StacktraceKey string `json:"stacktrace_key" yaml:"stacktrace_key"`
  EncodeLevel   string `json:"encode_level" yaml:"encode_level"`
  Level         string `json:"level" yaml:"level"`
  Type          string `json:"type" yaml:"type"`
}

func NewLoggerCfg(cfg ...*LoggerCfg) *LoggerCfg {

  l := &LoggerCfg{
    LogInFile:     false,
    Filename:      "",
    MaxSize:       100,
    MaxAge:        100,
    MaxBackups:    100,
    Compress:      false,
    ShowLine:      false,
    LogInConsole:  true,
    StacktraceKey: "tracekey",
    EncodeLevel:   "CapitalLevelEncoder",
    Level:         "debug",
    Type:          "console",
  }

  if len(cfg) >= 1 {
    l = cfg[0]
  }

  return l
}

func (b *LoggerCfg) SetFilename(filename string) *LoggerCfg {
  b.Filename = filename
  return b
}

func (b *LoggerCfg) SetMaxSize(maxSize int) *LoggerCfg {
  b.MaxSize = maxSize
  return b
}

func (b *LoggerCfg) SetMaxAge(maxAge int) *LoggerCfg {
  b.MaxAge = maxAge
  return b
}

func (b *LoggerCfg) SetMaxBackups(maxBackups int) *LoggerCfg {
  b.MaxBackups = maxBackups
  return b
}

func (b *LoggerCfg) SetCompress(compress bool) *LoggerCfg {
  b.Compress = compress
  return b
}

func (b *LoggerCfg) SetShowLine(showLine bool) *LoggerCfg {
  b.ShowLine = showLine
  return b
}

func (b *LoggerCfg) SetLogInConsole(t bool) *LoggerCfg {
  b.LogInConsole = t
  return b
}

func (b *LoggerCfg) SetLogInFile(t bool) *LoggerCfg {
  b.LogInFile = t
  return b
}

func (b *LoggerCfg) SetStacktraceKey(stacktraceKey string) *LoggerCfg {
  b.StacktraceKey = stacktraceKey
  return b
}

func (b *LoggerCfg) SetEncodeLevel(encodeLevel string) *LoggerCfg {
  b.EncodeLevel = encodeLevel
  return b
}

func (b *LoggerCfg) SetLevel(level string) *LoggerCfg {
  b.Level = level
  return b
}
