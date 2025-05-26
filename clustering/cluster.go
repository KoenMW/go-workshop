package clustering

import (
	"fmt"
	"go-workshop/csvutil"         // Our own package to load weapon data
	"image/color"
	"log"
	"math"
	"math/rand"
	"sort"
	"time"

	"gonum.org/v1/gonum/mat"      // For matrices and PCA
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"           // For plotting
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// DataPoint holds weapon stats and metadata (used for clustering)
type DataPoint struct {
	Features []float64
	Name     string
	Type     string
}

// GenerateDataPoints selects specific fields from each weapon for clustering
func GenerateDataPoints(weapons []csvutil.Weapon, fields []string) []DataPoint {
	var points []DataPoint

	for _, w := range weapons {
		values := []float64{}
		v := weaponToMap(w)
		for _, f := range fields {
			values = append(values, v[f]) // collect only selected fields
		}
		points = append(points, DataPoint{
			Features: values,
			Name:     w.Name,
			Type:     w.Type,
		})
	}
	return points
}

// Helper function to convert a Weapon struct to a map of stat names to values
func weaponToMap(w csvutil.Weapon) map[string]float64 {
	return map[string]float64{
		"Phy": w.Phy, "Mag": w.Mag, "Fir": w.Fir, "Lit": w.Lit,
		"Hol": w.Hol, "Sta": w.Sta, "Str": w.Str, "Dex": w.Dex,
		"Int": w.Int, "Fai": w.Fai, "Arc": w.Arc, "Bst": w.Bst,
		"Rst": w.Rst, "Wgt": w.Wgt,
	}
}

// RunKMeans applies basic k-means clustering to the datapoints
func RunKMeans(points []DataPoint, k int) []int {
	rand.Seed(time.Now().UnixNano())

	n := len(points)
	dim := len(points[0].Features)

	assignments := make([]int, n)        // cluster assignment per point
	centroids := make([][]float64, k)    // cluster centers

	// Randomly initialize centroids
	for i := 0; i < k; i++ {
		centroids[i] = make([]float64, dim)
		copy(centroids[i], points[rand.Intn(n)].Features)
	}

	// Repeat k-means for 20 iterations
	for iter := 0; iter < 20; iter++ {
		// Assign each point to the closest centroid
		for i, p := range points {
			closest := 0
			minDist := distance(p.Features, centroids[0])
			for j := 1; j < k; j++ {
				if d := distance(p.Features, centroids[j]); d < minDist {
					closest = j
					minDist = d
				}
			}
			assignments[i] = closest
		}

		// Recalculate centroids
		counts := make([]int, k)
		newCentroids := make([][]float64, k)
		for i := range newCentroids {
			newCentroids[i] = make([]float64, dim)
		}
		for i, a := range assignments {
			for d := 0; d < dim; d++ {
				newCentroids[a][d] += points[i].Features[d]
			}
			counts[a]++
		}
		for i := 0; i < k; i++ {
			if counts[i] == 0 {
				continue
			}
			for d := 0; d < dim; d++ {
				newCentroids[i][d] /= float64(counts[i])
			}
		}
		centroids = newCentroids
	}

	return assignments
}

// Euclidean distance between two vectors
func distance(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		diff := a[i] - b[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

// ProjectTo2D reduces dimensionality using PCA so we can visualize on a 2D plot
func ProjectTo2D(points []DataPoint) ([][2]float64, error) {
	numPoints := len(points)
	numFeatures := len(points[0].Features)

	data := mat.NewDense(numPoints, numFeatures, nil)
	for i, p := range points {
		data.SetRow(i, p.Features)
	}

	var pc stat.PC
	if ok := pc.PrincipalComponents(data, nil); !ok {
		return nil, fmt.Errorf("PCA failed")
	}

	vecs := mat.NewDense(numFeatures, numFeatures, nil)
	pc.VectorsTo(vecs)

	coords := make([][2]float64, numPoints)
	for i := 0; i < numPoints; i++ {
		x, y := 0.0, 0.0
		for j := 0; j < numFeatures; j++ {
			val := data.At(i, j)
			x += val * vecs.At(j, 0) // PC1
			y += val * vecs.At(j, 1) // PC2
		}
		coords[i] = [2]float64{x, y}
	}

	return coords, nil
}

// PlotClusters creates a scatterplot and draws a convex hull for each cluster
func PlotClusters(coords [][2]float64, clusters []int, filename string) {
	p := plot.New()
	p.Title.Text = "Weapon Clusters"
	p.X.Label.Text = "PC1"
	p.Y.Label.Text = "PC2"

	// Determine number of clusters
	k := 0
	for _, c := range clusters {
		if c+1 > k {
			k = c + 1
		}
	}

	// Predefined colors for each cluster
	colors := []color.RGBA{
		{255, 0, 0, 100}, {0, 200, 0, 100}, {0, 0, 255, 100},
		{255, 165, 0, 100}, {128, 0, 128, 100}, {0, 255, 255, 100},
		{255, 192, 203, 100}, {128, 128, 128, 100}, {0, 128, 128, 100},
	}

	// Group points by cluster
	clusterPlots := make([]plotter.XYs, k)
	for i := range clusterPlots {
		clusterPlots[i] = plotter.XYs{}
	}
	for i, coord := range coords {
		cluster := clusters[i]
		clusterPlots[cluster] = append(clusterPlots[cluster], plotter.XY{X: coord[0], Y: coord[1]})
	}

	// Draw scatter points and convex hull for each cluster
	for i, pts := range clusterPlots {
		if len(pts) == 0 {
			continue
		}

		hull := convexHull(pts)
		hull = append(hull, hull[0]) // Close the polygon

		poly, err := plotter.NewLine(hull)
		if err != nil {
			log.Fatal(err)
		}
		poly.Color = colors[i%len(colors)]
		poly.Width = vg.Points(1)

		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Color = colors[i%len(colors)]
		s.GlyphStyle.Radius = vg.Points(2)

		p.Add(poly, s)
		p.Legend.Add(fmt.Sprintf("Cluster %d", i), s)
	}

	if err := p.Save(6*vg.Inch, 6*vg.Inch, filename); err != nil {
		log.Fatal(err)
	}
}

// convexHull returns the outer boundary of a set of 2D points
func convexHull(points plotter.XYs) plotter.XYs {
	if len(points) < 3 {
		return points
	}

	// Sort points by X (then Y)
	sort.Slice(points, func(i, j int) bool {
		if points[i].X == points[j].X {
			return points[i].Y < points[j].Y
		}
		return points[i].X < points[j].X
	})

	// Cross product for checking convexity
	cross := func(o, a, b plotter.XY) float64 {
		return (a.X-o.X)*(b.Y-o.Y) - (a.Y-o.Y)*(b.X-o.X)
	}

	var lower plotter.XYs
	for _, p := range points {
		for len(lower) >= 2 && cross(lower[len(lower)-2], lower[len(lower)-1], p) <= 0 {
			lower = lower[:len(lower)-1]
		}
		lower = append(lower, p)
	}

	var upper plotter.XYs
	for i := len(points) - 1; i >= 0; i-- {
		p := points[i]
		for len(upper) >= 2 && cross(upper[len(upper)-2], upper[len(upper)-1], p) <= 0 {
			upper = upper[:len(upper)-1]
		}
		upper = append(upper, p)
	}

	return append(lower[:len(lower)-1], upper[:len(upper)-1]...)
}
