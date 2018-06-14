package foo

import (
	"fmt"
	"log"

	"github.com/golang/dep"
)

func logFoo() {
	log.Println("Foo")

	d := dep.Analyzer{}
	fmt.Println(d)
}
