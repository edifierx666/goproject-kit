package kmerge

import "github.com/imdario/mergo"

func Merge(dst, src interface{}, opts ...func(*mergo.Config)) error {
  return mergo.Merge(dst, src, opts...)
}

func Map(dst, src interface{}, opts ...func(*mergo.Config)) error {
  return mergo.Map(dst, src, opts...)
}
