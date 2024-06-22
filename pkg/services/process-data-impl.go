package services

import (
	"math"

	"github.com/sdblg/vrp/pkg/models"
)

func processData(loads []*models.Load) error {
	stop := false
	for !stop {
		updating := false
		for i := 0; i < len(loads); i++ {
			minCost := math.MaxFloat64
			var t int
			for j := i + 1; j < len(loads); j++ {
				if loads[j].Joined {
					continue
				}
				lCost := loads[i].Join(loads[j], false)
				minCost = min(minCost, lCost)
				if minCost > 12*60 {
					break
				}
				t = j
			}
			if t != 0 {
				loads[i].Join(loads[t], true) // joining minCost road
				updating = true
			}
		}
		if !updating {
			stop = true
		}
	}

	return nil
}
