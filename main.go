package main

import (
	"log"

	"delm.dev/music-recognizer/transform"
	"delm.dev/music-recognizer/wav"
)

func main() {
	data, err := wav.ReadWav("./assets/saul.wav")
	if err != nil {
		log.Fatalf("Unable to read info: %s\n", err.Error())
	}
	samples, _ := wav.BytesToSamples(data.Data)
	for _, sample := range samples {
		log.Printf("%f ", sample)
	}
	spectrogram, _ := transform.Spectrogram(samples, data.SampleRate)
	log.Printf("length: %d\n", len(spectrogram))
	transform.ExtractPeaks(spectrogram, data.Duration)
	err = transform.VisualizeSpectrogram(spectrogram, "./image.png")
	if err != nil {
		log.Printf("unable to visualize spectrogram: %v\n", err.Error())
	}
}
