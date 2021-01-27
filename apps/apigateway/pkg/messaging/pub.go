package messaging

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type Topic struct {
	ProjectID string
	TopicID   string
	client    *pubsub.Client
	topic     *pubsub.Topic
}

func (t Topic) New(ctx context.Context) (Topic, error) {
	client, err := pubsub.NewClient(ctx, t.ProjectID)

	if err != nil {
		return Topic{}, err
	}

	t.client = client
	t.topic = t.client.Topic(t.TopicID)

	return t, nil
}

func (t Topic) Publish(ctx context.Context, msg string) (string, error) {
	result := t.topic.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	return id, err
}
