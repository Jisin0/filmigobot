// (c) Jisin0
// Stuff with images like creating justwatch posters.

package plugins

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"

	"github.com/fogleman/gg"
	"golang.org/x/image/webp"
)

const (
	posterRadius       = 20.0
	posterStartPercent = 40
)

// Creates a poster image with a backdrop and poster as overlay.
func CreateJWPoster(backdropURL, posterURL, id string) *os.File {
	fileName := fmt.Sprintf("%s.jpg", id)

	if file, err := os.Open(fileName); err == nil {
		return file
	}

	// Load the backdrop image
	backdrop, err := loadImage(backdropURL)
	if err != nil {
		fmt.Println("Error loading backdrop:", err)
		return nil
	}

	// Load the poster image
	poster, err := loadImage(posterURL)
	if err != nil {
		fmt.Println("Error loading poster:", err)
		return nil
	}

	// Add rounded corners to the poster
	poster = addRoundedCorners(poster, posterRadius)

	// Calculate position for the poster to start at 80% of the width of the backdrop
	backdropBounds := backdrop.Bounds()
	posterBounds := poster.Bounds()

	offset := image.Pt(
		(backdropBounds.Max.X*posterStartPercent/100)-posterBounds.Max.X,
		(backdropBounds.Max.Y-posterBounds.Max.Y)/2, // Center vertically
	)

	// Create a new image with the size of the backdrop
	result := image.NewRGBA(backdrop.Bounds())

	// Draw the backdrop onto the new image
	draw.Draw(result, backdrop.Bounds(), backdrop, image.Point{}, draw.Src)

	// Calculate the rectangle for the poster
	posterRect := image.Rectangle{
		Min: offset,
		Max: offset.Add(poster.Bounds().Size()),
	}

	// Draw the poster onto the new image with the calculated rectangle
	draw.Draw(result, posterRect, poster, image.Point{}, draw.Over)

	// Save the result to a new file
	outFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return nil
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, result, &jpeg.Options{Quality: 100})
	if err != nil {
		fmt.Println("Error saving result:", err)
		return nil
	}

	returnFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return returnFile
}

// loadImage loads an image from a URL
func loadImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	var img image.Image

	switch {
	case strings.Contains(contentType, "jpeg"), strings.Contains(contentType, "jpg"):
		img, err = jpeg.Decode(resp.Body)
	case strings.Contains(contentType, "png"):
		img, err = png.Decode(resp.Body)
	case strings.Contains(contentType, "webp"):
		img, err = webp.Decode(resp.Body)
	default:
		return nil, fmt.Errorf("unsupported image type: %s", contentType)
	}

	if err != nil {
		return nil, err
	}

	return img, nil
}

// addRoundedCorners adds rounded corners to an image
func addRoundedCorners(src image.Image, radius float64) image.Image {
	w := src.Bounds().Dx()
	h := src.Bounds().Dy()

	dc := gg.NewContext(w, h)
	dc.DrawRoundedRectangle(0, 0, float64(w), float64(h), radius)
	dc.Clip()
	dc.DrawImage(src, 0, 0)

	return dc.Image()
}
