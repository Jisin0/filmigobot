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

// CbOpen handles callbacks from open_ buttons in search results.
func CbOpen(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.CallbackQuery

	// callback data structure: open_<method>_<id>

	split := strings.Split(update.Data, "_")
	if len(split) < 3 {
		update.Answer(bot, &gotgbot.AnswerCallbackQueryOpts{Text: "Bad Callback Data !", ShowAlert: true})
		return ext.EndGroups
	}

	var (
		method = split[1]
		id     = split[2]

		photo   gotgbot.InputMediaPhoto
		buttons [][]gotgbot.InlineKeyboardButton
		err     error
	)

	switch method {
	case searchMethodJW:
		photo, buttons, err = GetJWTitle(id)
	case searchMethodIMDb:
		photo, buttons, err = GetIMDbTitle(id)
	case searchMethodOMDb:
		photo, buttons, err = GetOMDbTitle(id)
	default:
		fmt.Println("unknown method on cbopen: " + method)
		photo, buttons, err = GetJWTitle(id)
	}

	if err != nil {
		fmt.Printf("cbopen: %v", err)
		update.Answer(bot, &gotgbot.AnswerCallbackQueryOpts{Text: "I Couldn't Fetch Data on That Movie 🤧\nPlease Try Again Later or Contact Admins !", ShowAlert: true})
		return nil
	}

	_, _, err = update.Message.EditMedia(bot, photo, &gotgbot.EditMessageMediaOpts{ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: buttons}})
	if err != nil {
		fmt.Printf("cbopen: %v", err)
	}

	return nil
}
