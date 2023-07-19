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
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ras0q/kinano-go/cmd"
	"github.com/ras0q/kinano-go/pkg/config"
	"github.com/ras0q/kinano-go/pkg/traqio"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func main() {
	bot, err := traqwsbot.NewBot(&traqwsbot.Options{AccessToken: config.AccessToken})
	if err != nil {
		panic(err)
	}

	bot.OnError(func(msg string) {
		log.Println(fmt.Errorf("bot error: %s", msg))
	})

	bot.OnMessageCreated(func(p *payload.MessageCreated) {
		if p.Message.User.Bot {
			return
		}

		var (
			ctx  = context.WithValue(context.Background(), config.PayloadKey, p)
			w    = traqio.NewWriter(bot.API(), p.Message.ChannelID, true)
			args = strings.Fields(p.Message.PlainText)
		)

		if e := p.Message.Embedded; len(e) > 0 {
			if e[0].Raw == args[0] {
				args = args[1:]
			}
		}

		if err := cmd.Execute(ctx, w, os.Stderr, args); err != nil {
			if !errors.Is(err, cmd.ErrCommandNotFound) {
				log.Println(fmt.Errorf("cmd.Execute: %w", err))
			}
		}
	})

	if err := bot.Start(); err != nil {
		panic(err)
	}
}
