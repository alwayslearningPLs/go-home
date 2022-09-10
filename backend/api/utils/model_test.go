package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type orderByAllower struct{}

func (o orderByAllower) OrderByColumnsAllowed() map[string]any {
	return map[string]any{"first": struct{}{}, "second": struct{}{}}
}

type queryParserImpl struct {
	queryDb map[string][]string
}

func (q queryParserImpl) QueryArray(key string) []string {
	if v, ok := q.queryDb[key]; ok && len(v) > 0 {
		return v
	}
	return []string{}
}

func (q queryParserImpl) Query(key string) string {
	if v := q.QueryArray(key); len(v) > 0 {
		return v[0]
	}
	return ""
}

func (q queryParserImpl) DefaultQuery(key string, d string) string {
	if q.Query(key) == "" {
		return d
	}
	return key
}

func TestDirectionFromInt(t *testing.T) {
	for _, each := range []struct {
		description          string
		src                  int
		want                 Direction
		wantType, wantString string
		wantErr              error
	}{
		{
			description: "direction is asc",
			src:         0,
			want:        asc,
			wantType:    "int",
			wantString:  "ASC",
		},
		{
			description: "direction is desc",
			src:         1,
			want:        desc,
			wantType:    "int",
			wantString:  "DESC",
		},
		{
			description: "direction is unknown",
			src:         2,
			wantType:    "int",
			wantString:  "ASC",
			wantErr:     errInvalidDirection,
		},
	} {
		t.Run(each.description, func(t *testing.T) {
			var d Direction

			gotErr := d.Set(each.src)

			assert.ErrorIs(t, each.wantErr, gotErr)
			assert.Equal(t, each.want, d)
			assert.Equal(t, each.wantString, d.String())
			assert.Equal(t, each.wantType, d.Type())
		})
	}
}

func TestDirectionFromText(t *testing.T) {
	for _, each := range []struct {
		description          string
		src                  string
		want                 Direction
		wantType, wantString string
		wantB                bool
	}{
		{
			description: "direction is asc",
			src:         "asc",
			want:        asc,
			wantType:    "int",
			wantString:  "ASC",
			wantB:       true,
		},
		{
			description: "direction is desc",
			src:         "desc",
			want:        desc,
			wantType:    "int",
			wantString:  "DESC",
			wantB:       true,
		},
		{
			description: "direction is unknown",
			src:         "unknown",
			wantType:    "int",
			wantString:  "ASC",
			wantB:       false,
		},
	} {
		t.Run(each.description, func(t *testing.T) {
			var d Direction

			gotB := d.unmarshalText(each.src)

			assert.Equal(t, each.wantB, gotB)
			assert.Equal(t, each.want, d)
			assert.Equal(t, each.wantString, d.String())
			assert.Equal(t, each.wantType, d.Type())
		})
	}
}

func TestParseOrderBy(t *testing.T) {
	for _, each := range []struct {
		description string
		input       []string
		want        []orderBy
	}{
		{
			description: "no value passed",
			input:       []string{},
			want:        []orderBy{},
		},
		{
			description: "Just one value correct from one len arr",
			input:       []string{"field_1 desc"},
			want:        []orderBy{{Field: "field_1", Direction: desc}},
		},
		{
			description: "All values are correct",
			input:       []string{"field_1 desc", "field_2 asc", "field_3 desc"},
			want:        []orderBy{{Field: "field_1", Direction: desc}, {Field: "field_2", Direction: asc}, {Field: "field_3", Direction: desc}},
		},
		{
			description: "Just one value incorrect from one len arr",
			input:       []string{"field_1  desc"},
			want:        []orderBy{},
		},
		{
			description: "All values are incorrect",
			input:       []string{"field_1  desc", "field_2 a sc", "field_3 d e s c"},
			want:        []orderBy{},
		},
		{
			description: "Some value are correct and others not",
			input:       []string{"field_1 desc", "field_2 a sc", "field_3 desc"},
			want:        []orderBy{{Field: "field_1", Direction: desc}, {Field: "field_3", Direction: desc}},
		},
	} {
		t.Run(each.description, func(t *testing.T) {
			got := parseArrOrderBy(each.input)

			assert.ElementsMatch(t, each.want, got)
		})
	}
}

func TestParseRequest(t *testing.T) {
	for _, each := range []struct {
		description string
		input       QueryParser
		want        WrapperRequest[orderByAllower]
	}{
		{
			description: "limit query param present and orderBy",
			input: queryParserImpl{queryDb: map[string][]string{
				limitQuery:   {"20"},
				orderByQuery: {"field_1 asc"},
			}},
			want: WrapperRequest[orderByAllower]{
				Limit:   20,
				OrderBy: []orderBy{{Field: "field_1", Direction: asc}},
				Body:    orderByAllower{},
			},
		},
		{
			description: "limit query param is not present and orderBy is",
			input: queryParserImpl{queryDb: map[string][]string{
				orderByQuery: {"field_1 asc"},
			}},
			want: WrapperRequest[orderByAllower]{
				Limit:   50,
				OrderBy: []orderBy{{Field: "field_1", Direction: asc}},
				Body:    orderByAllower{},
			},
		},
		{
			description: "limit query param is present, but it is lower than min allowed",
			input: queryParserImpl{queryDb: map[string][]string{
				limitQuery:   {"0"},
				orderByQuery: {"field_1 asc"},
			}},
			want: WrapperRequest[orderByAllower]{
				Limit:   50,
				OrderBy: []orderBy{{Field: "field_1", Direction: asc}},
				Body:    orderByAllower{},
			},
		},
		{
			description: "limit query param is present, but it is greater than max allowed",
			input: queryParserImpl{queryDb: map[string][]string{
				limitQuery:   {"101"},
				orderByQuery: {"field_1 asc"},
			}},
			want: WrapperRequest[orderByAllower]{
				Limit:   50,
				OrderBy: []orderBy{{Field: "field_1", Direction: asc}},
				Body:    orderByAllower{},
			},
		},
	} {
		t.Run(each.description, func(t *testing.T) {
			got := ParseRequest(each.input, orderByAllower{})

			assert.Equal(t, each.want, got)
		})
	}
}
