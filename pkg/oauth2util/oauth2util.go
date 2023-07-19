package oauth2util

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/ras0q/kinano-go/pkg/cache"
	traqoauth2 "github.com/ras0q/traq-oauth2"
	"github.com/spf13/cobra"
	"github.com/traPtitech/go-traq"
	"golang.org/x/oauth2"
)

const tokenKey = "token"

func RetrieveToken(ctx context.Context, cmd *cobra.Command, conf *traqoauth2.Config) (*oauth2.Token, error) {
	tok, ok := cache.Default().Get(tokenKey)
	if !ok || tok == nil {
		codeVerifier, authURL, err := conf.AuthorizeWithPKCE(traqoauth2.CodeChallengeS256, "state")
		if err != nil {
			return nil, fmt.Errorf("authorize: %w", err)
		}

		codeCh, err := startCallbackServer()
		if err != nil {
			return nil, fmt.Errorf("start callback server: %w", err)
		}

		cmd.Println("Please access the following URL and get the code.\n" + authURL)

		tok, err := conf.CallbackWithPKCE(ctx, codeVerifier, <-codeCh)
		if err != nil {
			return nil, fmt.Errorf("callback: %w", err)
		}

		cache.Default().SetDefault(tokenKey, tok)

		return tok, nil
	}

	if tok, ok := tok.(*oauth2.Token); ok {
		return tok, nil
	}

	return nil, fmt.Errorf("invalid token type")
}

func startCallbackServer() (chan string, error) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return nil, err
	}

	codeCh := make(chan string)

	//nolint:errcheck
	go http.Serve(listener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		codeCh <- r.FormValue("code")
		listener.Close()
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Login successful!"))
	}))

	return codeCh, nil
}

func NewTRAQClient(ctx context.Context, conf *traqoauth2.Config, tok *oauth2.Token) *traq.APIClient {
	traqconf := traq.NewConfiguration()
	traqconf.HTTPClient = conf.Client(ctx, tok)

	return traq.NewAPIClient(traqconf)
}
