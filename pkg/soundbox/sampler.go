package soundbox

import (
	"encoding/binary"
	"math"
	"os"
)

const defaultSampleRate = 44100

type sampler struct {
	sampleRate int
}

func NewDefaultSampler() *sampler {
	return &sampler{sampleRate: defaultSampleRate}
}

func (s *sampler) GenerateF32LE(sounds ...Sound) string {
	file, _ := os.CreateTemp("", "sound.raw")
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	for _, sound := range sounds {
		for t := 0; t < sound.DurationInSeconds(); t++ {
			for i := 0; i < s.sampleRate; i++ {
				sample := sound.Value(s.sampleRate, i)
				var buf [8]byte
				binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
				if _, err := file.Write(buf[:]); err != nil {
					panic(err)
				}
			}
		}
	}
	return file.Name()
}
