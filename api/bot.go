package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"slices"
	"strings"

	"github.com/Jisin0/filmigobot/plugins"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

var (
	allowedTokens    = strings.Split(os.Getenv("BOT_TOKENS"), " ")
	lenAllowedTokens = len(allowedTokens)
)

// Handles all incoming traffic from webhooks.
func Bot(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	_, botToken := path.Split(url)

	bot, _ := gotgbot.NewBot(botToken, &gotgbot.BotOpts{DisableTokenCheck: true})

	// Delete the webhook incase token is unauthorized.
	if lenAllowedTokens > 0 && !slices.Contains(allowedTokens, botToken) {
		bot.DeleteWebhook(&gotgbot.DeleteWebhookOpts{}) // nolint:errcheck
	}

	var update gotgbot.Update

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading request body: %v", err)
		w.WriteHeader(200)
		return
	}

	err = json.Unmarshal(body, &update)
	if err != nil {
		fmt.Println("failed to unmarshal body ", err)
		w.WriteHeader(200)
		return
	}

	ctx := ext.NewContext(&update, map[string]interface{}{})

	if msg := ctx.Message; msg != nil {
		if len(msg.Entities) > 0 {
			if msg.Entities[0].Type == "bot_command" {
				split := strings.Split(strings.ToLower(strings.Fields(msg.Text)[0]), "@")
				cmd := split[0][1:]

				switch cmd {
				case "start":
					plugins.Start(bot, ctx) // nolint:errcheck
				}
			}
		}
	} else if ctx.InlineQuery != nil {
		plugins.InlineQueryHandler(bot, ctx) // nolint:errcheck
	} else if ctx.ChosenInlineResult != nil {
		plugins.InlineResultHandler(bot, ctx) // nolint:errcheck
	}

	w.WriteHeader(200)
}
