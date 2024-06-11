// (c) Jisin0
// Handler for inline queries.

package plugins

import (
	"fmt"
	"os"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

var (
	defaulSearchMethod = searchMethodJW
)

var (
	startSearchingButton = &gotgbot.InlineQueryResultsButton{Text: "Start typing the name of your movie to search ...", StartParameter: "nvm"}
	searchResultsButton  = &gotgbot.InlineQueryResultsButton{Text: "Here Are Your Results 👇", StartParameter: "nvm2"}

	notFoundImage = "https://telegra.ph/file/d80303cbff7d4e93bb2e8.png"

	inlineSearchButtons = [][]gotgbot.InlineKeyboardButton{
		{{Text: "📺 Search IMDb", SwitchInlineQueryCurrentChat: &searchMethodIMDb}},
		{{Text: "💻 Search OTT", SwitchInlineQueryCurrentChat: &searchMethodJW}},
	}

	noResultsArticle = gotgbot.InlineQueryResultArticle{
		Id:    notAvailable,
		Title: "No Results Were Found for Your Query !",
		InputMessageContent: gotgbot.InputTextMessageContent{
			MessageText: "<i>👋 Sorry I didn't find anything for that !\nUse the buttons below to Search Again 👇</i>",
			ParseMode:   gotgbot.ParseModeHTML,
		},
		ReplyMarkup:  &gotgbot.InlineKeyboardMarkup{InlineKeyboard: inlineSearchButtons},
		ThumbnailUrl: notFoundImage,
	}
)

const (
	// The time in seconds that results for a query can be cached by a client.
	defaultCacheTime = 2000
)

func init() {
	if defaultMethod := os.Getenv("DEFAULT_SEARCH_METHOD"); defaultMethod != "" {
		if defaultMethod == searchMethodIMDb || defaultMethod == searchMethodOMDb || defaulSearchMethod == searchMethodJW {
			defaulSearchMethod = defaultMethod
		} else {
			fmt.Printf("error: unknown search method \"%s\", using default method \"%s\"\n", defaultMethod, defaulSearchMethod)
		}
	}
}

// Function that handles all inline queries.
func InlineQueryHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.InlineQuery

	fullQuery := update.Query
	if len(fullQuery) < 1 {
		update.Answer(bot, []gotgbot.InlineQueryResult{}, &gotgbot.AnswerInlineQueryOpts{CacheTime: defaultCacheTime, Button: startSearchingButton})
	}

	var (
		method = defaulSearchMethod
		query  = fullQuery
	)

	args := strings.SplitN(query, " ", 2)
	if len(args) > 1 {
		method = strings.ToLower(args[0])
		query = args[1]
	}

	results := getInlineResults(method, query, fullQuery)
	if len(results) < 1 {
		update.Answer(bot, []gotgbot.InlineQueryResult{noResultsArticle}, &gotgbot.AnswerInlineQueryOpts{
			CacheTime: defaultCacheTime,
			Button:    searchResultsButton,
		})
	}

	_, err := update.Answer(bot, results, &gotgbot.AnswerInlineQueryOpts{
		CacheTime: defaultCacheTime,
		Button:    searchResultsButton,
	})
	if err != nil {
		fmt.Printf("error while answering inline query %v\n", err)
	}

	return nil
}

// Returns inline results based on given method.
func getInlineResults(method, query, fullQuery string) []gotgbot.InlineQueryResult {
	switch method {
	case searchMethodJW:
		return JWInlineSearch(query)
	case searchMethodIMDb:
		return IMDbInlineSearch(query)
	case searchMethodOMDb:
		return OMDbInlineSearch(query)
	default:
		return getInlineResults(defaulSearchMethod, fullQuery, fullQuery)
	}
}
