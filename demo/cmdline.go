package main

import (
	"flag"
	"fmt"
	"os"

	clear "github.com/tjhorner/goclear"
)

func main() {
	region := flag.String("region", "", "The CodeDay region to get information for.")
	flag.Parse()

	if *region == "" {
		fmt.Println("Please specify a CodeDay region with the region flag!")
		os.Exit(1)
	}

	regionData := clear.GetRegion(*region)

	fmt.Printf("CodeDay Region: %s\n", regionData.Name)
	fmt.Printf("Venue this season: %s\n", regionData.CurrentEvent.Venue.Name)
}
