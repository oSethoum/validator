package validator

import (
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unsafe"
)

func parseConstraints(tag string) []Constraint {
	cs := []Constraint{}
	tgs := strings.Split(tag, ";")
	for _, c := range tgs {
		if strings.Contains(c, "=") {
			if s := strings.Split(c, "="); len(s) > 0 {
				c := Constraint{
					Tag:  c,
					Kind: s[0],
				}
				if len(s) > 1 {
					c.Param = s[1]
				}
				cs = append(cs, c)
			}
		} else {
			cs = append(cs, Constraint{
				Tag:  c,
				Kind: c,
			})
		}
	}
	return cs
}

func isString(t reflect.Type) bool {
	return t.String() == typeString || t.String() == typeStringPtr
}

func getStringValue(v reflect.Value) (string, bool) {
	if v.Type().String() == typeStringPtr {
		if v.IsNil() {
			return "", false
		}
		return v.Elem().String(), true
	}

	if v.Type().String() == typeString {
		return v.String(), true
	}
	return "", false
}

func getStringParam(param any) (string, bool) {
	v, ok := param.(string)
	if ok {
		return v, true
	}
	return "", false
}

func isInt(t reflect.Type) bool {
	return strings.HasPrefix(t.String(), typeInt) || strings.HasPrefix(t.String(), typeIntPtr)
}

func getIntValue(v reflect.Value) (int64, bool) {
	if strings.HasPrefix(v.Type().String(), typeIntPtr) {
		if v.IsNil() {
			return 0, false
		}
		return v.Elem().Int(), true
	}
	if strings.HasPrefix(v.Type().String(), typeInt) {
		return v.Int(), true
	}
	return 0, false
}

func getIntParam(param any) (int64, bool) {
	v, ok := param.(string)
	if ok {
		i, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return i, ok
		}
	}
	return 0, false
}

func isUint(t reflect.Type) bool {
	return strings.HasPrefix(t.String(), typeUint) || strings.HasPrefix(t.String(), typeUintPtr)
}

func getUintParam(param any) (uint64, bool) {
	v, ok := param.(string)
	if ok {
		i, err := strconv.ParseUint(v, 10, 64)
		if err == nil {
			return i, ok
		}
	}
	return 0, false
}

func getOneOfString(param any) []string {
	if p, ok := param.(string); ok {
		p = strings.ReplaceAll(p, " ", "")
		ss := strings.Split(p, ",")
		if len(ss) == 0 {
			return nil
		}
		return ss
	}
	return nil
}

func getUintValue(v reflect.Value) (uint64, bool) {
	if v.Type().String() == typeIntPtr {
		if !v.IsNil() {
			return 0, false
		}
		return v.Elem().Uint(), true

	}
	if v.Type().String() == typeInt {
		return v.Uint(), true
	}
	return 0, false
}

func isFloat(t reflect.Type) bool {
	return strings.HasPrefix(t.String(), typeFloat) || strings.HasPrefix(t.String(), typeFloatPtr)
}

func getFloatValue(v reflect.Value) (float64, bool) {
	if v.Type().String() == typeFloatPtr {
		if !v.IsNil() {
			return 0, false
		}
		return v.Elem().Float(), true
	}
	if v.Type().String() == typeInt {
		return v.Float(), true
	}
	return 0, false
}

func getFloatParam(param any) (float64, bool) {
	v, ok := param.(string)
	if ok {
		f, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return f, true
		}
	}
	return 0, false
}

func getStringListParam(param any) ([]string, bool) {
	s, ok := param.(string)
	if ok {
		ss := strings.Split(s, ",")
		return ss, true
	}
	return nil, false
}

func getIntListParam(param any) ([]int64, bool) {
	if s, ok := param.(string); ok {
		ss := strings.Split(s, ",")
		res := []int64{}
		for _, v := range ss {
			n, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, false
			}
			res = append(res, n)
		}
		if len(res) > 0 {
			return res, true
		}
	}
	return nil, false
}

func getUintListParam(param any) ([]uint64, bool) {
	if s, ok := param.(string); ok {
		ss := strings.Split(s, ",")
		res := []uint64{}
		for _, v := range ss {
			n, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return nil, false
			}
			res = append(res, n)
		}
		if len(res) > 0 {
			return res, true
		}
	}
	return nil, false
}

func getFloatListParam(param any) ([]float64, bool) {
	if s, ok := param.(string); ok {
		ss := strings.Split(s, ",")
		res := []float64{}

		for _, v := range ss {
			n, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, false
			}
			res = append(res, n)
		}
		if len(res) > 0 {
			return res, true
		}
	}
	return nil, false
}

func isStringArray(t reflect.Type) bool {
	return strings.HasPrefix(t.String(), typeStringArray)
}

func getStringArrayValue(v reflect.Value) ([]string, bool) {
	if !v.IsNil() {
		value, ok := v.Interface().([]string)
		return value, ok
	}
	return nil, false
}

func getIntArrayValue(v reflect.Value) ([]int64, bool) {
	if !v.IsNil() {
		i := v.Interface()
		values := *(*[]int64)(unsafe.Pointer(&i))
		return values, true
	}
	return nil, false
}

func getUintArrayValue(v reflect.Value) ([]uint64, bool) {
	if !v.IsNil() {
		i := v.Interface()
		values := *(*[]uint64)(unsafe.Pointer(&i))
		return values, true
	}
	return nil, false
}

func getFloatArrayValue(v reflect.Value) ([]float64, bool) {
	if !v.IsNil() {
		i := v.Interface()
		values := *(*[]float64)(unsafe.Pointer(&i))
		return values, true
	}
	return nil, false
}

func isIntArray(t reflect.Type) bool {
	return t.String() == typeIntArray
}

func isUintArray(t reflect.Type) bool {
	return t.String() == typeUintArray
}

func isFloatArray(t reflect.Type) bool {
	return t.String() == typeFloatArray
}

func camel(s string) string {
	switch s {
	case "":
		return s
	case "ID":
		return "id"
	default:
		return string(unicode.ToLower(rune(s[0]))) + s[1:]
	}
}

// Utils
func insArray[T comparable](array []T, values []T) bool {
	for _, value := range values {
		if !inArray(array, value) {
			return false
		}
	}
	return true
}

func outsArray[T comparable](array []T, values []T) bool {
	for _, value := range values {
		if inArray(array, value) {
			return false
		}
	}
	return true
}

func inArray[T comparable](array []T, value T) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}
