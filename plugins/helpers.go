// (c) Jisin0
// Helper methods.

package plugins

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/Jisin0/filmigo/types"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

// Parses the time returned from imdb into human-readable format.
func parseIMDbDuration(s string) string {
	s = strings.ReplaceAll(s, "PT", "")
	s = strings.ReplaceAll(s, "H", "h ")
	s = strings.ReplaceAll(s, "M", "min ")
	s = strings.ReplaceAll(s, "S", "s")

	return s
}

// Parses the duration of an imdb trailer ie. PT1M35S becomes 1:35.
func parseIMDbTrailerDuration(duration string) string {
	duration = strings.ReplaceAll(duration, "PT", "")

	var (
		minutes, seconds int
		err              error
	)

	duration = strings.TrimSpace(duration)

	for i := 0; i < len(duration); i++ {
		if unicode.IsDigit(rune(duration[i])) {
			continue
		}

		switch duration[i] {
		case 'M':
			minutes, err = strconv.Atoi(duration[:i])
			if err != nil {
				break
			}

			duration = duration[i+1:]
			i = -1
		case 'S':
			seconds, err = strconv.Atoi(duration[:i])
			if err != nil {
				break
			}

			duration = duration[i+1:]
		}
	}

	formattedDuration := fmt.Sprintf("%02d:%02d", minutes, seconds)
	return formattedDuration
}

// Returns a html-formatted string from the given list.
func htmlLinkList(elems types.Links, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		if elems[0].Text == "" {
			return ""
		}

		return fmt.Sprintf("<a href='%s'>%s</a>", elems[0].Href, elems[0].Text)
	}

	var b strings.Builder

	if elems[0].Text != "" {
		b.WriteString(fmt.Sprintf("<a href='%s'>%s</a>", elems[0].Href, elems[0].Text))
	}

	for _, e := range elems[1:] {
		if e.Text == "" {
			continue
		}

		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("<a href='%s'>%s</a>", e.Href, e.Text))
	}

	return b.String()
}

// Pretty much what it says.
func capitalizeFirstLetter(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}

// Returns a html formatted string that mention's the user
func mention(u *gotgbot.User) string {
	name := u.FirstName
	if u.LastName != "" {
		name = name + " " + u.LastName
	}

	return fmt.Sprintf("<a href='tg://user?id=%v'>%v</a>", u.Id, name)
}

// Checks if a string slice Contains an item.
func Contains(l []string, v string) bool {
	for _, i := range l {
		if i == v {
			return true
		}
	}

	return false
}
