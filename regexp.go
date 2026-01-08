package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`l([a-z])n`)

	fmt.Println(regex.MatchString("len"))
	fmt.Println(regex.MatchString("lin"))
	fmt.Println(regex.MatchString("lud"))

	fmt.Println(regex.FindAllString("len lin lun lud lad lOd lid", 10))
}
