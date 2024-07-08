// (c) Jisin0
// Basic commands.

package plugins

import (
	"fmt"

	"github.com/Jisin0/filmigobot/configs"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

var (
	startButtons = append(inlineSearchButtons, []gotgbot.InlineKeyboardButton{{Text: "Source Code", Url: "github.com/Jisin0/filmigobot"}})
)

func Start(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.EffectiveMessage

	_, err := bot.SendMessage(update.Chat.Id, fmt.Sprintf(configs.StartText, mention(update.From)), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML, LinkPreviewOptions: &gotgbot.LinkPreviewOptions{IsDisabled: true}, ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: startButtons}})
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
