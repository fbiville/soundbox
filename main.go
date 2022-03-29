package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
)

const (
	SoundDuration = 1
	SampleRate    = 44100

	Tau = math.Pi * 2
)

func main() {
	sampleCountPerFreq := SoundDuration * SampleRate
	file, err := os.CreateTemp("", "sound-")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	for i, frequency := range generateRandomFrequencies(15) {
		start := i * sampleCountPerFreq
		for sampleIndex := start; sampleIndex < start+sampleCountPerFreq; sampleIndex++ {
			sample := math.Sin(Tau / float64(sampleCountPerFreq) * frequency * float64(sampleIndex))
			var buffer [8]byte

			binary.LittleEndian.PutUint32(buffer[:], math.Float32bits(float32(sample)))
			if _, err := file.Write(buffer[:]); err != nil {
				panic(err)
			}
		}
	}

	command := exec.Command("ffplay",
		"-f", "f32le", // f32le: PCM 32-bit floating-point little-endian
		"-ar", fmt.Sprintf("%d", SampleRate),
		"-showmode", "1",
		file.Name())
	if err = command.Run(); err != nil {
		panic(err)
	}
}

func generateRandomFrequencies(count int) []float64 {
	result := make([]float64, count)
	for i := 0; i < count; i++ {
		result[i] = 20 + rand.Float64()*(20000-20)
	}
	return result
}
