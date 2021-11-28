package main

import (
	"fmt"
	"regexp" // github.com/google/re2/wiki/Syntax or golang.org/pkg/regexp
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")
	fmt.Println(regex.MatchString("elo")) // true (l is letter between a-z)
	fmt.Println(regex.MatchString("eto")) // true (l is letter between a-z)
	fmt.Println(regex.MatchString("eDo")) // false (D is not letter between a-z)

	search := regex.FindAllString("elo eto ewo ego eyi", -1) // show all true data of var regexp
	fmt.Println(search)
}
