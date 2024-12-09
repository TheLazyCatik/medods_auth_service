package main

import (
	gate "auth_service/infra/gateway/email"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSendEmailSuccess(t *testing.T) {
	emailGate := gate.NewEmailGate()

	err := emailGate.SendNewIPLoginNotificationEmail(context.Background(), "1.1.1.1", "ifigurin1411@gmail.com", time.Now().UTC())

	require.NoError(t, err)
}
