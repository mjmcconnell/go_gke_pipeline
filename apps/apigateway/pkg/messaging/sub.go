package messaging

import (
	"context"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/sirupsen/logrus"

	"github.com/mjmcconnell/go_gke_pipeline/apps/apigateway/pkg/monitoring"
)

type Sub struct {
	ProjectID      string
	SubscriptionID string
	client         *pubsub.Client
	sub            *pubsub.Subscription
}

func (s Sub) New(ctx context.Context) (Sub, error) {
	client, err := pubsub.NewClient(ctx, s.ProjectID)

	if err != nil {
		return Sub{}, err
	}

	s.client = client
	s.sub = s.client.Subscription(s.SubscriptionID)

	return s, nil
}

func (s Sub) Listen(ctx context.Context) error {
	var mu sync.Mutex
	err := s.sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()

		s.handleMessage(msg)
	})

	return err
}

func (s Sub) handleMessage(msg *pubsub.Message) {
	logger := monitoring.GetLogger()
	logger.WithFields(logrus.Fields{
		"ID":          msg.ID,
		"payload":     string(msg.Data),
		"attributes":  msg.Attributes,
		"publishedAt": msg.PublishTime,
		"attempt":     msg.DeliveryAttempt,
	}).Info("Recieved message, for subscription", s.SubscriptionID)
}
