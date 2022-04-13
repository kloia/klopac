package main

import (
	"entrypoint/pkg/helper"
	"entrypoint/pkg/klopac"
)

func main() {
	helper.InitializeLogger()
	klopac.Run()
}
