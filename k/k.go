package k

import (
  "context"
  "io"

  "github.com/edifierx666/goproject-kit/container/kvar"
  "github.com/edifierx666/goproject-kit/internal/empty"
  "github.com/gogf/gf/v2/net/ghttp"
  "github.com/gogf/gf/v2/util/gutil"
)

// NewVar returns a gvar.Var.
func NewVar(i interface{}, safe ...bool) *kvar.Var {
  return kvar.New(i, safe...)
}

// Dump dumps a variable to stdout with more manually readable.
func Dump(values ...interface{}) {
  gutil.Dump(values...)
}

// DumpTo writes variables `values` as a string in to `writer` with more manually readable
func DumpTo(writer io.Writer, value interface{}, option gutil.DumpOption) {
  gutil.DumpTo(writer, value, option)
}

// DumpWithType acts like Dump, but with type information.
// Also see Dump.
func DumpWithType(values ...interface{}) {
  gutil.DumpWithType(values...)
}

// DumpWithOption returns variables `values` as a string with more manually readable.
func DumpWithOption(value interface{}, option gutil.DumpOption) {
  gutil.DumpWithOption(value, option)
}

// Throw throws an exception, which can be caught by TryCatch function.
func Throw(exception interface{}) {
  gutil.Throw(exception)
}

// Try implements try... logistics using internal panic...recover.
// It returns error if any exception occurs, or else it returns nil.
func Try(ctx context.Context, try func(ctx context.Context)) (err error) {
  return gutil.Try(ctx, try)
}

// TryCatch implements try...catch... logistics using internal panic...recover.
// It automatically calls function `catch` if any exception occurs ans passes the exception as an error.
func TryCatch(ctx context.Context, try func(ctx context.Context), catch ...func(ctx context.Context, exception error)) {
  gutil.TryCatch(ctx, try, catch...)
}

// IsNil checks whether given `value` is nil.
// Parameter `traceSource` is used for tracing to the source variable if given `value` is type
// of pinter that also points to a pointer. It returns nil if the source is nil when `traceSource`
// is true.
// Note that it might use reflect feature which affects performance a little.
func IsNil(value interface{}, traceSource ...bool) bool {
  return empty.IsNil(value, traceSource...)
}

// IsEmpty checks whether given `value` empty.
// It returns true if `value` is in: 0, nil, false, "", len(slice/map/chan) == 0.
// Or else it returns true.
func IsEmpty(value interface{}) bool {
  return empty.IsEmpty(value)
}

// RequestFromCtx retrieves and returns the Request object from context.
func RequestFromCtx(ctx context.Context) *ghttp.Request {
  return ghttp.RequestFromCtx(ctx)
}
