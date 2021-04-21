package helper

import (
	"path/filepath"
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
