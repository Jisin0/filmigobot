// (c) Jisin0

package plugins

import "github.com/PaulSonOfLars/gotgbot/v2"

var allTexts map[string]string = map[string]string{
	"PRIVACY": `This bot does not connect to any datbase and hence does not store any user data in any form.`,

	"ABOUT": `
○ <b>Language</b>: <a href='https://go.dev'>GO</a>
○ <b>Library</b>: <a href='https://github.com/PaulSonOfLars/gotgbot'>GoTgbot</a>
○ <b>Source</b>: <a href='https://github.com/Jisin/filmigobot'>Repo</a>
○ <b>Support</b>: <a href='https://FractalProjects'>Fractal</a>
	`,

	"HELP": `
<i>Type my username into any chat to start searching for any movie.</i>

Here's a list of my available commands:
  /start : Start the bot.
  /about : Get some data about the bot.
  /help  : Display this help message.

<i>Use the buttons below to search for a movie here 👇</i>
`,

	"NOTFOUND": `
<i>😐 I couldn't find that command !
Check /help to see how to use me.</i>
`,
}

var allButtons map[string][][]gotgbot.InlineKeyboardButton = map[string][][]gotgbot.InlineKeyboardButton{}
