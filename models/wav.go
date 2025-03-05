package models

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
	"os"
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

type WavInfo struct {
	Channels   int
	SampleRate int
	Data       []byte
	Duration   float64
}

const WavHeaderLength = 44

// extracts metadata from the target file, assuming this path points to a WAV formatted sound file
func ReadWavInfo(filename string) (*WavInfo, error) {
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
	log.Printf("wav header: %v\n", header)
	if string(header.FileTypeBlocID[:]) != "RIFF" || string(header.FileFormatID[:]) != "WAVE" || header.AudioFormat != 1 {

		return nil, errors.New("invalid wav header")
	}
	info := &WavInfo{
		Channels:   int(header.NumChannels),
		SampleRate: int(header.Frequency),
		Data:       data[WavHeaderLength:],
	}
	if header.BitsPerSample == 16 {
		info.Duration = float64(len(info.Data)) / float64(int(header.NumChannels)) * 2.0 * float64(header.Frequency)
	} else {
		return nil, errors.New("unsupported bits per sample format")
	}
	return info, err
}
