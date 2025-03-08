package main

import (
	"log"

	"delm.dev/music-recognizer/shazam"
	"delm.dev/music-recognizer/wav"
)

func main() {

	// wav.ConvertToWav("./assets/Euphoria.mp3", 1)

	// data, _ := wav.ReadWav("./assets/Euphoria.wav")
	// db, _ := db.NewDBClient()
	// id, err := db.RegisterSong("Euphoria", "ar", "ytID et")
	// if err != nil {
	// 	log.Printf("unable to register song: %s\n", err.Error())
	// }
	// log.Printf("song ID: %d\n", id)
	// samples, _ := wav.BytesToSamples(data.Data)
	// spectrogram, _ := transform.Spectrogram(samples, data.SampleRate)
	// peaks := transform.ExtractPeaks(spectrogram, data.Duration)
	// fps := shazam.Fingerprint(peaks, id)
	// db.StoreFingerprints(fps)

	data, _ := wav.ReadWav("./assets/girl5s.wav")
	samples, _ := wav.BytesToSamples(data.Data)
	matches, _, _ := shazam.FindMatches(samples, data.Duration, data.SampleRate)
	for _, match := range matches {
		log.Printf("%s: %f\n", match.SongTitle, match.Score)
	}

}
