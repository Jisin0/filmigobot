// (c) Jisin0
// Handler for inline queries.

package plugins

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

var (
	defaultSearchMethod = searchMethodJW
	allSearchMethods    = []string{searchMethodIMDb, searchMethodOMDb, searchMethodJW}
)

// Search methods with a whitespace added after for a seamless search.
var (
	inlineIMDbSwitch = searchMethodIMDb + " "
	inlineJWSwitch   = searchMethodJW + " "
	inlineOMDbSwitch = searchMethodOMDb + " "
)

var (
	startSearchingButton = &gotgbot.InlineQueryResultsButton{Text: "Start typing the name of your movie to search ...", StartParameter: "nvm"}
	searchResultsButton  = &gotgbot.InlineQueryResultsButton{Text: "Here Are Your Results 👇", StartParameter: "nvm2"}

	notFoundImage = "https://telegra.ph/file/24788bfd2b087c292fbe2.jpg"

	inlineSearchButtons = [][]gotgbot.InlineKeyboardButton{
		{{Text: "📺 Search IMDb", SwitchInlineQueryCurrentChat: &inlineIMDbSwitch}},
		{{Text: "💻 Search OTT", SwitchInlineQueryCurrentChat: &inlineJWSwitch}},
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
	switch DefaultMethod {
	case "":
		DefaultMethod = defaultSearchMethod
	case searchMethodIMDb, searchMethodJW, searchMethodOMDb:
	default:
		fmt.Printf("error: unknown search method \"%s\", using default method \"%s\"\n", DefaultMethod, defaultSearchMethod)
	}
}

// Function that handles all inline queries.
func InlineQueryHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.InlineQuery

	fullQuery := update.Query
	if len(fullQuery) < 1 {
		_, err := update.Answer(bot, []gotgbot.InlineQueryResult{}, &gotgbot.AnswerInlineQueryOpts{CacheTime: defaultCacheTime, Button: startSearchingButton})

		return err
	}

	var (
		method string
		query  = fullQuery
	)

	args := strings.SplitN(query, " ", 2)
	if len(args) > 1 {
		method = strings.ToLower(args[0])
		query = args[1]
	}

	if Contains(allSearchMethods, method) && len(query) < 1 {
		_, err := update.Answer(bot, []gotgbot.InlineQueryResult{}, &gotgbot.AnswerInlineQueryOpts{CacheTime: defaultCacheTime, Button: startSearchingButton})

		return err
	}

	results := getInlineResults(method, query, fullQuery)
	if len(results) < 1 {
		_, err := update.Answer(bot, []gotgbot.InlineQueryResult{noResultsArticle}, &gotgbot.AnswerInlineQueryOpts{
			CacheTime: defaultCacheTime,
			Button:    searchResultsButton,
		})

		return err
	}

	_, err := update.Answer(bot, results, &gotgbot.AnswerInlineQueryOpts{
		CacheTime: defaultCacheTime,
		Button:    searchResultsButton,
	})

	return err
}

// Returns inline results based on given method.
//
//nolint:unparam // linter claims I don't use fullQuery.
func getInlineResults(method, query, fullQuery string) []gotgbot.InlineQueryResult {
	switch method {
	case searchMethodJW:
		return JWInlineSearch(query)
	case searchMethodIMDb:
		return IMDbInlineSearch(query)
	case searchMethodOMDb:
		return OMDbInlineSearch(query)
	default:
		return getInlineResults(DefaultMethod, fullQuery, fullQuery)
	}
}
