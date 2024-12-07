package chime

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings"
)

type ChimeClient struct {
	client *chimesdkmeetings.Client
}

func NewChimeClient() (*ChimeClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	svc := chimesdkmeetings.NewFromConfig(cfg)

	return &ChimeClient{
		client: svc,
	}, nil
}

func (c *ChimeClient) CreateMeeting(ctx context.Context, input *chimesdkmeetings.CreateMeetingInput) (*chimesdkmeetings.CreateMeetingOutput, error) {
	result, err := c.client.CreateMeeting(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to create meeting: %w", err)
	}

	return result, nil
}

func (c *ChimeClient) CreateAttendee(ctx context.Context, input *chimesdkmeetings.CreateAttendeeInput) (*chimesdkmeetings.CreateAttendeeOutput, error) {
	result, err := c.client.CreateAttendee(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to create attendee: %w", err)
	}

	return result, nil
}

func (c *ChimeClient) DeleteMeeting(ctx context.Context, input *chimesdkmeetings.DeleteMeetingInput) error {
	_, err := c.client.DeleteMeeting(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete meeting: %w", err)
	}

	return nil
}

func (c *ChimeClient) DeleteAttendee(ctx context.Context, input *chimesdkmeetings.DeleteAttendeeInput) error {
	_, err := c.client.DeleteAttendee(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete attendee: %w", err)
	}

	return nil
}
