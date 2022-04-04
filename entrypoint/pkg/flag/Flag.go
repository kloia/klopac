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

// Flag service and it helps us to parse command-line flags from os.Args[1:]
func (f flagService) Parse() {
	flag.Parse()
}

// This function help us to make use of a boolean logic to decide whether we are going to use some sort of functions or not. It defines a bool flag with specified name
func (f flagService) Bool(name string, value bool, usage string) *bool {
	return flag.Bool(name, value, usage)
}

// Defines a string flag with specified name, default value, and usage string. The return value is the address of a string variable that stores the value of the flag
func (f flagService) String(name string, value string, usage string) *string {
	return flag.String(name, value, usage)
}

func NewFlagService() Flag {
	return &flagService{}
}
