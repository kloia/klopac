package flag

import (
	"flag"
)

type Flag interface {
	Parse()
	Bool(string, bool, string) *bool
	String(string, string, string) *string
}

type flagService struct {
}

func (f flagService) Parse() {
	flag.Parse()
}

func (f flagService) Bool(name string, value bool, usage string) *bool {
	return flag.Bool(name, value, usage)
}

func (f flagService) String(name string, value string, usage string) *string {
	return flag.String(name, value, usage)
}

func NewFlagService() Flag {
	return &flagService{}
}
