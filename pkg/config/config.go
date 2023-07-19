package config

import (
	"os"

	traqoauth2 "github.com/ras0q/traq-oauth2"
)

var BotOAuth2Config = traqoauth2.NewConfig(
	os.Getenv("TRAQ_BOT_OAUTH2_CLIENT_ID"),
	os.Getenv("TRAQ_BOT_OAUTH2_REDIRECT_URL"),
)
