package mail

import (
	"testing"

	"github.com/JuanHeredia3/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	config, err := util.LoadConfig("../")
	require.NoError(t, err)

	sender := NewGmailSender(
		config.EmailSenderName,
		config.EmailSenderAddress,
		config.EmailSenderPassword,
	)

	err = sender.SendEmail(
		"Test Subject",
		"This is a test email content.",
		"recipient@example.com",
		"",
		"",
		[]string{"path/to/attachment1.txt", "path/to/attachment2.jpg"},
	)
	require.NoError(t, err)
}
