/*
Copyright © 2023 ras0q

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/traPtitech/go-traq"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
	"github.com/traPtitech/traq-ws-bot/payload"
)

// metamorphoseCmd represents the metamorphose command
func metamorphoseCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "metamorphose",
		Short: "metamorphose into you",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			bot, ok := ctx.Value("bot").(*traqwsbot.Bot)
			if !ok {
				return fmt.Errorf("failed to get bot from context")
			}

			p, ok := ctx.Value("payload").(*payload.MessageCreated)
			if !ok {
				return fmt.Errorf("failed to get payload from context")
			}

			handleError := func(err error, res *http.Response, msg string) error {
				if err != nil {
					return fmt.Errorf("%s: %w", msg, err)
				}
				if res.StatusCode != http.StatusOK {
					return fmt.Errorf("%s: %d %s", msg, res.StatusCode, res.Status)
				}

				return nil
			}

			f, res, err := bot.API().PublicApi.GetPublicUserIcon(ctx, p.Message.User.Name).Execute()
			if err := handleError(err, res, "get user icon"); err != nil {
				return err
			}
			if f == nil {
				return fmt.Errorf("user icon is nil")
			}
			defer (*f).Close()

			res, err = bot.API().MeApi.ChangeMyIcon(ctx).File(*f).Execute()
			if err := handleError(err, res, "change my icon"); err != nil {
				return err
			}

			displayName := p.Message.User.DisplayName + "きなの"
			res, err = bot.API().MeApi.EditMe(ctx).PatchMeRequest(traq.PatchMeRequest{
				DisplayName: &displayName,
			}).Execute()
			//nolint: revive
			if err := handleError(err, res, "edit me"); err != nil {
				return err
			}

			return nil
		},
	}
}
