package randomizer

import (
	"math"
	"math/rand"
	"time"
)

type Randomizer struct {
	maxChance float64
}

func (r *Randomizer) Create() *Randomizer {
	r = &Randomizer{
		maxChance: 100,
	}

	return r
}

func (r *Randomizer) IsChance(chance int) bool {
	rand.Seed(time.Now().UTC().UnixNano())
	val := rand.Intn(int(math.Round(r.maxChance))) + 1

	if val <= chance {
		return true
	}

	return false
}

func (r *Randomizer) GetRandSpread(spread int) int {
	if spread < 1 {
		return 0
	}

	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(2*spread/2) - spread/2
}
