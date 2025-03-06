package transform

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
)

// converts a spectrogram to a heat map image
func VisualizeSpectrogram(spectrogram [][]complex128, outputPath string) error {
	// Determine dimensions of the spectrogram
	numWindows := len(spectrogram)
	numFreqBins := len(spectrogram[0])

	// Create a new grayscale image
	img := image.NewRGBA(image.Rect(0, 0, numFreqBins, numWindows))

	log.Printf("num window: %d\n", numWindows)

	// Scale the values in the spectrogram to the range [0, 255]
	maxMagnitude := 0.0
	for i := 0; i < numWindows; i++ {
		for j := 0; j < numFreqBins; j++ {
			magnitude := cmplx.Abs(spectrogram[i][j])
			if magnitude > maxMagnitude {
				maxMagnitude = magnitude
			}
		}
	}

	// Convert spectrogram values to pixel intensities
	for i := 0; i < numWindows; i++ {
		for j := 0; j < numFreqBins; j++ {
			magnitude := cmplx.Abs(spectrogram[i][j])
			intensity := uint8(math.Floor(255 * (magnitude / maxMagnitude)))
			img.Set(j, i, color.RGBA{intensity, 0, 0, 255})
		}
	}

	// Save the image to a PNG file
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}
