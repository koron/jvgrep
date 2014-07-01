package main

import (
	gmigemo "github.com/koron/gomigemo/migemo"
	"regexp"
)

func migemoCompile(s string) (*regexp.Regexp, error) {
	dict, err := gmigemo.LoadDefault()
	if err != nil {
		return nil, err
	}
	rx, err := gmigemo.Compile(dict, s)
	if err != nil {
		return nil, err
	}
	return rx, nil
}
