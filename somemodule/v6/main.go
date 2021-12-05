package main

import (
	"fmt"

	"github.com/nechvatalp/GoVersioningTest/somemodule/v6/packageone"
	"github.com/nechvatalp/GoVersioningTest/somemodule/v6/packagetwo"
)

func main() {
	fmt.Println("Starting")
	packageone.HelloPackageOne()
	packagetwo.HelloPackageTwo()
}
