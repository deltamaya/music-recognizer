package wav

import (
	"testing"
)

// check the if data length equals bytesPerSample * channels * sampleRate * duration
func TestLength(t *testing.T) {
	wavData, err := ReadWav("../assets/saul.wav")
	if err != nil {
		t.Error(err)
	}
	bodyLength := len(wavData.Data)
	calcLength := 2.0 * float64(wavData.Channels) * float64(wavData.SampleRate) * wavData.Duration
	if bodyLength != int(calcLength) {
		t.Errorf("length does not match: body: %d, calc: %f\n", bodyLength, calcLength)
	}
}
