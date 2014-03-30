package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/phil-mansfield/num/rand"
	"github.com/phil-mansfield/objects/geom"
	"github.com/phil-mansfield/objects/vec"
)

func cosSqr(x) {
	cosX := math.Cos(x)
	return cosX * cosX
}

func main() {
	if len(os.Argv) != 2 {
		panic("")
	}

	trials := strconv.Atoi(os.Argv[1])
	gen := rand.NewTimeSeed()

	for i := 0; i < trials; i++ {
		theta := rand.MonteCarlo(cosSqr, 0, 2 * math.Pi, 0, 1)
	}
}
