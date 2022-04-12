package soundbox

import (
	"github.com/fbiville/soundbox/pkg/units"
	"math"
	"time"
)

type Sound interface {
	Value(sampleRate, sampleIndex int) float64
	DurationInSeconds() int
}

func Note(duration time.Duration, frequency units.PerSecond) Sound {
	return &complexSound{
		sounds: []sound{{
			frequency:       frequency,
			amplitudeFactor: 1,
		}},
		duration: duration,
	}
}

type complexSound struct {
	sounds   []sound
	duration time.Duration
}

func Chord(duration time.Duration, mainFrequencies ...units.PerSecond) Sound {
	return &complexSound{
		sounds:   harmonics(mainFrequencies),
		duration: duration,
	}
}

func (c *complexSound) Value(sampleRate, sampleIndex int) float64 {
	sum := 0.0
	for _, sound := range c.sounds {
		sum += sineWavePoint(sampleRate, sampleIndex, sound)
	}
	return sum
}

func (c *complexSound) DurationInSeconds() int {
	return int(c.duration.Seconds())
}

func harmonics(frequencies []units.PerSecond) []sound {
	var allFrequencies []sound
	for _, fundamental := range frequencies {
		allFrequencies = append(allFrequencies, sound{
			frequency:       fundamental,
			amplitudeFactor: 1,
		})
		for i := 1; i <= 4; i++ {
			harmonic := fundamental * units.PerSecond(1+i)
			allFrequencies = append(allFrequencies, sound{
				frequency:       harmonic,
				amplitudeFactor: 1.0 / math.Pow(2, float64(i)),
			})
		}
	}
	return allFrequencies
}

func sineWavePoint(samplingCount int, sampleIndex int, sound sound) float64 {
	return sound.amplitudeFactor * math.Sin(sound.frequency*2*math.Pi*float64(sampleIndex)/float64(samplingCount))
}

type sound struct {
	frequency       units.PerSecond
	amplitudeFactor float64
}
