package main_test

import (
	"strings"
	"testing"

	counter "github.com/baizhigit/go-cli-apps/counter"
)

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "5 words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "empty input",
			input: "",
			wants: 0,
		},
		{
			name:  "single space",
			input: " ",
			wants: 0,
		},
		{
			name:  "new lines",
			input: "one two three\nfour five",
			wants: 5,
		},
		{
			name:  "multiple spaces",
			input: "This is a sentence.  This is another",
			wants: 7,
		},
		{
			name:  "Prefixed multiple spaces",
			input: "  Hello",
			wants: 1,
		},
		{
			name:  "Suffixed multiple spaces",
			input: "Hello     ",
			wants: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			result := counter.CountWords(r)
			if result != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, result)
				t.Fail()
			}
		})
	}
}

func TestCountLines(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "simple five words, 1 new line",
			input: "one two three four five\n",
			wants: 1,
		},
		{
			name:  "empty file",
			input: "",
			wants: 0,
		},
		{
			name:  "no new line",
			input: "one two three four five",
			wants: 0,
		},
		{
			name:  "no new line at end",
			input: "one two three four five\nsix",
			wants: 1,
		},
		{
			name:  "multi newline string",
			input: "\n\n\n\n",
			wants: 4,
		},
		{
			name:  "multi word line string",
			input: "one\ntwo\nthree\nfour\nfive\n",
			wants: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			res := counter.Countlines(r)

			if res != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, res)
				t.Fail()
			}
		})
	}
}

func TestCountBytes(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "five words",
			input: "one two three four five",
			wants: 23,
		},
		{
			name:  "empty file",
			input: "",
			wants: 0,
		},
		{
			name:  "all spaces",
			input: "       ",
			wants: 7,
		},
		{
			name:  "newLines and words",
			input: "one\ntwo\nthree\nfour\nfive\t\n",
			wants: 25,
		},
		{
			name:  "unicode chars",
			input: "Җ",
			wants: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)

			res := counter.CountBytes(r)

			if res != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, res)
				t.Fail()
			}
		})
	}
}
