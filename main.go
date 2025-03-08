// (c) Jisin0
// Run the bot in polling environments

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Jisin0/filmigobot/plugins"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

const (
	defaultGetUpdatesSleep = 15 // number of seconds to sleep when getUpdates fails
)

func main() {
	// Run a useless http server to get a healthy build on koyeb/render
	go func() {
		port := plugins.Port
		if port == "" {
			port = "8080"
		}

		http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprintf(w, "healthcheck")
		})

		//nolint:gosec // I frankly don't care if it isn't ideal.
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			fmt.Printf("failed to start web server : %v\n", err)
		}
	}()

	token := plugins.BotToken
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

	updater := ext.NewUpdater(plugins.Dispatcher, &ext.UpdaterOpts{})

	// Start receiving updates.
	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true, GetUpdatesOpts: &gotgbot.GetUpdatesOpts{AllowedUpdates: []string{"message", "callback_query", "inline_query", "chosen_inline_result"}}})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}

	fmt.Printf("@%s Started !\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}
