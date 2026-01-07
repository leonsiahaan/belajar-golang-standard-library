package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Leon Siahaan", "Leon"))
	fmt.Println(strings.Split("Leon Siahaan", " "))
	fmt.Println(strings.ToLower("Leon Siahaan"))
	fmt.Println(strings.ToUpper("Leon Siahaan"))
	fmt.Println(strings.Trim("    Leon Siahaan    ", " "))
	fmt.Println(strings.ReplaceAll("Leon Siahaan Hinalang", "Leon", "Juan"))

}
