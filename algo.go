// Copyright 2024 Tristan Isham. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// cvm is an implementation of the CVM algorithm for the distinct elements problem.
// 	- https://arxiv.org/abs/2301.10191
// The CVM algorithm is best suited for large datasets where sorting unique values would require
// more than available memory. 
package cvm

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

// Count uses the CVM algorithm to calculate the number of unique item in an array.
// The accuracy of the CVM algorithm is dependent on the size of its memory. The larger the memory
// in comparison to the orignial slice size, the more accurate the estimation, with 100% accuracy when
// the size and memory are equal.
func Count[T comparable](source []T, memory int) float64 {
	mem := memory
	if memory == 0 {
		mem = 100
	}

	pen := make(map[T]struct{})
	var round int = 1

	for _, t := range source {
		if len(pen) < mem {
			if _, found := pen[t]; found {
				previousCoin := coin()
				for n := 0; n < round; n++ {
					if c := coin(); c != previousCoin {
						delete(pen, t)
						break
					} else {
						previousCoin = c
					}
				}

			} else {
				pen[t] = struct{}{}
			}
		} else {
			randClean(pen)
			round++
		}
	}

	probability := math.Pow(0.5, 6.0)
	return float64(len(pen)) / probability
}
