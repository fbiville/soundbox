package main

import (
	. "github.com/fbiville/soundbox/pkg/soundbox"
	. "github.com/fbiville/soundbox/pkg/units"
	. "time"
)

const (
	d2      = 73.4 * Hertz
	e2      = 82.41 * Hertz
	g2      = 98 * Hertz
	a2      = 110 * Hertz
	c3      = 130.81 * Hertz
	b3      = 246.94 * Hertz
	csharp4 = 277.18 * Hertz
	d4      = 293.66 * Hertz
	e4      = 329.63 * Hertz
	fsharp4 = 369.99 * Hertz
	g4      = 392 * Hertz
	a4      = 440 * Hertz
)

// go run ./cmd/02-complex
func main() {
	niceMelody()
	Sleep(2 * Second)
	uglyChords()
}

func niceMelody() {
	rawFile := NewDefaultSampler().GenerateF32LE(
		Note(1*Second, d4),
		Note(1*Second, g4),
		Note(1*Second, fsharp4),
		Note(1*Second, g4),
		Note(1*Second, a4),
		Note(1*Second, e4),
		Note(2*Second, a4),
	)
	NewDefaultPlayer().PlayF32LE(rawFile, 8*Second)
}

func uglyChords() {
	rawFile := NewDefaultSampler().GenerateF32LE(
		Chord(1*Second, d4, b3, g2),
		Note(1*Second, g4),
		Chord(1*Second, fsharp4, b3, d2),
		Note(1*Second, g4),
		Chord(1*Second, a4, csharp4, a2),
		Note(1*Second, e4),
		Chord(2*Second, a4, c3, e2),
	)
	NewDefaultPlayer().PlayF32LE(rawFile, 8*Second)
}
