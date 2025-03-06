package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalc(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		input := ""
		res := calc(input)

		if !assert.Equal(t, res, 0) {
			t.Fatalf("expexted %d, got=%d", 0, res)
		}
	})

	t.Run("single number", func(t *testing.T) {
		inputs := []struct {
			input    string
			expected int
		}{
			{
				input:    "1",
				expected: 1,
			},
			{
				input:    "5",
				expected: 5,
			},
		}

		for _, tt := range inputs {
			res := calc(tt.input)
			if !assert.Equal(t, tt.expected, res) {
				t.Fatalf("expected %d, got %d", tt.expected, res)
			}
		}
	})

	t.Run("two numbers comma delimered return sum", func(t *testing.T) {
		inputs := []struct {
			input    string
			expected int
		}{
			{
				input:    "1,2",
				expected: 3,
			},
			{
				input:    "1,5",
				expected: 6,
			},
		}

		for _, tt := range inputs {
			res := calc(tt.input)
			if !assert.Equal(t, tt.expected, res) {
				t.Fatalf("expected %d, got %d", tt.expected, res)
			}
		}
	})

	t.Run("two numbers, newline delimed", func(t *testing.T) {
		inputs := []struct {
			input    string
			expected int
		}{
			{
				input:    "1\n2",
				expected: 3,
			},
			{
				input:    "1\n5",
				expected: 6,
			},
		}

		for _, tt := range inputs {
			res := calc(tt.input)
			if !assert.Equal(t, tt.expected, res) {
				t.Fatalf("expected %d, got %d", tt.expected, res)
			}
		}
	})

	t.Run("three numbers, delimed , or new line", func(t *testing.T) {
		inputs := []struct {
			input    string
			expected int
		}{
			{
				input:    "1,2,3",
				expected: 6,
			},
			{
				input:    "1\n5\n1",
				expected: 7,
			},
			{
				input:    "1,5\n1",
				expected: 7,
			},
		}

		for _, tt := range inputs {
			res := calc(tt.input)
			if !assert.Equal(t, tt.expected, res) {
				t.Fatalf("expected %d, got %d", tt.expected, res)
			}
		}
	})

	t.Run("negative numbers retun error", func(t *testing.T) {
		inputs := []struct {
			input    string
			expected int
			panics   bool
		}{
			{
				input:    "-1,2,3",
				expected: -1,
				panics:   true,
			},
			{
				input:    "1\n5\n1",
				expected: 7,
				panics:   false,
			},
			{
				input:    "1,5\n1",
				expected: 7,
				panics:   false,
			},
		}

		for _, tt := range inputs {
			if tt.panics {
				if !assert.Panics(t, func() {
					calc(tt.input)
				}) {
					t.Fatalf("expected to panic but didn't")
				}
				return
			}
			res := calc(tt.input)
			if !assert.Equal(t, tt.expected, res) {
				t.Fatalf("expected %d got %d", tt.expected, res)
			}
		}

	})

	t.Run("numbers greate than 1000 are ignored", func(t *testing.T) {
		inputs := []struct {
			input    string
			expected int
			panics   bool
		}{
			{
				input:    "-1,2,3",
				expected: -1,
				panics:   true,
			},
			{
				input:    "1\n5\n10000",
				expected: 6,
				panics:   false,
			},
			{
				input:    "1,5\n1",
				expected: 7,
				panics:   false,
			},
		}

		for _, tt := range inputs {
			if tt.panics {
				if !assert.Panics(t, func() {
					calc(tt.input)
				}) {
					t.Fatalf("expected to panic but didn't")
				}
				return
			}
			res := calc(tt.input)
			if !assert.Equal(t, tt.expected, res) {
				t.Fatalf("expected %d got %d", tt.expected, res)
			}
		}

	})

	t.Run("a single delimeter can be defined", func(t *testing.T) {
		inputs := []struct {
			input    string
			expected int
			panics   bool
		}{
			{
				input:    "-1,2,3",
				expected: -1,
				panics:   true,
			},
			{
				input:    "1\n5\n10000",
				expected: 6,
				panics:   false,
			},
			{
				input:    "//#1#5#1",
				expected: 7,
				panics:   false,
			},
		}

		for _, tt := range inputs {
			if tt.panics {
				if !assert.Panics(t, func() {
					calc(tt.input)
				}) {
					t.Fatalf("expected to panic but didn't")
				}
				return
			}
			res := calc(tt.input)
			if !assert.Equal(t, tt.expected, res) {
				t.Fatalf("expected %d got %d", tt.expected, res)
			}
		}
	})

	t.Run("multi char delimeter", func(t *testing.T) {
		inputs := []struct {
			input    string
			expected int
			panics   bool
		}{
			{
				input:    "-1,2,3",
				expected: -1,
				panics:   true,
			},
			{
				input:    "1\n5\n10000",
				expected: 6,
				panics:   false,
			},
			{
				input:    "//#1#5#1",
				expected: 7,
				panics:   false,
			},
			{
				input:    "//#1#5#1",
				expected: 7,
				panics:   false,
			},
		}

		for _, tt := range inputs {
			if tt.panics {
				if !assert.Panics(t, func() {
					calc(tt.input)
				}) {
					t.Fatalf("expected to panic but didn't")
				}
				return
			}
			res := calc(tt.input)
			if !assert.Equal(t, tt.expected, res) {
				t.Fatalf("expected %d got %d", tt.expected, res)
			}
		}
	})

}
