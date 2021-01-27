package config

import (
	"os"
)

type Settings struct {
	AppName                 string
	TracingHost             string
	TracingPort             string
	MessagingProjectID      string
	MessagingTopicID        string
	MessagingSubscriptionID string
}

func (s Settings) New() Settings {
	s.AppName = GetWithDefault("APP_NAME", "api-gateway-foobar")
	s.TracingHost = GetWithDefault("TRACING_HOST", "jaeger")
	s.TracingPort = GetWithDefault("TRACING_PORT", "14268")
	s.MessagingProjectID = Get("PUBSUB_PROJECT_ID")
	s.MessagingTopicID = Get("PUBSUB_REQUEST_TOPIC")
	s.MessagingSubscriptionID = Get("PUBSUB_RESPONSE_SUBSCRIPTION")

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
