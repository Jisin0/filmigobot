// (c) Jisin0
// Basic commands.

package plugins

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

var (
	startText = `
<i><b>👋 Hey there <tg-spoiler>%s</tg-spoiler>,</b>
This bot can search multiple movie databases to get details about any movie or show.</i>

<blockquote expandable><u><b>DISCLAIMER</b></u>

<b>IMDb:</b> Public, commercial, and/or non-private use of the IMDb data provided by this bot is not allowed.
This bot is only for limited non-commercial use of <a href='https://help.imdb.com/article/imdb/general-information/can-i-use-imdb-data-in-my-software/G5JTRESSHJBBHTGX'>IMDb data</a>.

<b>OMDb:</b> Data distributed by the omdbapi is licensed under the <a href='https://creativecommons.org/licenses/by-nc/4.0/'>CreativeCommons 4.0 License</a> and is free to share, copy, remix, transform and build upon for any Non-Commercial application.

<b>JustWatch:</b> <a href='https://support.justwatch.com/hc/en-us/articles/9567105189405-JustWatch-s-Terms-of-Use#h_01HM8Z0MS8WT2S38ND9KNEJY0Y'>Privacy Policy</a> states not to modify, copy, reverse engineer, reverse assemble or otherwise attempt to discover any source code in the Website, or to frame, scrape, rent, lease, loan, sell, assign, sublicense, distribute or create derivative works based on, or reproduce, display, publicly perform, or otherwise use the Service Content in any way for any public or commercial purpose.
</blockquote>
`

	startButtons = append(inlineSearchButtons, []gotgbot.InlineKeyboardButton{{Text: "🚀 𝘜𝘱𝘥𝘢𝘵𝘦𝘴 🚀", Url: "https://t.me/piroxbots"}})
)

func Start(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.EffectiveMessage

	_, err := bot.SendMessage(update.Chat.Id, fmt.Sprintf(startText, mention(update.From)), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML, LinkPreviewOptions: &gotgbot.LinkPreviewOptions{IsDisabled: true}, ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: startButtons}})
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
