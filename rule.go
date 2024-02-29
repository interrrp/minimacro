package main

import "strings"

type Rule struct {
	From string
	To   string
}

type Ruleset []Rule

func (r *Rule) Apply(str string) string {
	return strings.ReplaceAll(str, r.From, r.To)
}

func (rs *Ruleset) Apply(str string) string {
	out := str
	for _, r := range *rs {
		out = r.Apply(out)
	}
	return out
}
