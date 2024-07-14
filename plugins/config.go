// (c) Jisin0

package plugins

import "github.com/PaulSonOfLars/gotgbot/v2"

var allTexts map[string]string = map[string]string{
	"PRIVACY": `<i>This bot does not connect to any datbase and hence <b>does not store any user data</b> in any form.</i>`,

	"ABOUT": `
○ <b>Language</b>: <a href='https://go.dev'>GO</a>
○ <b>Library</b>: <a href='https://github.com/PaulSonOfLars/gotgbot'>GoTgbot</a>
○ <b>Support</b>: <a href='https://FractalProjects'>Fractal</a>
	`,

	"HELP": `
<i>Type my <b>username</b> into any chat to start <b>searching</b> for any movie 👉</i>

<i>Here's a list of my available commands:</i>

/start : Start the bot.
/about : Get some data about the bot.
/help  : Display this help message.
/privacy: Leran how this bot uses your data.
/imdb: Search or get a movie from IMDb.
/jw: Search or get a movie from Justwatch

<i>Use the <b>buttons</b> below to search for a movie here 👇</i>
`,

	"NOTFOUND": `
<i>😐 I don't recognize that command !
Check /help to see how to use me.</i>
`,
}

var allButtons map[string][][]gotgbot.InlineKeyboardButton = map[string][][]gotgbot.InlineKeyboardButton{
	"ABOUT": {{homeButton, helpButton}, {{Text: "Source 🔗", Url: "https://github.com/Jisin0/filmigobot"}}},
	"HELP":  append(inlineSearchButtons, []gotgbot.InlineKeyboardButton{aboutButton, homeButton}),
}

// Single buttons used to build composite markups.
var (
	aboutButton = gotgbot.InlineKeyboardButton{Text: "About ℹ️", CallbackData: "cmd_ABOUT"}
	helpButton  = gotgbot.InlineKeyboardButton{Text: "Help ❓", CallbackData: "cmd_HELP"}
	homeButton  = gotgbot.InlineKeyboardButton{Text: "Home 🏠", CallbackData: "cmd_START"}
)
