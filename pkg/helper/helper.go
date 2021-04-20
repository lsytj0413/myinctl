package helper

import (
	"path/filepath"
	"reflect"
	"runtime"
)

// currentProjectRootPath is the root path of current project src
var currentProjectRootPath string

// JoinWithProjectAbsPath return the path's full-path under project root
func JoinWithProjectAbsPath(path string) string {
	return filepath.Join(currentProjectRootPath, path)
}

func currentFilePath() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

func init() {
	path := currentFilePath()
	var err error
	currentProjectRootPath, err = filepath.Abs(filepath.Join(filepath.Dir(path), "../.."))
	if err != nil {
		panic(err)
	}
}

// DataSize returns the number of bytes the actual data represented by v occupies in memory.
// For compound structures, it sums the sizes of the elements.
// For instance or slice, it returns the length of the slice times the element size and doesn't count the memory occupied by the header.
// If the type of v is not acceptable, returns -1.
func DataSize(data interface{}) int {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Ptr:
		return dataSize(v.Elem())
	case reflect.Slice:
		return dataSize(v)
	}

	return -1
}

func dataSize(v reflect.Value) int {
	if v.Kind() == reflect.Slice {
		if s := sizeof(v.Type().Elem()); s >= 0 {
			return s * v.Len()
		}

		return -1
	}

	return sizeof(v.Type())
}

func sizeof(t reflect.Type) int {
	switch t.Kind() {
	case reflect.Array:
		if s := sizeof(t.Elem()); s >= 0 {
			return s * t.Len()
		}

	case reflect.Struct:
		sum := 0
		for i, n := 0, t.NumField(); i < n; i++ {
			s := sizeof(t.Field(i).Type)
			if s < 0 {
				return -1
			}
			sum += s
		}
		return sum

	case reflect.Bool,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return int(t.Size())
	}

	return -1
}
