package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

// Copy Pasted from github.com/StarkBotsIndustries/telegraph
// Upload photo/video to Telegra.ph on the '/upload' endpoint.
// Media type should either be "video" or "photo". "Animation" is considered "video" here.
func UploadTelegraph(f io.Reader, mediaType string) (string, error) {
	var (
		name string
		b    = &bytes.Buffer{}
		w    = multipart.NewWriter(b)
	)

	if mediaType == "video" {
		name = "file.mp4"
	} else {
		name = "file.jpg"
	}

	part, err := w.CreateFormFile(mediaType, name)
	if err != nil {
		return "", err
	}

	io.Copy(part, f)
	w.Close()

	r, err := http.NewRequest("POST", "https://te.legra.ph/upload", bytes.NewReader(b.Bytes()))
	if err != nil {
		return "", err
	}

	r.Header.Set("Content-Type", w.FormDataContentType())

	c := &http.Client{}

	resp, err := c.Do(r)
	if err != nil {
		return "", err
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var jsonData uploadResult

	json.Unmarshal(content, &jsonData.Source)
	if jsonData.Source == nil {
		var err errorUpload
		json.Unmarshal(content, &err)
		return "", fmt.Errorf(err.Error)
	}

	return "telegra.ph" + jsonData.Source[0].Src, err
}

type uploadResult struct {
	Source []source
}

type source struct {
	Src string `json:"src"`
}

type errorUpload struct {
	Error string `json:"error"`
}
