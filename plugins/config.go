// (c) Jisin0

package plugins

import "github.com/PaulSonOfLars/gotgbot/v2"

var allTexts map[string]string = map[string]string{
	"PRIVACY": `<i>This bot does not connect to any datbase and hence <b>does not store any user data</b> in any form.</i>`,

	"ABOUT": `
○ <b>Language</b>: <a href='https://go.dev'>GO</a>
○ <b>Library</b>: <a href='https://github.com/PaulSonOfLars/gotgbot'>GoTgbot</a>
○ <b>Source</b>: <a href='https://github.com/Jisin/filmigobot'>Repo</a>
○ <b>Support</b>: <a href='https://FractalProjects'>Fractal</a>
	`,

	"HELP": `
<i>Type my <b>username</b> into any chat to start <b>searching</b> for any movie 👉</i>

<i>Here's a list of my available commands:</i>
  /start : Start the bot.
  /about : Get some data about the bot.
  /help  : Display this help message.
  /privacy: Read how this bot uses your data.

<i>Use the buttons below to search for a movie here 👇</i>
`,

	"NOTFOUND": `
<i>😐 I couldn't find that command !
Check /help to see how to use me.</i>
`,
}

var allButtons map[string][][]gotgbot.InlineKeyboardButton = map[string][][]gotgbot.InlineKeyboardButton{
	"ABOUT": {{{Text: "Home", CallbackData: "cmd_START"}, {Text: "Help", CallbackData: "cmd_HELP"}}},
	"HELP":  append(inlineSearchButtons, []gotgbot.InlineKeyboardButton{{Text: "About", CallbackData: "cmd_ABOUT"}, {Text: "Home", CallbackData: "cmd_START"}}),
}
