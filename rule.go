package main

import "strings"

type Rule struct {
	From string
	To   string
}

func (r *Rule) Apply(str string) string {
	return strings.ReplaceAll(str, r.From, r.To)
}
