package wav

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func ConvertToWav(inputFilepath string, channels int) (wavFilepath string, err error) {

	if channels < 1 || channels > 2 {
		channels = 1
	}
	fileExt := filepath.Ext(inputFilepath)
	outputFilepath := strings.TrimSuffix(inputFilepath, fileExt) + ".wav"
	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", inputFilepath,
		"-c", "pcm_s16le",
		"-ar", "44100",
		"-ac", strconv.FormatInt(int64(channels), 10),
		outputFilepath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to convert to WAV: %v, output %v", err, string(output))
	}

	return outputFilepath, nil
}

func ReformatWAV(inputFilePath string, channels int) (reformatedFilePath string, errr error) {
	if channels < 1 || channels > 2 {
		channels = 1
	}

	fileExt := filepath.Ext(inputFilePath)
	outputFile := strings.TrimSuffix(inputFilePath, fileExt) + "rfm.wav"

	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-i", inputFilePath,
		"-c", "pcm_s16le",
		"-ar", "44100",
		"-ac", strconv.FormatInt(int64(channels), 10),
		outputFile,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to convert to WAV: %v, output %v", err, string(output))
	}

	return outputFile, nil
}
