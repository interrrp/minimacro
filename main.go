package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	rulesetPath string
	glob        string
)

func main() {
	flag.StringVar(&rulesetPath, "ruleset", ".minimacro", "path to ruleset file")
	flag.StringVar(&glob, "glob", "", "files to affect")
	flag.Parse()

	rsBytes, err := os.ReadFile(rulesetPath)
	if err != nil {
		fatal(err)
	}
	ruleset := ParseRuleset(string(rsBytes))

	matches, err := filepath.Glob(glob)
	if err != nil {
		fatal(err)
	}
	for _, match := range matches {
		cont, err := os.ReadFile(match)
		if err != nil {
			fmt.Fprintf(os.Stderr, "skipping %s: %s", match, err)
			continue
		}

		out := string(cont)
		for _, r := range ruleset {
			out = r.Apply(out)
		}
		os.WriteFile(match, []byte(out), 0777)
		fmt.Println(match)
	}
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "fatal error: %s", err)
	os.Exit(1)
}
