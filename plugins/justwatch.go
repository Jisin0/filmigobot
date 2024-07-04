// (c) Jisin0
// Methods to create justwatch search results.

package plugins

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Jisin0/filmigo/justwatch"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

const (
	jWBanner   = "https://telegra.ph/file/23a8bea137de034392f29.jpg"
	jWLogo     = "https://upload.wikimedia.org/wikipedia/commons/e/e1/JustWatch.png"
	jWHomepage = "https://justwatch.com"

	decriptionMaxLength = 200
	jWCountryCode       = "US"
)

var (
	searchMethodJW = "jw"
	// map to cache tgraph url of images.
	jWPosterCache = make(map[string]string)
)

var jWClient = justwatch.NewClient(&justwatch.JustwatchClientOpts{Country: jWCountryCode})

// JWInlineSearch searches for query on justwatch and returns results to be used in inline queries.
func JWInlineSearch(query string) []gotgbot.InlineQueryResult {
	var results []gotgbot.InlineQueryResult

	rawResults, err := jWClient.SearchTitle(query)
	if err != nil {
		return results
	}

	for _, item := range rawResults.Results {
		posterURL := item.Poster.FullURL()
		if posterURL == "" {
			posterURL = jWBanner
		}

		results = append(results, gotgbot.InlineQueryResultPhoto{
			Id:           searchMethodJW + "_" + item.ID,
			PhotoUrl:     posterURL,
			ThumbnailUrl: posterURL,
			Title:        fmt.Sprintf("%s (%v)", item.Title, item.OriginalReleaseYear),
			Description:  item.ShortDescription,
			Caption:      buildSearchCaption(item.TitlePreview),
			ParseMode:    gotgbot.ParseModeHTML,
			ReplyMarkup: &gotgbot.InlineKeyboardMarkup{InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
				{{Text: "Open JustWatch", CallbackData: fmt.Sprintf("open_%s_%s", searchMethodJW, item.ID)}},
			}},
		})
	}

	return results
}

// buildSearchCaption creates the caption for a search item.
func buildSearchCaption(item *justwatch.TitlePreview) string {
	var (
		builder     strings.Builder
		description = item.ShortDescription
	)

	if len(description) > decriptionMaxLength {
		description = description[0:decriptionMaxLength]
	}

	builder.WriteString(fmt.Sprintf("🎯 <b><a href='%s'>%s (%v)</a></b>\n", jWHomepage+item.Path, item.OriginalTitle, item.OriginalReleaseYear))
	builder.WriteString(fmt.Sprintf("<i>%s</i>\n\n", item.Genres.ToString(" | ")))
	builder.WriteString(fmt.Sprintf("<tg-spoiler><i>%s</i></tg-spoiler>", description))

	return builder.String()
}

// Gets a justwatch title by id and build the message that should be sent or edited.
func GetJWTitle(id string) (gotgbot.InputMediaPhoto, [][]gotgbot.InlineKeyboardButton, error) {
	var (
		photo   gotgbot.InputMediaPhoto
		buttons [][]gotgbot.InlineKeyboardButton
	)

	title, err := jWClient.GetTitle(id)
	if err != nil {
		return photo, buttons, err
	}

	var captionBuilder strings.Builder

	content := title.Content

	if content == nil {
		fmt.Println("no content found !")
		return photo, buttons, errors.New("title content not found : " + id)
	}

	captionBuilder.WriteString(fmt.Sprintf("<a href='%s'><b>%s</b>", jWHomepage+content.URLPath, content.Title))

	if content.ReleaseYear != 0 {
		captionBuilder.WriteString(fmt.Sprintf("<b> (%v)</b>", content.ReleaseYear))
	}

	captionBuilder.WriteString("</a>")

	if content.AgeCertification != "" {
		captionBuilder.WriteString(fmt.Sprintf(" [<tg-spoiler>%s Rated</tg-spoiler>]", content.AgeCertification))
	}

	captionBuilder.WriteRune('\n')

	if content.OriginalTitle != content.Title {
		captionBuilder.WriteString(fmt.Sprintf("<i>  AKA : %s\n</i>", content.OriginalTitle))
	}

	if content.Interactions != nil {
		captionBuilder.WriteString(fmt.Sprintf("<i>👍 %v | %v 👎</i>", content.Interactions.Likes, content.Interactions.Dislikes))
	}

	if content.Scores != nil {
		captionBuilder.WriteString(fmt.Sprintf("  (<i>%.1f%% ❤️</i>)", content.Scores.JustwatchRating*100))
	}

	captionBuilder.WriteRune('\n')

	if content.ExteranlIDs != nil && content.ExteranlIDs.ImdbID != "" {
		captionBuilder.WriteString(fmt.Sprintf("<b>🚦IMDb:</b> <i><a href='imdb.com/title/%s'>%s", content.ExteranlIDs.ImdbID, content.ExteranlIDs.ImdbID))

		if content.Scores != nil && content.Scores.ImdbRating > 0 {
			captionBuilder.WriteString(fmt.Sprintf(" | %v/10 ⭐", content.Scores.ImdbRating))
		}

		captionBuilder.WriteString("</a></i>\n")
	}

	if content.ReleaseDate != "" {
		captionBuilder.WriteString(fmt.Sprintf("<b>🗓️ Released:</b> %s\n", content.ReleaseDate))
	}

	if content.Runtime != 0 {
		captionBuilder.WriteString(fmt.Sprintf("<b>📟 Runtime:</b> %vmins\n", content.Runtime))
	}

	if len(*content.Genres) > 0 {
		captionBuilder.WriteString(fmt.Sprintf("<b>🎭 Genres:</b> <i>%s</i>\n", content.Genres.ToString(", ")))
	}

	if len(title.Offers) > 0 {
		captionBuilder.WriteString("\n<blockquote expandable>")

		var savedOffers []string

		for _, offer := range title.Offers {
			if contains(savedOffers, offer.URL) {
				continue
			}

			captionBuilder.WriteString(fmt.Sprintf("[<b><a href='%s'>%s</a></b>] ", offer.URL, offer.Package.ClearName))

			savedOffers = append(savedOffers, offer.URL)
		}

		captionBuilder.WriteString("</blockquote>")
	} else {
		captionBuilder.WriteString("<b>No Offers Available</b>")
	}

	var posterURL = content.Poster.FullURL()

	if len(content.Backdrops) > 0 {
		if s, ok := jWPosterCache[id]; ok {
			posterURL = s
		} else {
			file := CreateJWPoster(content.FullBackdrops[0].FullURL(), posterURL, id)

			tGraphURL, err := UploadTelegraph(file, "photo")
			if err == nil {
				posterURL = tGraphURL
				jWPosterCache[id] = tGraphURL
			} else {
				fmt.Println("failed to upload to telegraph " + err.Error())
			}
		}
	}

	if posterURL == "" {
		posterURL = jWBanner
	}

	if len(content.Clips) > 0 {
		var row []gotgbot.InlineKeyboardButton

		switch len(content.Clips) {
		case 1:
			row = append(row, gotgbot.InlineKeyboardButton{Text: content.Clips[0].Name, Url: content.Clips[0].URL})
		default:
			for n, clip := range content.Clips {
				if clip == nil {
					continue
				}

				row = append(row, gotgbot.InlineKeyboardButton{Text: fmt.Sprintf("Clip %v", n+1), Url: clip.URL})

				if n >= 2 {
					break
				}
			}
		}

		buttons = append(buttons, row)
	}

	photo = gotgbot.InputMediaPhoto{
		Media:      gotgbot.InputFileByURL(posterURL),
		Caption:    captionBuilder.String(),
		ParseMode:  gotgbot.ParseModeHTML,
		HasSpoiler: true,
	}

	return photo, buttons, nil
}
