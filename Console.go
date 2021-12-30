package Util

import "fmt"

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

// Color Format text with custom color
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

// PrintIfNotNil Print a text if not nil
func PrintIfNotNil(a interface{}, description ...interface{}) {
	if a != nil {
		if description != nil {
			fmt.Print(description, ": ")
		}
		fmt.Println(a)
	}
}

// PrintSuccessIfNotNil Print a text green if not nil
func PrintSuccessIfNotNil(a interface{}, description ...interface{}) {
	if a != nil {
		if description != nil {
			fmt.Print(Green(description), Green(": "))
		}
		fmt.Println(Green(a))
	}
}

// PrintErrorIfNotNil Print a text red if not nil
func PrintErrorIfNotNil(a interface{}, description ...interface{}) {
	if a != nil {
		if description != nil {
			fmt.Print(Red(description), Red(": "))
		}
		fmt.Println(Red(a))
	}
}

// PrintWarnIfNotNil Print a text yellow if not nil
func PrintWarnIfNotNil(a interface{}, description ...interface{}) {
	if a != nil {
		if description != nil {
			fmt.Print(Yellow(description), Yellow(": "))
		}
		fmt.Println(Yellow(a))
	}
}
