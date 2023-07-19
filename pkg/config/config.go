package config

import (
	"os"

	traqoauth2 "github.com/ras0q/traq-oauth2"
)

type ContextKey string

const PayloadKey ContextKey = "payload"

var (
	AccessToken = os.Getenv("TRAQ_BOT_ACCESS_TOKEN")

	BotOAuth2Config = traqoauth2.NewConfig(
		os.Getenv("TRAQ_BOT_OAUTH2_CLIENT_ID"),
		os.Getenv("TRAQ_BOT_OAUTH2_REDIRECT_URL"),
	)
)
