package tests

import (
	"GoStream/internal/server"
	"bytes"
	"os"
	"testing"
)


func TestConvertWAVToFLAC(t *testing.T) {

	wavData, err := os.ReadFile("tests/testdata/music.wav")
	if err != nil {
		t.Fatalf("Failed to load test WAV file: %v", err)
	}

	flacData, err := server.ConvertWAVToFLAC(wavData)
	if err != nil {
		t.Errorf("Conversion failed: %v", err)
	}

	if len(flacData) == 0 {
		t.Errorf("FLAC data is empty after conversion")
	}

	expectedSignature := []byte{0x66, 0x4C, 0x61, 0x43}
	if !bytes.HasPrefix(flacData, expectedSignature) {
		t.Errorf("Data does not start with expected FLAC signature")
	}
}
