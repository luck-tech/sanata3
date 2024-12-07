package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
)

const (
	chatTable = "chats"
)

type DynamoRepository struct {
	client *dynamodb.Client
}

func NewDynamoRepository(client *dynamodb.Client) *DynamoRepository {
	return &DynamoRepository{client: client}
}

var _ dai.Chat = (*DynamoRepository)(nil)

func (r *DynamoRepository) CreateChat(ctx context.Context, chat *entity.Chat) error {
	item, err := attributevalue.MarshalMap(chat)
	if err != nil {
		return err
	}

	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("chats"),
		Item:      item,
	})
	return err
}

func (r *DynamoRepository) GetChats(ctx context.Context, roomID string, lastTime time.Time) ([]entity.Chat, error) {
	chats := []entity.Chat{}
	input := &dynamodb.QueryInput{
		TableName: aws.String(chatTable),
		KeyConditions: map[string]types.Condition{
			"roomId": {
				ComparisonOperator: types.ComparisonOperatorEq,
				AttributeValueList: []types.AttributeValue{
					&types.AttributeValueMemberS{Value: roomID},
				},
			},
		},
	}

	paginator := dynamodb.NewQueryPaginator(r.client, input)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, item := range page.Items {
			var chat entity.Chat
			err := attributevalue.UnmarshalMap(item, &chat)
			if err != nil {
				return nil, err
			}

			if !lastTime.IsZero() {
				if chat.CreatedAt.Before(lastTime) {
					continue
				}
			}
			chats = append(chats, chat)
		}
	}

	return chats, nil
}
func (r *DynamoRepository) GetChatByID(ctx context.Context, chatID string) (*entity.Chat, error) {
	result, err := r.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(chatTable),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: chatID},
		},
	})
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	var chat entity.Chat
	err = attributevalue.UnmarshalMap(result.Item, &chat)
	if err != nil {
		return nil, err
	}

	return &chat, nil
}

func (r *DynamoRepository) UpdateChat(ctx context.Context, chat *entity.Chat) error {
	expressionAttributeValues := map[string]types.AttributeValue{
		"updatedAt": &types.AttributeValueMemberS{Value: chat.UpdatedAt.Format(time.RFC3339Nano)},
	}

	if chat.DeletedAt != nil {
		expressionAttributeValues["deletedAt"] = &types.AttributeValueMemberS{Value: chat.DeletedAt.Format(time.RFC3339Nano)}
	}

	_, err := r.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(chatTable),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: chat.ID},
		},
		ExpressionAttributeValues: expressionAttributeValues,
	})
	if err != nil {
		return err

	}

	return nil
}
