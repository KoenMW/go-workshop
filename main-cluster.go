package main

import (
	"fmt"
	"go-workshop/clustering"
	"go-workshop/csvutil"
)

func main() {
	// Load weapons from CSV file (assuming the file always exists and is valid)
	weapons, _ := csvutil.LoadWeapons("data/elden_ring_weapon.csv")

	// Select features to use for clustering
	fields := []string{"Phy", "Mag", "Fir", "Lit", "Hol", "Sta", "Str", "Dex"}
	points := clustering.GenerateDataPoints(weapons, fields)

	// Perform k-means clustering with k = 3
	assignments := clustering.RunKMeans(points, 3)

	// Reduce to 2D using PCA
	coords, _ := clustering.ProjectTo2D(points)

	// Plot clusters and save to output.png
	clustering.PlotClusters(coords, assignments, "output.png")

	// Print weapon names and their assigned clusters
	for i, cluster := range assignments {
		fmt.Printf("Weapon: %-35s | Type: %-20s | Cluster: %d\n", points[i].Name, points[i].Type, cluster)
	}
}
