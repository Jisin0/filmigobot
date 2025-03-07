// Setup and initialize environment variables.

package plugins

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	BotToken      string // bot token from @botfather
	Port          string // port to run webapp
	DefaultMethod string // default search method
	OmdbApiKey    string // omdb api key
)

const stringTrue = true

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("load environment variables failed")
	}

	BotToken = os.Getenv("BOT_TOKEN")
	Port = os.Getenv("PORT")
	DefaultMethod = os.Getenv("DEFAULT_SEARCH_METHOD")
	OmdbApiKey = os.Getenv("OMDB_API_KEY")
}
