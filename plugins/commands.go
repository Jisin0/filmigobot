// (c) Jisin0
// Basic commands.

package plugins

import (
	"fmt"
	"strings"

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

	PrivacyText = `
This bot does not connect to any datbase and hence does not store any user data in any form.
`
)

var (
	startButtons = append([][]gotgbot.InlineKeyboardButton{{{Text: "About ℹ️", CallbackData: "cmd_ABOUT"}, {Text: "Help ❓", CallbackData: "cmd_HELP"}}}, inlineSearchButtons...)
)

func Start(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.EffectiveMessage

	_, err := bot.SendMessage(update.Chat.Id, fmt.Sprintf(startText, mention(update.From)), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML, LinkPreviewOptions: &gotgbot.LinkPreviewOptions{IsDisabled: true}, ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: startButtons}})
	if err != nil {
		fmt.Println(err)
	}

	return ext.EndGroups
}

// CbCommand handles callback from command buttons.
func CbCommand(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.CallbackQuery

	split := strings.SplitN(update.Data, "_", 2)
	if len(split) < 2 {
		update.Answer(bot, &gotgbot.AnswerCallbackQueryOpts{Text: "Bad Callback Data !", ShowAlert: true})
		return nil
	}

	var (
		cmd     = strings.ToUpper(split[1])
		text    string
		buttons [][]gotgbot.InlineKeyboardButton
	)

	switch cmd {
	case "START":
		text = fmt.Sprintf(startText, mention(ctx.EffectiveUser))
		buttons = startButtons
	default:
		if s, k := allTexts[cmd]; k {
			text = s
		} else {
			text, _ = allTexts["NOTFOUND"]
		}

		buttons, _ = allButtons[cmd]
	}

	_, _, err := update.Message.EditText(bot, text, &gotgbot.EditMessageTextOpts{ParseMode: gotgbot.ParseModeHTML, ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: buttons}, LinkPreviewOptions: &gotgbot.LinkPreviewOptions{IsDisabled: true}})
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

// CommandHandler handles any command except start.
func CommandHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.EffectiveMessage

	cmd := strings.ToUpper(strings.Split(strings.ToLower(strings.Fields(update.GetText())[0]), "@")[0][1:])

	var text string
	if s, k := allTexts[cmd]; k {
		text = s
	} else {
		if update.Chat.Type != gotgbot.ChatTypePrivate {
			return nil
		}

		text, _ = allTexts["NOTFOUND"]
	}

	buttons, _ := allButtons[cmd]

	_, err := bot.SendMessage(update.Chat.Id, text, &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML, LinkPreviewOptions: &gotgbot.LinkPreviewOptions{IsDisabled: true}, ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: buttons}})
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
