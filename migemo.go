package main

import (
	gmigemo "github.com/koron/gomigemo/migemo"
	"os"
	"path"
	"regexp"
)

func dictdir() string {
	d := os.Getenv("GMIGEMO_DICTDIR")
	if d != "" {
		return d
	}
	d = os.Getenv("GOPATH")
	if d == "" {
		return "./_dict"
	}
	return path.Join(d, "src", "github.com", "koron", "gomigemo", "_dict")
}

func migemoCompile(s string) (*regexp.Regexp, error) {
	dict, err := gmigemo.Load(dictdir())
	if err != nil {
		return nil, err
	}
	rx, err := gmigemo.Compile(dict, s)
	if err != nil {
		return nil, err
	}
	return rx, nil
}
