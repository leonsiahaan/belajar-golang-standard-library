package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"leon", "siahaan", "hinalang"})
	_ = writer.Write([]string{"albert", "paulina", "harianja"})
	_ = writer.Write([]string{"caped", "baldy", "king"})

	writer.Flush()
}
