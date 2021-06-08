package urls

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestValidateURLItem struct {
	input string
	expected bool
}

type TestValidateSlugItem struct {
	input string
	expected bool
}

func TestValidateURL(t *testing.T) {
	testCases := []TestValidateURLItem {
		{`https://medium.com`, true},
		{`medium.com`, false},
		{`aaa`, false},
		{`https://stackoverflow.com/questions/45267125/how-to-generate-unique-random-alphanumeric-tokens-in-golang`, true},
	}

	for _, c := range testCases {
		result := ValidateURL(c.input)
		assert.Equal(t, c.expected, result, fmt.Sprintf("URL [%v] -> Should be [%v]", c.input, c.expected))
	}
}

func TestValidateSlug(t *testing.T) {
	testCases := []TestValidateSlugItem {
		{`a`, false},
		{`abcdeefghahdgasfkhvlkcvjxch`, false},
		{`alive`, false},
		{`register-url`, false},
		{`correct`, true},
		{`cj.asd`, false},
		{`cj^asd`, false},
		{`cj:asd`, false},
		{`hi-there`, true},
		{`h0123ere`, true},
	}

	for _, c := range testCases {
		result := ValidateSlug(c.input)
		assert.Equal(t, c.expected, result, fmt.Sprintf("Slug [%v] -> Should be [%v]", c.input, c.expected))
	}
}