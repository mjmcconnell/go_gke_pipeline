package config

import (
	"os"
)

type Settings struct {
	AppName     string
	TracingHost string
	TracingPort string
}

func (s Settings) New() Settings {
	s.AppName = GetWithDefault("APP_NAME", "api-gateway-foobar")
	s.TracingHost = GetWithDefault("TRACING_HOST", "jaeger")
	s.TracingPort = GetWithDefault("TRACING_PORT", "14268")

	return s
}

func Get(k string) string {
	return os.Getenv(k)
}

func GetWithDefault(k string, d string) string {
	v, ok := os.LookupEnv(k)

	if !ok {
		v = d
	}

	return v
}
