package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoundaries(t *testing.T) {
	for _, each := range []struct {
		description, input string
		filter             Filter[int8]
		d, want            int8
	}{
		{
			description: "int8 being greater than max, so default is applied",
			input:       "120",
			filter:      Boundaries[int8](1, 100),
			d:           50,
			want:        50,
		},
		{
			description: "int8 being lower than min, so default is applied",
			input:       "0",
			filter:      Boundaries[int8](1, 100),
			d:           50,
			want:        50,
		},
		{
			description: "int8 between boundaries",
			input:       "30",
			filter:      Boundaries[int8](1, 100),
			d:           50,
			want:        30,
		},
	} {
		t.Run(each.description, func(t *testing.T) {
			got := ParseNumber(each.input, each.d, each.filter)

			assert.Equal(t, got, each.want)
		})
	}
}

func TestParseNumber(t *testing.T) {
	t.Run("parse number with invalid input", func(t *testing.T) {
		var (
			input = "hello world"

			want = 50
		)

		got := ParseNumber(input, want)

		assert.Equal(t, want, got)
	})
}
