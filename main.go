package main

import "github.com/MangioneAndrea/GoUtils/console"

func main() {
	console.Log(nil)
	console.Log(nil, "do not show", console.ShowIfNotNil)
	console.Error(nil, "show")
	console.Error(nil, "do not show", console.ShowIfNotNil)
	console.Success(nil, "show")
	console.Success(nil, "do not show", console.ShowIfNotNil)
	console.Warn(nil, "show")
	console.Warn("hello", "do not show", console.ShowIfNotNil)

}
