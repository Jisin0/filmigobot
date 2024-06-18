// (c) Jisin0
// Run the bot in polling environments

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Jisin0/filmigobot/plugins"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/choseninlineresult"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/inlinequery"
)

const (
	defaultGetUpdatesSleep = 15 // number of seconds to sleep when getUpdates fails
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("exiting because no BOT_TOKEN provided")
	}

	b, err := gotgbot.NewBot(token, &gotgbot.BotOpts{
		BotClient: &gotgbot.BaseBotClient{
			Client: http.Client{},
			DefaultRequestOpts: &gotgbot.RequestOpts{
				Timeout: gotgbot.DefaultTimeout,
				APIURL:  gotgbot.DefaultAPIURL,
			},
		},
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// To make sure no other instance of the bot is running
	_, err = b.GetUpdates(&gotgbot.GetUpdatesOpts{})
	if err != nil {
		fmt.Println("duplicate instance found: waiting 15s to fetch updates")
		time.Sleep(time.Second * defaultGetUpdatesSleep)

		return
	}

	// Create updater and dispatcher.
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			fmt.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})

	dispatcher.AddHandlerToGroup(handlers.NewInlineQuery(inlinequery.All, plugins.InlineQueryHandler), 0)
	dispatcher.AddHandlerToGroup(handlers.NewChosenInlineResult(choseninlineresult.All, plugins.InlineResultHandler), 0)
	dispatcher.AddHandlerToGroup(handlers.NewCommand("start", plugins.Start), 0)

	updater := ext.NewUpdater(dispatcher, &ext.UpdaterOpts{})

	// Start receiving updates.
	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true, GetUpdatesOpts: &gotgbot.GetUpdatesOpts{AllowedUpdates: []string{"message", "callback_query", "inline_query", "chosen_inline_result"}}})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}

	fmt.Printf("@%s Started !\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}
