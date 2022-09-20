package main

import "testing"

func TestWrapValue(t *testing.T) {
	width := 10

	testCases := [][2]int{
		{0, 0},
		{5, 5},
		{10, 0},
		{11, 1},
		{22, 2},
		{-1, 9},
		{-2, 8},
		{-3, 7},
		{-4, 6},
		{-5, 5},
		{-10, 0},
		{-20, 0},
	}

	for _, tc := range testCases {
		val, expected := tc[0], tc[1]
		actual := wrapValue(val, width)
		if actual != expected {
			t.Errorf("expected %d, got %d", expected, actual)
		}
	}
}
