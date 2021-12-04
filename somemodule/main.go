package main

import (
	"fmt"

	"github.com/nechvatalp/GoVersioningTest/somemodule/packageone"
	"github.com/nechvatalp/GoVersioningTest/somemodule/packagetwo"
)

func main() {
	fmt.Println("Starting")
	packageone.HelloPackageOne()
	packagetwo.HelloPackageTwo()
}
