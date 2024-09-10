package plugins

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

// UploadEnvssh uploads photo/video to envs.sh.
// Media type should either be "video" or "photo". "Animation" is considered "video" here.
func UploadEnvssh(data *bytes.Buffer, id string) (string, error) {
	// Create a buffer and a multipart writer
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create the file field
	part, err := writer.CreateFormFile("file", id+".jpg")
	if err != nil {
		return "", fmt.Errorf("error creating form file: %v", err)
	}

	// Copy the file into the field
	_, err = io.Copy(part, data)
	if err != nil {
		return "", fmt.Errorf("error copying file: %v", err)
	}

	// Close the writer to finalize the request
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("error closing writer: %v", err)
	}

	// Send the POST request
	req, err := http.NewRequest("POST", "https://envs.sh", &requestBody)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	// Print the response from the server
	return string(body), nil
}
