package problem1

import "math/rand"

func GetRandName(weights map[string]int) string {
	totalWeight := 0
	for _, weight := range weights {
		totalWeight += weight
	}
	r := rand.Intn(totalWeight)
	accumulatedWeighted := 0
	for name, weight := range weights {
		accumulatedWeighted += weight
		if r < accumulatedWeighted {
			return name
		}
	}
	return ""
}
