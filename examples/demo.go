package main

import (
	"fmt"

	clear "github.com/tjhorner/goclear"
)

func main() {
	// Get the current regions on the CodeDay website
	regions := clear.GetCurrentRegions()

	// List them all
	for _, region := range regions {
		fmt.Println(region.Name)
	}
}
