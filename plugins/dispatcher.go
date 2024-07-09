// (c) Jisin0

package plugins

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/choseninlineresult"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/inlinequery"
)

var Dispatcher *ext.Dispatcher = ext.NewDispatcher(&ext.DispatcherOpts{
	// If an error is returned by a handler, log it and continue going.
	Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
		fmt.Println("an error occurred while handling update:", err.Error())
		return ext.DispatcherActionNoop
	},
	MaxRoutines: ext.DefaultMaxRoutines,
})

const (
	commandHandlerGroup = 1
)

func init() {
	Dispatcher.AddHandlerToGroup(handlers.NewInlineQuery(inlinequery.All, InlineQueryHandler), 0)
	Dispatcher.AddHandlerToGroup(handlers.NewChosenInlineResult(choseninlineresult.All, InlineResultHandler), 0)
	Dispatcher.AddHandlerToGroup(handlers.NewCallback(callbackquery.All, CbCommand), 0)

	Dispatcher.AddHandlerToGroup(handlers.NewCommand("start", Start), commandHandlerGroup)
	Dispatcher.AddHandlerToGroup(handlers.NewCommand("imdb", IMDbCommand), commandHandlerGroup)
	Dispatcher.AddHandlerToGroup(handlers.NewMessage(allCommand, CommandHandler), commandHandlerGroup)
}

func allCommand(msg *gotgbot.Message) bool {
	ents := msg.GetEntities()
	if len(ents) != 0 && ents[0].Offset == 0 && ents[0].Type != "bot_command" {
		return false
	}

	text := msg.GetText()

	if r, _ := utf8.DecodeRuneInString(text); r != '/' {
		return false
	}

	split := strings.Split(strings.ToLower(strings.Fields(text)[0]), "@")
	cmd := split[0][1:]

	if cmd == "" {
		return false
	}

	return true
}
