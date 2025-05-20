package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Weapon struct represents all attributes from the CSV
type Weapon struct {
	Name    string
	Type    string
	Phy     float64
	Mag     float64
	Fir     float64
	Lit     float64
	Hol     float64
	Cri     float64
	Sta     float64
	Str     float64
	Dex     float64
	Int     float64
	Fai     float64
	Arc     float64
	Any     float64
	Bst     float64
	Rst     float64
	Wgt     float64
	Upgrade string
}

// neighbor type for KNN implementation
type neighbor struct {
	distance float64
	label    string
}

// loadWeapons loads data from CSV file
func loadWeapons(filename string) ([]Weapon, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var weapons []Weapon

	// Assuming first row is header
	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}

		weapon := Weapon{
			Name:    record[0],
			Type:    record[1],
			Upgrade: record[18], // Assuming Upgrade is the last column
		}

		// Parse all numerical values
		for j := 2; j < 18; j++ {
			val, err := strconv.ParseFloat(record[j], 64)
			if err != nil {
				val = 0 // Default value if parsing fails
			}

			switch j {
			case 2:
				weapon.Phy = val
			case 3:
				weapon.Mag = val
			case 4:
				weapon.Fir = val
			case 5:
				weapon.Lit = val
			case 6:
				weapon.Hol = val
			case 7:
				weapon.Cri = val
			case 8:
				weapon.Sta = val
			case 9:
				weapon.Str = val
			case 10:
				weapon.Dex = val
			case 11:
				weapon.Int = val
			case 12:
				weapon.Fai = val
			case 13:
				weapon.Arc = val
			case 14:
				weapon.Any = val
			case 15:
				weapon.Bst = val
			case 16:
				weapon.Rst = val
			case 17:
				weapon.Wgt = val
			}
		}

		weapons = append(weapons, weapon)
	}

	return weapons, nil
}

// KNN implementation
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
	var neighbors []neighbor

	for _, instance := range c.Instances {
		dist := euclideanDistance(features, instance.Features)
		neighbors = append(neighbors, neighbor{distance: dist, label: instance.Label})
	}

	sortNeighbors(neighbors)

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

func sortNeighbors(neighbors []neighbor) {
	for i := 0; i < len(neighbors); i++ {
		for j := i + 1; j < len(neighbors); j++ {
			if neighbors[i].distance > neighbors[j].distance {
				neighbors[i], neighbors[j] = neighbors[j], neighbors[i]
			}
		}
	}
}

func maxVote(votes map[string]int) string {
	max := 0
	var label string
	for k, v := range votes {
		if v > max {
			max = v
			label = k
		}
	}
	return label
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Load data from CSV
	weapons, err := loadWeapons("elden_ring_weapon.csv")
	if err != nil {
		fmt.Printf("Error loading data: %v\n", err)
		return
	}

	// Prepare instances for KNN
	var instances []Instance
	for _, weapon := range weapons {
		features := []float64{
			weapon.Phy, weapon.Mag, weapon.Fir, weapon.Lit, weapon.Hol,
			weapon.Cri, weapon.Sta, weapon.Str, weapon.Dex, weapon.Int,
			weapon.Fai, weapon.Arc, weapon.Any, weapon.Bst, weapon.Rst,
			weapon.Wgt,
		}
		instances = append(instances, Instance{
			Features: features,
			Label:    weapon.Type,
		})
	}

	// Shuffle data
	rand.Shuffle(len(instances), func(i, j int) {
		instances[i], instances[j] = instances[j], instances[i]
	})

	// Split into training and test sets (80/20)
	split := int(0.8 * float64(len(instances)))
	trainSet := instances[:split]
	testSet := instances[split:]

	// Create and train classifier
	classifier := NewClassifier(5)
	classifier.Train(trainSet)

	// Evaluate
	correct := 0
	for _, test := range testSet {
		predicted := classifier.Predict(test.Features)
		if predicted == test.Label {
			correct++
		}
	}

	accuracy := float64(correct) / float64(len(testSet)) * 100
	fmt.Printf("Model accuracy: %.2f%%\n", accuracy)

	// Example prediction
	sampleFeatures := []float64{
		128, 0, 0, 0, 0, // Phy, Mag, Fir, Lit, Hol
		100, 45, 16, 10, 0, // Cri, Sta, Str, Dex, Int
		0, 0, 0, 46, 30, 5.5, // Fai, Arc, Any, Bst, Rst, Wgt
	}

	predictedType := classifier.Predict(sampleFeatures)
	fmt.Printf("Predicted weapon type for sample: %s\n", predictedType)
}
