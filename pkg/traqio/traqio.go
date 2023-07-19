package traqio

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/traPtitech/go-traq"
)

type traqWriter struct {
	api       *traq.APIClient
	channelID string
	embed     bool
}

func NewWriter(api *traq.APIClient, channelID string, embed bool) io.Writer {
	return &traqWriter{
		api:       api,
		channelID: channelID,
		embed:     embed,
	}
}

func (w *traqWriter) Write(p []byte) (int, error) {
	_, _, err := w.api.
		MessageApi.
		PostMessage(context.Background(), w.channelID).
		PostMessageRequest(traq.PostMessageRequest{
			Content: fmt.Sprintf("```\n%s\n```", strings.TrimSpace(string(p))),
			Embed:   &w.embed,
		}).
		Execute()
	if err != nil {
		return 0, fmt.Errorf("PostMessage: %w", err)
	}

	return len(p), nil
}
