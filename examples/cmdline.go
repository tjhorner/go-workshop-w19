package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	clear "github.com/tjhorner/goclear"
)

func main() {
	// Get the region that the user wants
	region := flag.String("region", "", "The CodeDay region to get info for.")
	flag.Parse()

	// Check to see if they actually provided it.
	if *region == "" {
		fmt.Printf("Please specify a region to get info for.\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Get the info
	regionData := clear.GetRegion(*region)

	fmt.Printf("CodeDay Region: %s\n", regionData.Name)

	if regionData.CurrentEvent.ID != "" {
		startsAt := time.Unix(int64(regionData.CurrentEvent.StartsAt), 0)
		endsAt := time.Unix(int64(regionData.CurrentEvent.EndsAt), 0)

		fmt.Printf(
			"Current Event: %s\nStarts At: %v\nEnds At: %s\n",
			regionData.CurrentEvent.Name,
			startsAt,
			endsAt,
		)
	} else {
		fmt.Printf("  (no event this season)\n")
	}
}
