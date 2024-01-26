package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testOne = []string{"RL", "", "AAA = (BBB, CCC)", "BBB = (DDD, EEE)", "CCC = (ZZZ, GGG)", "DDD = (DDD, DDD)", "EEE = (EEE, EEE)", "GGG = (GGG, GGG)", "ZZZ = (ZZZ, ZZZ)"}
var testTwo = []string{"LLR", "", "AAA = (BBB, BBB)", "BBB = (AAA, ZZZ)", "ZZZ = (ZZZ, ZZZ)"}

func TestDay1(t *testing.T) {
	assert.Equal(t, 2, day1(testOne))
	assert.Equal(t, 6, day1(testTwo))
}
