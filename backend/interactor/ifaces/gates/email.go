package gate

import (
	"context"
	"time"
)

type EmailGate interface {
	SendNewIPLoginNotificationEmail(ctx context.Context, ipAddress string, email string, timestamp time.Time) error
}
