package server

import (
	"bytes"
	"fmt"
	"os/exec"
)


func ConvertWAVToFLAC(wavData []byte) ([]byte, error) {
    cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-c:a", "flac", "-f", "flac", "pipe:1")

    inputPipe, err := cmd.StdinPipe()
    if err != nil {
        return nil, fmt.Errorf("failed to create FFmpeg stdin pipe: %v", err)
    }

    var flacBuffer bytes.Buffer
    cmd.Stdout = &flacBuffer
    var stderr bytes.Buffer
    cmd.Stderr = &stderr

    if err := cmd.Start(); err != nil {
        return nil, fmt.Errorf("error starting FFmpeg: %v", err)
    }

    go func() {
        defer inputPipe.Close()
        _, writeErr := inputPipe.Write(wavData)
        if writeErr != nil {
            fmt.Printf("error writing to FFmpeg stdin: %v\n", writeErr)
        }
    }()

    if err := cmd.Wait(); err != nil {
        return nil, fmt.Errorf("FFmpeg conversion error: %v, %s", err, stderr.String())
    }

    return flacBuffer.Bytes(), nil
}
