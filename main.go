package main

import (
	"log"

	"delm.dev/music-recognizer/transform"
	"delm.dev/music-recognizer/wav"
)

func main() {
	info, err := wav.ReadWavInfo("./assets/saul.wav")
	if err != nil {
		log.Fatalf("Unable to read info: %s\n", err.Error())
	}
	samples, _ := wav.BytesToSamples(info.Data)
	spectrogram, _ := transform.Spectrogram(samples, info.SampleRate)
	log.Printf("length: %d\n", len(spectrogram))
}
