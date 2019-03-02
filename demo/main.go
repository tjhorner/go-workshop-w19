package main

import (
	"fmt"

	clear "github.com/tjhorner/goclear"
)

func main() {
	regions := clear.GetCurrentRegions()

	for _, region := range regions {
		fmt.Printf("CodeDay Region: %s\n", region.Name)
	}
}
