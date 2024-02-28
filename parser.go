package main

import (
	"regexp"
	"strings"
)

type Rule struct {
	From string
	To   string
}

var ruleRegex = regexp.MustCompile("\"(?P<from>.+)\" *-> *\"(?P<to>.+)\"")

func ParseRule(ruleStr string) Rule {
	groups := getNamedGroups(ruleStr, ruleRegex)
	return Rule{
		From: groups["from"],
		To:   groups["to"],
	}
}

func ParseRuleset(set string) []Rule {
	rules := []Rule{}
	for _, line := range strings.Split(set, "\n") {
		if isEmpty(line) {
			continue
		}
		rules = append(rules, ParseRule(line))
	}
	return rules
}

func getNamedGroups(str string, re *regexp.Regexp) map[string]string {
	groups := map[string]string{}
	match := re.FindStringSubmatch(str)
	if match != nil {
		groupNames := re.SubexpNames()
		for i, name := range groupNames {
			if i != 0 && name != "" {
				groups[name] = match[i]
			}
		}
	}
	return groups
}

func isEmpty(line string) bool {
	return strings.TrimSpace(line) == ""
}
