package main

import (
	"fmt"
	"go-workshop/clustering"
	"go-workshop/csvutil"
)

func main() {
	weapons, err := csvutil.LoadWeapons("data/elden_ring_weapon.csv")
	if err != nil {
		panic(err)
	}

	fields := []string{"Phy", "Mag", "Fir", "Lit", "Hol", "Sta", "Str", "Dex"}
	points := clustering.GenerateDataPoints(weapons, fields)

	assignments := clustering.RunKMeans(points, 3)

	coords, err := clustering.ProjectTo2D(points)
	if err != nil {
		panic(err)
	}

	clustering.PlotClusters(coords, assignments, "output.png")

	for i, cluster := range assignments {
		fmt.Printf("Weapon: %-35s | Type: %-20s | Cluster: %d\n", points[i].Name, points[i].Type, cluster)
	}
}
