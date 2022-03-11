package main

import (
	Logger "spotify2gether.server/pkg/helpers"
)

func main() {
	logger := Logger.Logger{Name: "main"}
	logger.Verbose("Starting")
}
