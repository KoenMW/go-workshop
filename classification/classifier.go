package classification

import (
	"math"
	"sort"
)

// Instance and Classifier definitions

type Instance struct {
	Features []float64
	Label    string
}

type Classifier struct {
	Instances []Instance
	K         int
}

func NewClassifier(k int) *Classifier {
	return &Classifier{K: k}
}

func (c *Classifier) Train(instances []Instance) {
	c.Instances = instances
}

func (c *Classifier) Predict(features []float64) string {
	type neighbor struct {
		distance float64
		label    string
	}

	var neighbors []neighbor
	for _, instance := range c.Instances {
		dist := euclideanDistance(features, instance.Features)
		neighbors = append(neighbors, neighbor{distance: dist, label: instance.Label})
	}

	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].distance < neighbors[j].distance
	})

	votes := make(map[string]int)
	for i := 0; i < c.K && i < len(neighbors); i++ {
		votes[neighbors[i].label]++
	}

	return maxVote(votes)
}

func euclideanDistance(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		diff := a[i] - b[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

func maxVote(votes map[string]int) string {
	max := 0
	label := ""
	for k, v := range votes {
		if v > max {
			max = v
			label = k
		}
	}
	return label
}
