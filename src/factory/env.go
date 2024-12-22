package factory

import (
	"fmt"
	"os"
)

type environmentRuntime int

const (
	EnvLocal environmentRuntime = iota

	local = "local"
)

func CheckRuntimeEnvironment() environmentRuntime {
	environment := MustLookupEnv("RUNTIME_ENVIRONMENT")
	switch environment {
	case local:
		return EnvLocal
	default:
		panic(fmt.Sprintf("factory: unknown runtime environment: '%s'", environment))
	}
}

func MustLookupEnv(key string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		panic(fmt.Sprintf("factory: environment '%s' not present", key))
	}

	return value
}
