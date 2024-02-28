package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRule(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		r := ParseRule("\"Goodbye\" -> \"Hello\"")
		assert.Equal(t, "Goodbye", r.From)
		assert.Equal(t, "Hello", r.To)
	})
	t.Run("with whitespace", func(t *testing.T) {
		r := ParseRule("\" Goodbye  \"->  \" Hello \"")
		assert.Equal(t, " Goodbye  ", r.From)
		assert.Equal(t, " Hello ", r.To)
	})
}

func TestParseRuleset(t *testing.T) {
	rs := ParseRuleset(`
"Goodbye" -> "Hello "
" Bye  "  ->"Hi"
	`)
	assert.Equal(t, []Rule{
		{From: "Goodbye", To: "Hello "},
		{From: " Bye  ", To: "Hi"},
	}, rs)
}
