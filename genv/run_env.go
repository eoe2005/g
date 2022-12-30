package genv

import (
	"os"
	"strings"
)

func SetRunEnv(val string) {
	os.Setenv("RUN_ENV", strings.ToUpper(val))
}
func GetRunEnv() string {
	return os.Getenv("RUN_ENV")
}
