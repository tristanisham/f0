// Copyright 2024 Tristan Isham. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// f0 is an implementation of the CVM algorithm for the distinct elements problem.
//   - https://arxiv.org/abs/2301.10191
//
// The CVM algorithm is best suited for large datasets where sorting unique values would require
// more than available memory.
package f0

import (
	"math"
	"math/rand/v2"
)

// coin simulates a coin flip. Heads is true, tails is false.
func coin() bool {
	return rand.IntN(2) == 0
}

func randClean[T comparable](s map[T]struct{}) {
	for item := range s {

		if coin() {
			delete(s, item)
		}

	}

}

// Estimate uses the CVM algorithm to calculate the number of unique item in an array.
// e (epsilon) represents acceptable relative error
// d (delta) represents the probability of failure for the algorithm to produce an estimation within the specific
// e-bound.
func Estimate[T comparable](source []T, e, d float64) float64 {

	x := make(map[T]struct{})

	p := 1.0
	m := len(source)

	thresh := math.Ceil(12.0/e*e) * math.Log2(8.0*float64(m)/d)

	for _, item := range source {

		if p == 0.5 {
			if coin() {
				x[item] = struct{}{}
			}
		} else if p == 1 {
			if _, ok := x[item]; !ok {
				x[item] = struct{}{}
			}
		}

		if math.Abs(float64(len(x))) == thresh {
			randClean(x)
			p = p / 2

		}

	}

	return math.Abs(float64(len(x))) / p
}
