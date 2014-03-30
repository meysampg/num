package rand

import (
	"testing"
	
	"github.com/phil-mansfield/num/objects/geom"
)

func BenchmarkTauswortheNext(b *testing.B) {
	rand := NewTimeSeed(Tausworthe)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Uniform(0, 2)
	}
}

func BenchmarkGoRandNext(b *testing.B) {
	rand := NewTimeSeed(GoRand)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Uniform(0, 2)
	}
}

func BenchmarkXorshiftNext(b *testing.B) {
	rand := NewTimeSeed(Xorshift)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Uniform(0, 2)
	}
}

func BenchmarkMultiplyWithCarryNext(b *testing.B) {
	rand := NewTimeSeed(MultiplyWithCarry)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Uniform(0, 2)
	}
}

func BenchmarkGslRandNext(b *testing.B) {
	rand := NewTimeSeed(GslRand)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Uniform(0, 2)
	}
}

func BenchmarkTauswortheGaussian(b *testing.B) {
	rand := NewTimeSeed(Tausworthe)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = rand.Gaussian(2, 2)
	}
}

func BenchmarkSphereAt(b *testing.B) {
	rand := NewTimeSeed(Tausworthe)

	b.ReportAllocs()
	b.ResetTimer()

	target := make([]float64, 3)
	for i := 0; i < b.N; i++ {
		rand.SphereAt(2, target)
	}
}

func BenchmarkFiniteLineAt(b *testing.B) {
	rand := NewTimeSeed(Xorshift)

	b.ReportAllocs()
	b.ResetTimer()

	target := make([]float64, 3)
	anchor := []float64{1, 1, 1}
	norm := []float64{1, 2, 1}
	line := geom.NewFiniteLine(anchor, norm, 3)
	for i := 0; i < b.N; i++ {
		rand.FiniteLineAt(line, target)
	}
}


func TestFrequncyTausworthe(t *testing.T) {
	rand := NewTimeSeed(Tausworthe)
	chis := FrequencyTest(rand, 10, 1000 * 1000, 5)
	badCount := 0
	for _, chi := range chis {
		if chi < 0.05 || chi > 0.95 {
			badCount += 1
		}
	}
	if badCount >= 2 {
		t.Errorf("Got Chi^2 values of %v", chis)
	}
}

func randomEnough(chis []float64) bool {
	badCount := 0
	for _, chi := range chis {
		if chi < 0.05 || chi > 0.95 {
			badCount += 1
		}
	}
	return badCount < len(chis)
}

func TestFrequncyGoRand(t *testing.T) {
	rand := NewTimeSeed(GoRand)
	chis := FrequencyTest(rand, 12, 5 * 1000 * 1000, 4)
	if !randomEnough(chis) {
		t.Errorf("Got Chi^2 values of %v", chis)
	}
}

func TestFrequncyXorshift(t *testing.T) {
	rand := NewTimeSeed(GoRand)
	chis := FrequencyTest(rand, 12, 5 * 1000 * 1000, 4)
	if !randomEnough(chis) {
		t.Errorf("Got Chi^2 values of %v", chis)
	}
}

func TestFrequncyMultiplyWithCarry(t *testing.T) {
	rand := NewTimeSeed(GoRand)
	chis := FrequencyTest(rand, 12, 5 * 1000 * 1000, 4)
	if !randomEnough(chis) {
		t.Errorf("Got Chi^2 values of %v", chis)
	}
}

func TestSerialTausworthe(t *testing.T) {
	rand := NewTimeSeed(GoRand)
	chis := SerialTest(rand, 8, 5 * 1000 * 1000, 4)
	if !randomEnough(chis) {
		t.Errorf("Got Chi^2 values of %v", chis)
	}
}

func TestSerialGoRand(t *testing.T) {
	rand := NewTimeSeed(GoRand)
	chis := SerialTest(rand, 8, 5 * 1000 * 1000, 4)
	if !randomEnough(chis) {
		t.Errorf("Got Chi^2 values of %v", chis)
	}
}

func TestSerialXorshift(t *testing.T) {
	rand := NewTimeSeed(GoRand)
	chis := SerialTest(rand, 8, 5 * 1000 * 1000, 4)
	if !randomEnough(chis) {
		t.Errorf("Got Chi^2 values of %v", chis)
	}
}

func TestSerialMultiplyWithCarry(t *testing.T) {
	rand := NewTimeSeed(GoRand)
	chis := SerialTest(rand, 8, 5 * 1000 * 1000, 4)
	if !randomEnough(chis) {
		t.Errorf("Got Chi^2 values of %v", chis)
	}
}
