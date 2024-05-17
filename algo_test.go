package f0

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

// This test should work because it's memory is significantly higher than the unique count.
// TODO: I need to tweak the algorithm so that this works.
func TestBible(t *testing.T) {

	source, err := os.ReadFile("bible.txt")
	if err != nil {
		t.Error(err)
	}

	vx := strings.ReplaceAll(string(source), ".", "")
	ax := strings.Split(vx, " ")

	mapMethod := time.Now()

	og := make(map[string]struct{})
	for _, word := range ax {
		if _, ok := og[word]; !ok {
			og[word] = struct{}{}
		}
	}

	fmt.Println("Map time:", time.Since(mapMethod))

	uniqueCount := float64(len(og))

	// At this scale, with such small memory, inaccuracies are expected.

	estTime := time.Now()

	if n := Estimate(ax, 0.01, 0.001); n != uniqueCount {
		t.Errorf("wanted %f, got %f", uniqueCount, n)
	}

	fmt.Println("Estimate time:", time.Since(estTime))

}
