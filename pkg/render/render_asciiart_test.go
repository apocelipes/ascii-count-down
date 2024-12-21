package render

import (
	"testing"
	"time"

	"github.com/apocelipes/ascii-count-down/pkg/char"
)

func mustParseDuration(t *testing.T, d string) time.Duration {
	t.Helper()
	duration, err := time.ParseDuration(d)
	if err != nil {
		t.Fatal(err)
	}
	return duration
}

func TestASCIIArtTotalWidth(t *testing.T) {
	testCases := []struct {
		duration    time.Duration
		expectWidth int
	}{
		{
			duration:    0,
			expectWidth: char.MaxASCIIArtCharWidth(0)*6 + asciiArtNonDigitTotalWidth,
		},
		{
			duration:    mustParseDuration(t, "0h0m1s"),
			expectWidth: char.MaxASCIIArtCharWidth(0)*5 + char.MaxASCIIArtCharWidth(1) + asciiArtNonDigitTotalWidth,
		},
		{
			duration: mustParseDuration(t, "0h0m10s"),
			expectWidth: char.MaxASCIIArtCharWidth(0)*4 +
				char.MaxASCIIArtCharWidth(1) +
				char.MaxASCIIArtCharWidth(9) +
				asciiArtNonDigitTotalWidth,
		},
		{
			duration: mustParseDuration(t, "0h1m0s"),
			expectWidth: char.MaxASCIIArtCharWidth(0)*3 + char.MaxASCIIArtCharWidth(1) +
				char.MaxASCIIArtCharWidth(5) +
				char.MaxASCIIArtCharWidth(9) +
				asciiArtNonDigitTotalWidth,
		},
		{
			duration: mustParseDuration(t, "0h10m0s"),
			expectWidth: char.MaxASCIIArtCharWidth(0)*2 + char.MaxASCIIArtCharWidth(1) +
				char.MaxASCIIArtCharWidth(9) +
				char.MaxASCIIArtCharWidth(5) +
				char.MaxASCIIArtCharWidth(9) +
				asciiArtNonDigitTotalWidth,
		},
		{
			duration: mustParseDuration(t, "1h0m0s"),
			expectWidth: char.MaxASCIIArtCharWidth(0) +
				char.MaxASCIIArtCharWidth(1) +
				char.MaxASCIIArtCharWidth(5) +
				char.MaxASCIIArtCharWidth(9) +
				char.MaxASCIIArtCharWidth(5) +
				char.MaxASCIIArtCharWidth(9) +
				asciiArtNonDigitTotalWidth,
		},
		{
			duration: mustParseDuration(t, "10h0m0s"),
			expectWidth: char.MaxASCIIArtCharWidth(1) +
				char.MaxASCIIArtCharWidth(9) +
				char.MaxASCIIArtCharWidth(5) +
				char.MaxASCIIArtCharWidth(9) +
				char.MaxASCIIArtCharWidth(5) +
				char.MaxASCIIArtCharWidth(9) +
				asciiArtNonDigitTotalWidth,
		},
		{
			duration: mustParseDuration(t, "99h59m59s"),
			expectWidth: char.MaxASCIIArtCharWidth(9) +
				char.MaxASCIIArtCharWidth(9) +
				char.MaxASCIIArtCharWidth(5) +
				char.MaxASCIIArtCharWidth(9) +
				char.MaxASCIIArtCharWidth(5) +
				char.MaxASCIIArtCharWidth(9) +
				asciiArtNonDigitTotalWidth,
		},
		{
			duration: mustParseDuration(t, "90h0m0s"),
			expectWidth: char.MaxASCIIArtCharWidth(9) +
				char.MaxASCIIArtCharWidth(9) +
				char.MaxASCIIArtCharWidth(5) +
				char.MaxASCIIArtCharWidth(9) +
				char.MaxASCIIArtCharWidth(5) +
				char.MaxASCIIArtCharWidth(9) +
				asciiArtNonDigitTotalWidth,
		},
	}

	for _, tc := range testCases {
		if w := asciiArtTotalWidth(tc.duration); w != tc.expectWidth {
			t.Errorf("%v got %v, want %v", tc.duration, w, tc.expectWidth)
		}
	}
}
