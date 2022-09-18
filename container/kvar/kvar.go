package kvar

import "github.com/gogf/gf/v2/container/gvar"

func New(val interface{}, safe bool) *gvar.Var {
  return gvar.New(val, safe)
}
