package wav

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
)

type WavHeader struct {
	// Master RIFF chunk
	FileTypeBlocID [4]byte
	FileSize       uint32
	FileFormatID   [4]byte

	// Chunk describing the data format
	FormatBlocID  [4]byte
	BlocSize      uint32
	AudioFormat   uint16
	NumChannels   uint16
	Frequency     uint32
	BytesPerSec   uint32
	BytePerBloc   uint16
	BitsPerSample uint16

	// Chunk containing the sampled data
	DataBlocID [4]byte
	DataSize   uint32
}

type Wav struct {
	Channels   int
	SampleRate int
	// audio duartion in seconds
	Duration float64
	Data     []byte
}

const WavHeaderLength = 44

// extracts metadata from the target file, assuming this path points to a WAV formatted sound file
func ReadWav(filename string) (*Wav, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(data) < WavHeaderLength {
		return nil, errors.New("invalid wav file size")
	}
	header := WavHeader{}
	err = binary.Read(bytes.NewReader(data[:WavHeaderLength]), binary.LittleEndian, &header)
	if err != nil {
		return nil, err
	}
	if string(header.FileTypeBlocID[:]) != "RIFF" || string(header.FileFormatID[:]) != "WAVE" || header.AudioFormat != 1 {
		return nil, errors.New("invalid wav header")
	}
	info := &Wav{
		Channels:   int(header.NumChannels),
		SampleRate: int(header.Frequency),
		Data:       data[WavHeaderLength:],
	}
	if header.BitsPerSample == 16 {
		info.Duration = float64(len(info.Data)) / float64(uint32(header.NumChannels)*2*header.Frequency)
	} else {
		return nil, errors.New("unsupported bits per sample format")
	}
	return info, err
}

// FFmpegMetadata represents the metadata structure returned by ffprobe.
type FFmpegMetadata struct {
	Streams []struct {
		Index         int               `json:"index"`
		CodecName     string            `json:"codec_name"`
		CodecLongName string            `json:"codec_long_name"`
		CodecType     string            `json:"codec_type"`
		SampleFmt     string            `json:"sample_fmt"`
		SampleRate    string            `json:"sample_rate"`
		Channels      int               `json:"channels"`
		ChannelLayout string            `json:"channel_layout"`
		BitsPerSample int               `json:"bits_per_sample"`
		Duration      string            `json:"duration"`
		BitRate       string            `json:"bit_rate"`
		Disposition   map[string]int    `json:"disposition"`
		Tags          map[string]string `json:"tags"`
	} `json:"streams"`
	Format struct {
		Streams        int               `json:"nb_streams"`
		FormFilename   string            `json:"filename"`
		NbatName       string            `json:"format_name"`
		FormatLongName string            `json:"format_long_name"`
		StartTime      string            `json:"start_time"`
		Duration       string            `json:"duration"`
		Size           string            `json:"size"`
		BitRate        string            `json:"bit_rate"`
		Tags           map[string]string `json:"tags"`
	} `json:"format"`
}

// GetMetadata retrieves metadata from a file using ffprobe.
func GetMetadata(filePath string) (FFmpegMetadata, error) {
	var metadata FFmpegMetadata

	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", filePath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return metadata, err
	}

	err = json.Unmarshal(out.Bytes(), &metadata)
	if err != nil {
		return metadata, err
	}

	return metadata, nil
}

func BytesToSamples(data []byte) ([]float64, error) {
	if len(data)%2 != 0 {
		return nil, errors.New("invalid input length")
	}

	numSamples := len(data) / 2
	output := make([]float64, numSamples)

	for i := 0; i < len(data); i += 2 {
		// Interpret bytes as a 16-bit signed integer (little-endian)
		sample := int16(binary.LittleEndian.Uint16(data[i : i+2]))

		// Scale the sample to the range [-1, 1]
		output[i/2] = float64(sample) / 32768.0
	}

	return output, nil
}
