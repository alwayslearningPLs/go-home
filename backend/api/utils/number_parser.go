package utils

import (
	"strconv"
)

type Number interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | int | uint
}

type Filter[T Number] func(input T) bool

func ParseNumber[T Number](input string, d T, filter ...Filter[T]) T {
	i, err := strconv.Atoi(input)
	if err != nil {
		return d
	}

	result := T(i)

	for _, f := range filter {
		if !f(result) {
			return d
		}
	}

	return result
}

func Boundaries[T Number](min, max T) Filter[T] {
	return func(input T) bool {
		return min <= input && input <= max
	}
}
