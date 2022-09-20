package kvar

import (
  "github.com/gogf/gf/v2/container/gvar"
)

type Var struct {
  *gvar.Var
}

func New(val interface{}, safe ...bool) *Var {
  if len(safe) > 0 && !safe[0] {
    return &Var{
      Var: gvar.New(val, true),
    }
  }
  return &Var{
    Var: gvar.New(val),
  }
}
