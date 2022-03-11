package helpers

import (
	"github.com/fatih/color"
)

type LoggerFuncs interface {
	Error()
	Warn()
	Verbose()
}

type Logger struct {
	Name string
}

//Used for a formatted and colored outputt{o console.
//this improves visibility
func Log(mode string, name string, message string) {
	c := color.New(color.FgWhite)
	if mode == "Error" {
		c = color.New(color.FgRed)
	} else if mode == "Warning" {
		c = color.New(color.FgYellow)
	} else if mode == "Verbose" {
		c = color.New(color.FgGreen)
	}
	c.Println("[ " + mode + " ] " + name + ": \t " + message)
}

func (log Logger) Error(message string) {
	Log("Error", log.Name, message)
}

func (log Logger) Warn(message string) {
	Log("Warning", log.Name, message)
}
func (log Logger) Verbose(message string) {
	Log("Verbose", log.Name, message)
}
