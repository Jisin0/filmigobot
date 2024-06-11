// (c) Jisin0
// Handle the chosen_inline_result event.

package plugins

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// Edit message after the result is chosen.
func InlineResultHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	var (
		update = ctx.ChosenInlineResult
		data   = update.ResultId
	)

	if data == notAvailable {
		return nil
	}

	args := strings.Split(data, "_")
	if len(args) < 2 {
		fmt.Println("bad resultid on choseninlineresult : " + data)
		return nil
	}

	var (
		method = args[0]
		id     = args[1]
	)

	media, buttons, err := getChosenResult(method, id)
	if err != nil {
		// Must do something here
		fmt.Println(err)
		return nil
	}

	_, _, err = bot.EditMessageMedia(
		media,
		&gotgbot.EditMessageMediaOpts{
			InlineMessageId: update.InlineMessageId,
			ReplyMarkup:     gotgbot.InlineKeyboardMarkup{InlineKeyboard: buttons},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

// Returns the inputmessage content to edit with.
func getChosenResult(method, id string) (gotgbot.InputMediaPhoto, [][]gotgbot.InlineKeyboardButton, error) {
	switch method {
	case searchMethodJW:
		return GetJWTitle(id)
	case searchMethodIMDb:
		return GetIMDbTitle(id)
	case searchMethodOMDb:
		return GetOMDbTitle(id)
	default:
		fmt.Println("unknown method on choseninlineresult : " + method)
		return GetJWTitle(id)
	}
}
