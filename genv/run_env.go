package genv

import "os"

func SetRunEnv(val string) {
	os.Setenv("RUN_ENV", val)
}
func GetRunEnv() string {
	return os.Getenv("RUN_ENV")
}
