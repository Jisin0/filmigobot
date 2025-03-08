// (c) Jisin0
// Functions and types to process imdb results.

package plugins

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/Jisin0/filmigo/imdb"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

var (
	imdbClient       = imdb.NewClient()
	searchMethodIMDb = "imdb"
)

const (
	imdbLogo     = "https://telegra.ph/file/1720930421ae2b00d9bab.jpg"
	imdbBanner   = "https://telegra.ph/file/2dd6f7c9ebfb237db4826.jpg"
	imdbHomepage = "https://imdb.com"
)

// ImdbInlineSearch searches for query on imdb and returns results to be used in inline queries.
func IMDbInlineSearch(query string) []gotgbot.InlineQueryResult {
	rawResults, err := imdbClient.SearchTitles(query)
	if err != nil {
		return nil
	}

	results := make([]gotgbot.InlineQueryResult, 0, len(rawResults.Results))

	for _, item := range rawResults.Results {
		posterURL := item.Image.URL
		if posterURL == "" {
			posterURL = imdbLogo
		}

		title := fmt.Sprintf("%s (%v)", item.Title, item.Year)
		url := fmt.Sprintf("https://imdb.com/title/%s", item.ID)

		results = append(results, gotgbot.InlineQueryResultArticle{
			Id:           searchMethodIMDb + "_" + item.ID,
			Url:          url,
			ThumbnailUrl: posterURL,
			Title:        title,
			Description:  item.Description,
			InputMessageContent: gotgbot.InputTextMessageContent{
				MessageText: fmt.Sprintf("<b><a href='%s'>%s</a></b>", url, title),
				ParseMode:   gotgbot.ParseModeHTML,
				LinkPreviewOptions: &gotgbot.LinkPreviewOptions{
					PreferSmallMedia: true,
				},
			},
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
		captionBuilder.WriteString(fmt.Sprintf("   [ᴀᴋᴀ: <code>%s</code>]\n", title.Aka))
	}

	if rating := title.Rating; rating.Value > 0 {
		captionBuilder.WriteString(fmt.Sprintf("<b>🏆 Usᴇʀ Rᴀᴛɪɴɢs: %.1f / 10 </b>", rating.Value))
		captionBuilder.WriteString(fmt.Sprintf("<code>(based on %v votes ", rating.Votes))

		if rating.Best > 0 {
			captionBuilder.WriteString(fmt.Sprintf("best %.1f & worst %.1f", rating.Best, rating.Worst))
		}

		captionBuilder.WriteString(")</code>\n")
	}

	if title.Releaseinfo != "" {
		captionBuilder.WriteString(fmt.Sprintf("<b>🗓 Rᴇʟᴇᴀsᴇ Iɴғᴏ:</b> <a href='%s'>%s</a>\n", title.URL+"releaseinfo", title.Releaseinfo))
	}

	if title.Runtime != "" {
		captionBuilder.WriteString(fmt.Sprintf("<b>🕰 Dᴜʀᴀᴛɪᴏɴ:</b> <code>%s</code>\n", parseIMDbDuration(title.Runtime)))
	}

	if len(title.Languages) > 0 {
		captionBuilder.WriteString(fmt.Sprintf("<b>🎧 Lᴀɴɢᴜᴀɢᴇ:</b> %s\n", htmlLinkList(title.Languages, "|")))
	}

	if len(title.Genres) > 0 {
		captionBuilder.WriteString(fmt.Sprintf("<b>🎭 Gᴇɴʀᴇs:</b> <i>%s</i>\n", strings.Join(title.Genres, ", ")))
	}

	if title.Plot != "" {
		captionBuilder.WriteString(fmt.Sprintf("<b>📋 Sᴛᴏʀy Lɪɴᴇ:</b> <tg-spoiler>%s<a href='%s'>..</a></tg-spoiler>\n", title.Plot, title.URL+"plotsummary"))
	}

	if len(title.Directors) > 0 {
		captionBuilder.WriteString(fmt.Sprintf("<b>🎥 Dɪʀᴇᴄᴛᴏʀ:</b> %s\n", htmlLinkList(title.Directors, " ")))
	}

	if len(title.Actors) > 0 {
		captionBuilder.WriteString(fmt.Sprintf("<b>🎎 Aᴄᴛᴏʀs:</b> %s\n", htmlLinkList(title.Actors, " ")))
	}

	if len(title.Writers) > 0 {
		if str := htmlLinkList(title.Writers, " "); str != "" { // th writers field can contain companies whose names aren't available resulting in an empty string
			captionBuilder.WriteString(fmt.Sprintf("<b>✍️ Wʀɪᴛᴇʀ:</b> %s\n", str))
		}
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

// IMDbCommand handles the /imdb command.
func IMDbCommand(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.EffectiveMessage

	split := strings.SplitN(update.GetText(), " ", 2)
	if len(split) < 2 {
		text := "<i>Please provide a search query or movie id along with this command !\nFor Example:</i>\n  <code>/imdb Inception</code>\n  <code>/imdb tt1375666</code>"
		update.Reply(bot, text, &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})
		return nil
	}

	input := split[1]

	var (
		photo   gotgbot.InputMediaPhoto
		buttons [][]gotgbot.InlineKeyboardButton
		err     error
	)

	if id := regexp.MustCompile(`tt\d+`).FindString(input); id != "" {
		photo, buttons, err = GetIMDbTitle(id)
	} else {
		results, e := imdbClient.SearchTitles(input)
		if e != nil {
			err = e
		} else {
			if len(results.Results) < 1 {
				err = errors.New("No results found")
			} else {
				for _, r := range results.Results {
					buttons = append(buttons, []gotgbot.InlineKeyboardButton{{Text: fmt.Sprintf("%s (%d)", r.Title, r.Year), CallbackData: fmt.Sprintf("open_%s_%s", searchMethodIMDb, r.ID)}})
				}

				photo = gotgbot.InputMediaPhoto{
					Media:     gotgbot.InputFileByURL(imdbBanner),
					Caption:   fmt.Sprintf("<i>👋 Hey <tg-spoiler>%s</tg-spoiler> I've got %d Results for you 👇</i>", mention(ctx.EffectiveUser), len(results.Results)),
					ParseMode: gotgbot.ParseModeHTML,
				}
			}
		}
	}

	if err != nil {
		photo = gotgbot.InputMediaPhoto{
			Caption:   fmt.Sprintf("<i>I'm Sorry %s I Couldn't find Anything for <code>%s</code> 🤧</i>", mention(ctx.EffectiveUser), input),
			Media:     gotgbot.InputFileByURL(imdbBanner),
			ParseMode: gotgbot.ParseModeHTML,
		}

		buttons = [][]gotgbot.InlineKeyboardButton{{{Text: "Search On Google 🔎", Url: fmt.Sprintf("https://google.com/search?q=%s", url.QueryEscape(input))}}}
	}

	_, err = bot.SendPhoto(ctx.EffectiveChat.Id, photo.Media, &gotgbot.SendPhotoOpts{
		Caption:     photo.Caption,
		ParseMode:   photo.ParseMode,
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: buttons},
		HasSpoiler:  photo.HasSpoiler})
	if err != nil {
		fmt.Printf("imdbcommand: %v", err)
	}

	return ext.EndGroups
}
