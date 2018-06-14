package main

import (
	"fmt"
	"log"
	"github.com/golang/dep"
)


func printFoo() {
	fmt.Println("Foo")
}


func logFoo() {
	log.Println("Foo")

	d := dep.Analyzer{}
	fmt.Println(d)
}
