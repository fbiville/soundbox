package soundbox

import (
	"fmt"
	"os/exec"
	"time"
)

type player struct {
	sampleRate int
}

func NewDefaultPlayer() *player {
	return &player{sampleRate: defaultSampleRate}
}

func (p *player) PlayF32LE(file string, duration time.Duration) {
	command := exec.Command("ffplay",
		"-f", "f32le", // f32le: PCM 32-bit floating-point little-endian
		"-t", fmt.Sprintf("%d", int(duration.Seconds())), // f32le: PCM 32-bit floating-point little-endian
		"-ar", fmt.Sprintf("%d", p.sampleRate*2), // FIXME: not sure why factor fixes the weird pitch division
		"-autoexit",
		"-nodisp",
		file)
	if err := command.Run(); err != nil {
		panic(err)
	}
}
