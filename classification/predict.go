package classification

import (
	"fmt"
	"go-workshop-1/csvutil"
	"log"
)

// classifier instance stored at package level for reuse
var clf *Classifier

// Init loads weapons, converts to Instances, trains classifier — call once at program start
func Init(k int, csvPath string) {
	weapons, err := csvutil.LoadWeapons(csvPath)
	if err != nil {
		log.Fatalf("Failed to load weapons: %v", err)
	}

	var instances []Instance
	for _, w := range weapons {
		features := []float64{
			w.Phy, w.Mag, w.Fir, w.Lit, w.Hol, w.Cri, w.Sta,
			w.Str, w.Dex, w.Int, w.Fai, w.Arc, w.Any, w.Bst, w.Rst, w.Wgt,
		}
		instances = append(instances, Instance{
			Features: features,
			Label:    w.Type,
		})
	}

	clf = NewClassifier(k)
	clf.Train(instances)
}

// Predict accepts exactly 16 float64 stats and returns the predicted weapon type
func Predict(stats ...float64) string {
	if clf == nil {
		panic("Classifier not initialized. Call classification.Init first.")
	}
	if len(stats) != 16 {
		panic(fmt.Sprintf("Expected 16 stats, got %d", len(stats)))
	}
	return clf.Predict(stats)
}
