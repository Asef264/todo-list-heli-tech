package ports

import "context"

type SQS interface {
	SendMessage(context.Context, string) error
}
