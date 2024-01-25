package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

var testRaces = []Race{{7, 9}, {15, 40}, {30, 200}}

func TestDay1(t *testing.T) {
    assert.Equal(t, 288, day1(testRaces))
}
