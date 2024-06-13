package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"slices"
	"strings"
)

var allowedTokens = strings.Split(os.Getenv("BOT_TOKENS"), " ")
var lenAllowedTokens = len(allowedTokens)

type AuthRequest struct {
	Token string `json:"token"`
}

func Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error": "request has no body"}`))
		return
	}

	var update AuthRequest

	err = json.Unmarshal(body, &update)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error": "failed to decode request body"}`))
		return
	}

	if lenAllowedTokens > 0 && !slices.Contains(allowedTokens, update.Token) {
		w.WriteHeader(203)
		w.Write([]byte(`{"error": "unathorized bot token"}`))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(`{"message": "authorized"}`))
}
