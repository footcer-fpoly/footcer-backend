package main

import (
	"fmt"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"log"
)

func main()  {
	c, err := maps.NewClient(maps.WithAPIKey("key"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DistanceMatrixRequest{
		Origins:      []string{"10.8094459,106.6753573"},
		Destinations: []string{"10.8538546,106.6261668"},
		Units:        maps.UnitsImperial,
		Language:     "en",
		// Must specify DepartureTime in order to get DurationInTraffic in response
		DepartureTime: "now",
	}
	route, err := c.DistanceMatrix(context.Background(), r)

	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	fmt.Printf("Duration in minutes: %f\n", route.Rows[0].Elements[0].Distance.Meters)
}
