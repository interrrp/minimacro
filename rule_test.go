package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyRule(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		r := Rule{From: "Goodbye", To: "Hello"}
		assert.Equal(t, "Hello, world!", r.Apply("Goodbye, world!"))
	})
	t.Run("with whitespace", func(t *testing.T) {
		r := Rule{From: "Goodbye, ", To: "Hello, "}
		assert.Equal(t, "Hello, world!", r.Apply("Goodbye, world!"))
	})
}

func TestApplyRuleset(t *testing.T) {
	rs := Ruleset{
		{From: "Goodbye", To: "Hello"},
		{From: "adieu", To: "bonjour"},
	}
	assert.Equal(t, "Hello (bonjour), world (monde)!", rs.Apply("Goodbye (adieu), world (monde)!"))
}
