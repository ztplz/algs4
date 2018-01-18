package main

import (
	"fmt"

	"github.com/ztplz/algs4/bag"
)

func main() {
	sbag := bag.NewSliceBag()

	sbag.Add("asgasgd")
	sbag.Add("asgasgd")
	sbag.Add("asgasgd")
	sbag.Add("asgasgd")

	sbag.Add("asgasgd")
	fmt.Println(sbag.Size())
}
