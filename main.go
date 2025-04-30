package main

import (
	"Assessment/app"
	_ "Assessment/docs" // Swagger generated files
	"flag"
	"Assessment/tests"

)



func main() {

	var runUnitTests bool

	flag.BoolVar(&runUnitTests, "runUnitTests", true, "Setting to true will run unit tests")
	flag.Parse()

	if runUnitTests {
		tests.Start()
	} else {
		app.Start()
	}
}
