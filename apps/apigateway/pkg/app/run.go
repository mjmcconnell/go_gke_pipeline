package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	middleware "go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"

	"github.com/mjmcconnell/go_gke_pipeline/apps/apigateway/pkg/config"
	"github.com/mjmcconnell/go_gke_pipeline/apps/apigateway/pkg/endpoints"
	"github.com/mjmcconnell/go_gke_pipeline/apps/apigateway/pkg/messaging"
	"github.com/mjmcconnell/go_gke_pipeline/apps/apigateway/pkg/monitoring"
)

func Run() error {
	ctx := context.Background()

	settings := config.Settings{}.New()

	logger := monitoring.GetLogger()
	tracingCleanup := monitoring.InitTracer(settings)
	defer tracingCleanup()

	srvError := make(chan error)
	msgError := make(chan error)
	// Start the private server
	go func() { srvError <- startPrivateServer(settings) }()
	// Start the public server
	go func() { srvError <- startPublicServer(settings) }()
	// Start message listener
	go func() { msgError <- startMessageListener(ctx, settings) }()

	// Check if an error has occured during a servers startup
	// Else wait until the system sends an interrupt to the application
	select {
	case err := <-srvError:
		logger.Error(err)
		return err
	case err := <-msgError:
		logger.Error(err)
		return err
	case <-sysInterrupt():
	}

	return nil
}

func startPublicServer(settings config.Settings) error {
	router := mux.NewRouter()
	router.Use(monitoring.LoggingMiddleware)
	router.Use(middleware.Middleware(settings.AppName))

	endpoints.MainHandler{Settings: settings}.Register(router)
	err := http.ListenAndServe(":8080", router)
	return err
}

func startPrivateServer(settings config.Settings) error {
	router := mux.NewRouter()
	endpoints.MetaHandler{}.Register(router)
	err := http.ListenAndServe(":8888", router)
	return err
}

func startMessageListener(ctx context.Context, settings config.Settings) error {
	subscription, err := messaging.Sub{
		ProjectID:      settings.MessagingProjectID,
		SubscriptionID: settings.MessagingSubscriptionID,
	}.New()

	if err != nil {
		return err
	}

	err = subscription.Listen(ctx)

	return err
}

func sysInterrupt() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	return c
}
