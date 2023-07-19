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
package callevent

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ras0q/kinano-go/pkg/config"
	"github.com/ras0q/kinano-go/pkg/oauth2util"
	"github.com/spf13/cobra"
)

// getMyUserTagsCmd represents the getMyUserTags command
func GetMyUserTagsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "getMyUserTags",
		Short: "get my tags",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			conf := config.BotOAuth2Config

			tok, err := oauth2util.RetrieveToken(ctx, cmd, conf)
			if err != nil {
				return fmt.Errorf("get token: %w", err)
			}

			traqClient := oauth2util.NewTRAQClient(ctx, conf, tok)
			tags, res, err := traqClient.MeApi.GetMyUserTags(ctx).Execute()
			if err != nil {
				return fmt.Errorf("get my tags: %w", err)
			}
			if res.StatusCode != http.StatusOK {
				return fmt.Errorf("get my tags: %d, %s", res.StatusCode, res.Status)
			}

			tagContents := make([]string, len(tags))
			for i, tag := range tags {
				tagContents[i] = tag.Tag
			}
			cmd.Println(strings.Join(tagContents, "\n"))

			return nil
		},
	}
}
