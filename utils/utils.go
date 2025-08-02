package utils

import (
	"encoding/base64"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
)

var rootPath string

func init() {
	rootPath = findProjectRoot()
}

func GetRootPath() string {
	return rootPath
}

func findProjectRoot() string {
	if isDevelopment() {
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			dir := filepath.Dir(filename)
			for {
				goModPath := filepath.Join(dir, "go.mod")
				if _, err := os.Stat(goModPath); err == nil {
					return dir
				}
				parent := filepath.Dir(dir)
				if parent == dir {
					break
				}
				dir = parent
			}
		}
	}
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	exeDir := filepath.Dir(exePath)
	return exeDir
}

func isDevelopment() bool {
	_, err := os.Stat("go.mod")
	return err == nil
}

func GenerateClientTransactionID() string {
	randomBytes := make([]byte, 48)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	id := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(randomBytes)
	return id
}
