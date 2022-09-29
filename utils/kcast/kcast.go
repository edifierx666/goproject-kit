package kcast

import (
  "time"

  "github.com/spf13/cast"
)

// ToBool casts an interface to a bool type.
func ToBool(i interface{}, fallbackVal ...bool) bool {
  v, e := cast.ToBoolE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToTime casts an interface to a time.Time type.
func ToTime(i interface{}, fallbackVal ...time.Time) time.Time {
  v, e := cast.ToTimeE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

func ToTimeInDefaultLocation(i interface{}, location *time.Location, fallbackVal ...time.Time) time.Time {
  v, e := cast.ToTimeInDefaultLocationE(i, location)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToDuration casts an interface to a time.Duration type.
func ToDuration(i interface{}, fallbackVal ...time.Duration) time.Duration {
  v, e := cast.ToDurationE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToFloat64 casts an interface to a float64 type.
func ToFloat64(i interface{}, fallbackVal ...float64) float64 {
  v, e := cast.ToFloat64E(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToFloat32 casts an interface to a float32 type.
func ToFloat32(i interface{}, fallbackVal ...float32) float32 {
  v, e := cast.ToFloat32E(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToInt64 casts an interface to an int64 type.
func ToInt64(i interface{}, fallbackVal ...int64) int64 {
  v, e := cast.ToInt64E(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToInt32 casts an interface to an int32 type.
func ToInt32(i interface{}, fallbackVal ...int32) int32 {
  v, e := cast.ToInt32E(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToInt16 casts an interface to an int16 type.
func ToInt16(i interface{}, fallbackVal ...int16) int16 {
  v, e := cast.ToInt16E(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToInt8 casts an interface to an int8 type.
func ToInt8(i interface{}, fallbackVal ...int8) int8 {
  v, e := cast.ToInt8E(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToInt casts an interface to an int type.
func ToInt(i interface{}, fallbackVal ...int) int {
  v, e := cast.ToIntE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToUint casts an interface to a uint type.
func ToUint(i interface{}, fallbackVal ...uint) uint {
  v, e := cast.ToUintE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToUint64 casts an interface to a uint64 type.
func ToUint64(i interface{}, fallbackVal ...uint64) uint64 {
  v, e := cast.ToUint64E(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToUint32 casts an interface to a uint32 type.
func ToUint32(i interface{}, fallbackVal ...uint32) uint32 {
  v, e := cast.ToUint32E(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToUint16 casts an interface to a uint16 type.
func ToUint16(i interface{}, fallbackVal ...uint16) uint16 {
  v, e := cast.ToUint16E(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToUint8 casts an interface to a uint8 type.
func ToUint8(i interface{}, fallbackVal ...uint8) uint8 {
  v, e := cast.ToUint8E(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToString casts an interface to a string type.
func ToString(i interface{}, fallbackVal ...string) string {
  v, e := cast.ToStringE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToStringMapString casts an interface to a map[string]string type.
func ToStringMapString(i interface{}, fallbackVal ...map[string]string) map[string]string {
  v, e := cast.ToStringMapStringE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToStringMapStringSlice casts an interface to a map[string][]string type.
func ToStringMapStringSlice(i interface{}, fallbackVal ...map[string][]string) map[string][]string {
  v, e := cast.ToStringMapStringSliceE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToStringMapBool casts an interface to a map[string]bool type.
func ToStringMapBool(i interface{}, fallbackVal ...map[string]bool) map[string]bool {
  v, e := cast.ToStringMapBoolE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToStringMapInt casts an interface to a map[string]int type.
func ToStringMapInt(i interface{}, fallbackVal ...map[string]int) map[string]int {
  v, e := cast.ToStringMapIntE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToStringMapInt64 casts an interface to a map[string]int64 type.
func ToStringMapInt64(i interface{}, fallbackVal ...map[string]int64) map[string]int64 {
  v, e := cast.ToStringMapInt64E(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToStringMap casts an interface to a map[string]interface{} type.
func ToStringMap(i interface{}, fallbackVal ...map[string]interface{}) map[string]interface{} {
  v, e := cast.ToStringMapE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToSlice casts an interface to a []interface{} type.
func ToSlice(i interface{}, fallbackVal ...[]interface{}) []interface{} {
  v, e := cast.ToSliceE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToBoolSlice casts an interface to a []bool type.
func ToBoolSlice(i interface{}, fallbackVal ...[]bool) []bool {
  v, e := cast.ToBoolSliceE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToStringSlice casts an interface to a []string type.
func ToStringSlice(i interface{}, fallbackVal ...[]string) []string {
  v, e := cast.ToStringSliceE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToIntSlice casts an interface to a []int type.
func ToIntSlice(i interface{}, fallbackVal ...[]int) []int {
  v, e := cast.ToIntSliceE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}

// ToDurationSlice casts an interface to a []time.Duration type.
func ToDurationSlice(i interface{}, fallbackVal ...[]time.Duration) []time.Duration {
  v, e := cast.ToDurationSliceE(i)
  if e != nil && len(fallbackVal) > 0 {
    return fallbackVal[0]
  }
  return v
}
