package plugins

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

// https://github.com/StarkBotsIndustries/telegraph
// Upload photo/video to Telegra.ph on the '/upload' endpoint.
// Media type should either be "video" or "photo". "Animation" is considered "video" here.
func UploadTelegraph(data *bytes.Buffer, mediaType string) (string, error) {
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

	_, err = part.Write(data.Bytes())
	if err != nil {
		return "", err
	}

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

	err = json.Unmarshal(content, &jsonData.Source)
	if err != nil {
		return "", err
	}

	if jsonData.Source == nil {
		var err errorUpload

		er := json.Unmarshal(content, &err)
		if er != nil {
			fmt.Println(er)
		}

		return "", errors.New(err.Error)
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
