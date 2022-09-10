package utils

import (
	"encoding/xml"
	"errors"
	"math"
	"reflect"
	"strings"
)

var (
	errInvalidDirection = errors.New("invalid direction for the value supplied")
	errInvalidOrderBy   = errors.New("invalid order by for the value supplied")
)

const (
	limitQuery   = "limit"
	skipQuery    = "skip"
	orderByQuery = "order_by"

	limitDefault = 50
	limitMax     = 100
	limitMin     = 1

	skipDefault = 0
	skipMax     = math.MaxInt32
	skipMin     = 0
)

type Direction int

const (
	asc Direction = iota
	desc
)

// Type returns the string representation of the inner type
// of direction
func (d Direction) Type() string {
	return reflect.Int.String()
}

// Set is used to set the value to a direction type, checking
// if the input is inside the boundaries. If the value is not correct,
// by default, will be assigned asc
func (d *Direction) Set(src int) error {
	if !d.unmarshal(src) {
		return errInvalidDirection
	}
	return nil
}

func (d *Direction) unmarshal(src int) bool {
	switch src {
	case 0:
		*d = asc
	case 1:
		*d = desc
	default:
		return false
	}
	return true
}

func (d *Direction) unmarshalText(src string) bool {
	switch src {
	case "asc", "ASC":
		*d = asc
	case "desc", "DESC":
		*d = desc
	default:
		return false
	}
	return true
}

// String returns the string representation of the type direction
func (d Direction) String() string {
	var result string

	switch d {
	case asc:
		result = "ASC"
	case desc:
		result = "DESC"
	}

	return result
}

type WrapperRequest[T OrderByAllower] struct {
	Limit   int
	Skip    int
	OrderBy []orderBy
	Body    T
}

type orderBy struct {
	Field     string
	Direction Direction
}

func (o orderBy) String() string {
	return o.Field + " " + o.Direction.String()
}

type WrapperResponse struct {
	XMLName xml.Name `json:"-" xml:"Response"`
	Code    int      `json:"code" xml:"Code"`
	Msg     string   `json:"message" xml:"Message"`
}

func ParseRequest[T OrderByAllower](qParser QueryParser, body T) WrapperRequest[T] {
	return WrapperRequest[T]{
		Limit:   ParseNumber(qParser.Query(limitQuery), limitDefault, Boundaries(limitMin, limitMax)),
		Skip:    ParseNumber(qParser.Query(skipQuery), skipDefault, Boundaries(skipMin, skipMax)),
		OrderBy: parseArrOrderBy(qParser.QueryArray(orderByQuery)),
		Body:    body,
	}
}

func parseArrOrderBy(arr []string) []orderBy {
	var (
		result = make([]orderBy, len(arr))
		j      = 0
	)

	for i := range arr {
		if o, err := parseOrderBy(arr[i]); err == nil {
			result[j] = o
			j++
		}
	}

	return result[:j]
}

func parseOrderBy(input string) (orderBy, error) {
	var (
		splitted = strings.Split(input, " ")

		d Direction
	)

	if len(splitted) != 2 {
		return orderBy{}, errInvalidOrderBy
	}

	d.unmarshalText(splitted[1])

	return orderBy{
		Field:     splitted[0],
		Direction: d,
	}, nil
}
