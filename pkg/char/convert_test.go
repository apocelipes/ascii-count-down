package char

import (
	"slices"
	"testing"
	"time"
)

func charsEqual(a, b [][]string) bool {
	return slices.EqualFunc(a, b, func(i, j []string) bool {
		return slices.Equal(i, j)
	})
}

func TestConvertToChars(t *testing.T) {
	testCases := []struct {
		duration time.Duration
		result   [][]string
	}{
		{
			duration: 0,
			result: [][]string{
				SharpChars[0],
				SharpChars[0],
				SharpChars[0],
				SharpChars[0],
				SharpChars[0],
				SharpChars[0],
			},
		},
		{
			duration: 1 * time.Hour,
			result: [][]string{
				SharpChars[0],
				SharpChars[1],
				SharpChars[0],
				SharpChars[0],
				SharpChars[0],
				SharpChars[0],
			},
		},
		{
			duration: 24 * time.Hour,
			result: [][]string{
				SharpChars[2],
				SharpChars[4],
				SharpChars[0],
				SharpChars[0],
				SharpChars[0],
				SharpChars[0],
			},
		},
		{
			duration: 1*time.Hour + 1*time.Minute + 1*time.Second,
			result: [][]string{
				SharpChars[0],
				SharpChars[1],
				SharpChars[0],
				SharpChars[1],
				SharpChars[0],
				SharpChars[1],
			},
		},
		{
			duration: 1*time.Second + 900*time.Millisecond,
			result: [][]string{
				SharpChars[0],
				SharpChars[0],
				SharpChars[0],
				SharpChars[0],
				SharpChars[0],
				SharpChars[1],
			},
		},
	}
	for _, tc := range testCases {
		res := ConvertToChars(tc.duration, SharpChars, nil)
		if !charsEqual(res, tc.result) {
			t.Errorf("%s convert failed, got: %v, want: %v", tc.duration, res, tc.result)
		}
	}
}

func TestConvertToCharsPanic(t *testing.T) {
	mustPanic := func(t *testing.T, fn func()) {
		t.Helper()
		defer func() {
			if err := recover(); err == nil {
				t.Error("want panic")
			}
		}()
		fn()
	}
	mustPanic(t, func() {
		ConvertToChars(0, nil, nil)
	})
	mustPanic(t, func() {
		ConvertToChars(100*time.Hour, SharpChars, nil)
	})
}
