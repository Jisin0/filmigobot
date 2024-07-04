// (c) Jisin0
// Functions and types to process imdb results.

package plugins

import (
	"fmt"
	"strings"

	"github.com/Jisin0/filmigo/imdb"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

var (
	imdbClient       = imdb.NewClient()
	searchMethodIMDb = "imdb"
)

const (
	imdbLogo     = "https://telegra.ph/file/1720930421ae2b00d9bab.jpg"
	imdbHomepage = "https://imdb.com"
)

// ImdbInlineSearch searches for query on imdb and returns results to be used in inline queries.
func IMDbInlineSearch(query string) []gotgbot.InlineQueryResult {
	var results []gotgbot.InlineQueryResult

	rawResults, err := imdbClient.SearchTitles(query)
	if err != nil {
		return results
	}

	for _, item := range rawResults.Results {
		posterURL := item.Image.URL
		if posterURL == "" {
			posterURL = imdbLogo
		}

		title := fmt.Sprintf("%s (%v)", item.Title, item.Year)

		results = append(results, gotgbot.InlineQueryResultPhoto{
			Id:           searchMethodIMDb + "_" + item.ID,
			PhotoUrl:     posterURL,
			ThumbnailUrl: posterURL,
			Title:        title,
			Description:  item.Description,
			Caption:      fmt.Sprintf("<b><a href='https://imdb.com/title/%s'>%s</a></b>", item.ID, title),
			ParseMode:    gotgbot.ParseModeHTML,
			ReplyMarkup: &gotgbot.InlineKeyboardMarkup{InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
				{{Text: "Open IMDb", CallbackData: fmt.Sprintf("open_%s_%s", searchMethodIMDb, item.ID)}},
			}},
		})
	}

	return results
}

// Gets an imdb title by it's id and returns an InputPhoto to be used.
func GetIMDbTitle(id string) (gotgbot.InputMediaPhoto, [][]gotgbot.InlineKeyboardButton, error) {
	var (
		photo   gotgbot.InputMediaPhoto
		buttons [][]gotgbot.InlineKeyboardButton
	)

	title, err := imdbClient.GetMovie(id)
	if err != nil {
		return photo, buttons, err
	}

	var captionBuilder strings.Builder

	captionBuilder.WriteString(fmt.Sprintf("<b>🎪 %s: <a href='%s'>%s", title.Type, title.URL, title.Title))

	if title.ReleaseYear != "" {
		captionBuilder.WriteString(fmt.Sprintf(" (%s)", title.ReleaseYear))
	}

	captionBuilder.WriteString("</a></b>\n")

	if title.Aka != "" {
		captionBuilder.WriteString(fmt.Sprintf("   [AKA: <code>%s</code>]\n", title.Aka))
	}

	if rating := title.Rating; rating.Value > 0 {
		captionBuilder.WriteString(fmt.Sprintf("<b>🏆 User Ratings: %.1f / 10 </b>", rating.Value))
		captionBuilder.WriteString(fmt.Sprintf("<code>(based on %v votes ", rating.Votes))

		if rating.Best > 0 {
			captionBuilder.WriteString(fmt.Sprintf("best %.1f & worst %.1f", rating.Best, rating.Worst))
		}

		captionBuilder.WriteString(")</code>\n")
	}

	if title.Releaseinfo != "" {
		captionBuilder.WriteString(fmt.Sprintf("<b>🗓 Release Info:</b> <a href='%s'>%s</a>\n", title.URL+"releaseinfo", title.Releaseinfo))
	}

	if title.Runtime != "" {
		captionBuilder.WriteString(fmt.Sprintf("<b>🕰 Duration:</b> <code>%s</code>\n", parseIMDbDuration(title.Runtime)))
	}

	if len(title.Languages) > 0 {
		captionBuilder.WriteString(fmt.Sprintf("<b>🎧 Language:</b> %s\n", htmlLinkList(title.Languages, "|")))
	}

	if len(title.Genres) > 0 {
		captionBuilder.WriteString(fmt.Sprintf("<b>🎭 Genres:</b> <i>%s</i>\n", strings.Join(title.Genres, ", ")))
	}

	if title.Plot != "" {
		captionBuilder.WriteString(fmt.Sprintf("<b>📋 Story Line:</b> <tg-spoiler>%s<a href='%s'>..</a></tg-spoiler>\n", title.Plot, title.URL+"plotsummary"))
	}

	if len(title.Directors) > 0 {
		captionBuilder.WriteString(fmt.Sprintf("<b>🎥 Director:</b> %s\n", htmlLinkList(title.Directors, " ")))
	}

	if len(title.Actors) > 0 {
		captionBuilder.WriteString(fmt.Sprintf("<b>🎎 Actors/Actress:</b> %s\n", htmlLinkList(title.Actors, " ")))
	}

	if len(title.Writers) > 0 {
		captionBuilder.WriteString(fmt.Sprintf("<b>✍️ Writer:</b> %s\n", htmlLinkList(title.Writers, " ")))
	}

	if title.Trailer.URL != "" {
		buttons = append(buttons, []gotgbot.InlineKeyboardButton{{Text: fmt.Sprintf("🎞 IMDb Trailer (%s)", parseIMDbTrailerDuration(title.Trailer.Duration)), Url: title.Trailer.URL}})
	}

	photo = gotgbot.InputMediaPhoto{
		Media:      gotgbot.InputFileByURL(title.PosterURL),
		Caption:    captionBuilder.String(),
		ParseMode:  gotgbot.ParseModeHTML,
		HasSpoiler: true,
	}

	return photo, buttons, nil
}
