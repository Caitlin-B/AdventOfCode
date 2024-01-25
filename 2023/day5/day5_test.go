package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testSeeds = []string{"79", "14", "55", "13"}
var testMaps = []string{"", "seed-to-soil map:", "50 98 2", "52 50 48", "", "soil-to-fertilizer map:",
	"0 15 37", "37 52 2", "39 0 15", "", "fertilizer-to-water map:", "49 53 8", "0 11 42", "42 0 7", "57 7 4", "",
	"water-to-light map:", "88 18 7", "18 25 70", "", "light-to-temperature map:", "45 77 23", "81 45 19",
	"68 64 13", "", "temperature-to-humidity map:", "0 69 1", "1 0 69", "", "humidity-to-location map:", "60 56 37", "56 93 4"}

func TestLocationToSeedAndInRange(t *testing.T) {
	cases := []struct {
		name      string
		inputLoc  int
		assertion func(*testing.T, bool)
	}{{
		name:     "seed within range",
		inputLoc: 46,
		assertion: func(t *testing.T, got bool) {
			assert.True(t, got)
		},
	}, {
		name:     "seed within range",
		inputLoc: 82,
		assertion: func(t *testing.T, got bool) {
			assert.True(t, got)
		},
	}, {
		name:     "seed within range",
		inputLoc: 86,
		assertion: func(t *testing.T, got bool) {
			assert.True(t, got)
		},
	}, {
		name:     "seed not within range",
		inputLoc: 35,
		assertion: func(t *testing.T, got bool) {
			assert.False(t, got)
		},
	}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := getRanges(testSeeds)
			seed := locationToSeed(testMaps, tc.inputLoc)
			tc.assertion(t, isInRanges(seed, r))
		})
	}
}
