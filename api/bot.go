package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/Jisin0/filmigobot/plugins"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// Handles all incoming traffic from webhooks.
func Bot(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	_, botToken := path.Split(url)
	fmt.Println(botToken)

	bot, _ := gotgbot.NewBot(botToken, &gotgbot.BotOpts{DisableTokenCheck: true})

	var update gotgbot.Update

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading request body: %v", err)
		return
	}

	json.Unmarshal(body, &update)

	ctx := ext.NewContext(&update, map[string]interface{}{})

	if msg := ctx.Message; msg != nil {
		if len(msg.Entities) > 0 {
			if msg.Entities[0].Type == "bot_command" {
				split := strings.Split(strings.ToLower(strings.Fields(msg.Text)[0]), "@")
				cmd := split[0][1:]

				switch cmd {
				case "start":
					plugins.Start(bot, ctx)
				}
			}
		}
	} else if ctx.InlineQuery != nil {
		plugins.InlineQueryHandler(bot, ctx)
	} else if ctx.ChosenInlineResult != nil {
		plugins.InlineResultHandler(bot, ctx)
	}

	w.WriteHeader(200)
}
