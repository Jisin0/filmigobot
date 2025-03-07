// (c) Jisin0
// Functions and types to search using omdb.

package plugins

import (
	"fmt"
	"strings"

	"github.com/Jisin0/filmigo/omdb"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

const (
	omdbBanner   = "https://telegra.ph/file/e810982a269773daa42a9.png"
	omdbHomepage = "https://omdbapi.com"
	notAvailable = "N/A"
)

var (
	omdbClient       *omdb.OmdbClient
	searchMethodOMDb = "omdb"
)

func init() {
	if OmdbApiKey != "" {
		omdbClient = omdb.NewClient(OmdbApiKey)

		inlineSearchButtons = append(inlineSearchButtons, []gotgbot.InlineKeyboardButton{{Text: "🔍 Search OMDb", SwitchInlineQueryCurrentChat: &inlineOMDbSwitch}})
	}
}

// OmdbInlineSearch searches for query on omdb and returns results to be used in inline queries.
func OMDbInlineSearch(query string) []gotgbot.InlineQueryResult {
	var results []gotgbot.InlineQueryResult

	if omdbClient == nil {
		return results
	}

	rawResults, err := omdbClient.Search(query)
	if err != nil {
		return results
	}

	for _, item := range rawResults.Results {
		posterURL := item.Poster
		if posterURL == notAvailable {
			posterURL = omdbBanner
		}

		results = append(results, gotgbot.InlineQueryResultPhoto{
			Id:           searchMethodOMDb + "_" + item.ImdbID,
			PhotoUrl:     posterURL,
			ThumbnailUrl: posterURL,
			Title:        item.Title,
			Description:  fmt.Sprintf("%s | %s", item.Type, item.Year),
			Caption:      fmt.Sprintf("<b><a href='https://imdb.com/title/%s'>%s</a></b>", item.ImdbID, item.Title),
			ParseMode:    gotgbot.ParseModeHTML,
			ReplyMarkup: &gotgbot.InlineKeyboardMarkup{InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
				{{Text: "Open OMDb", CallbackData: fmt.Sprintf("open_%s_%s", searchMethodOMDb, item.ImdbID)}},
			}},
		})
	}

	return results
}

// Gets an imdb title by it's id and returns an InputPhoto to be used.
func GetOMDbTitle(id string) (gotgbot.InputMediaPhoto, [][]gotgbot.InlineKeyboardButton, error) {
	var (
		photo   gotgbot.InputMediaPhoto
		buttons [][]gotgbot.InlineKeyboardButton
	)

	title, err := omdbClient.GetMovie(&omdb.GetMovieOpts{ID: id})
	if err != nil {
		return photo, buttons, err
	}

	var captionBuilder strings.Builder

	url := imdbHomepage + "/title/" + title.ImdbID

	captionBuilder.WriteString(fmt.Sprintf("<b>🎪 %s: <a href='%s'>%s", capitalizeFirstLetter(title.Type), url, title.Title))

	if title.Year != notAvailable {
		captionBuilder.WriteString(fmt.Sprintf(" (%s)", title.Year))
	}

	captionBuilder.WriteString("</a></b>\n")

	if title.Rated != notAvailable {
		captionBuilder.WriteString(fmt.Sprintf("   [<code>%s</code> ʀᴀᴛᴇᴅ]\n", title.Rated))
	}

	if title.ImdbRating != notAvailable {
		captionBuilder.WriteString(fmt.Sprintf("<b>🏆 Usᴇʀ Rᴀᴛɪɴɢs: %s / 10 </b>", title.ImdbRating))

		if title.ImdbVotes != notAvailable {
			captionBuilder.WriteString(fmt.Sprintf("<code>(based on %v users rating)</code>", title.ImdbVotes))
		}

		captionBuilder.WriteRune('\n')
	}

	if title.Released != notAvailable {
		captionBuilder.WriteString(fmt.Sprintf("<b>🗓 Rᴇʟᴇᴀsᴇ Dᴀᴛᴇ:</b> <a href='%s'>%s</a>\n", url+"/releaseinfo", title.Released))
	}

	if title.Runtime != notAvailable {
		captionBuilder.WriteString(fmt.Sprintf("<b>🕰 Dᴜʀᴀᴛɪᴏɴ:</b> <code>%s</code>\n", title.Runtime))
	}

	if title.Languages != notAvailable {
		captionBuilder.WriteString(fmt.Sprintf("<b>🎧 Lᴀɴɢᴜᴀɢᴇ:</b> %s\n", title.Languages))
	}

	if title.Genres != notAvailable {
		captionBuilder.WriteString(fmt.Sprintf("<b>🎭 Gᴇɴʀᴇs:</b> <i>%s</i>\n", title.Genres))
	}

	if title.BoxOffice != notAvailable {
		captionBuilder.WriteString(fmt.Sprintf("<b>💸 Bᴏx Oғғɪᴄᴇ:</b> %s\n", title.BoxOffice))
	}

	if title.Plot != notAvailable {
		captionBuilder.WriteString(fmt.Sprintf("<b>📋 Sᴛᴏʀy Lɪɴᴇ:</b> <tg-spoiler>%s<a href='%s'>..</a></tg-spoiler>\n", title.Plot, url+"/plotsummary"))
	}

	if title.Director != notAvailable {
		captionBuilder.WriteString(fmt.Sprintf("<b>🎥 Dɪʀᴇᴄᴛᴏʀ:</b> <a href='%s'>%s</a>\n", url+"/fullcredits#director", title.Director))
	}

	if title.Actors != notAvailable {
		captionBuilder.WriteString(fmt.Sprintf("<b>🎎 Aᴄᴛᴏʀs:</b> <a href='%s'>%s</a>\n", url+"/fullcredits#cast", title.Actors))
	}

	if title.Writers != notAvailable {
		captionBuilder.WriteString(fmt.Sprintf("<b>✍️ Wʀɪᴛᴇʀ:</b> <a href='%s'>%s</a>\n", url+"/fullcredits#writer", title.Writers))
	}

	buttons = append(buttons, []gotgbot.InlineKeyboardButton{{Text: "🔗 Read More ...", Url: url}})

	photo = gotgbot.InputMediaPhoto{
		Media:      gotgbot.InputFileByURL(title.Poster),
		Caption:    captionBuilder.String(),
		ParseMode:  gotgbot.ParseModeHTML,
		HasSpoiler: true,
	}

	return photo, buttons, nil
}
