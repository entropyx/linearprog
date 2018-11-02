package linearprog

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

type Metrics struct {
	Date                 []string
	AvgCPC               []float64
	Clicks               []float64
	Roas                 []float64
	Prices               []float64
	Impressions          []float64
	Conversions          []float64
	TotalValueConversion []float64
	Cost                 []float64
	AvgCPM               []float64
	ClickShare           []float64
	ImpressionShare      []float64
	ConversionRate       []float64
	Ctr                  []float64
	CostPerConversion    []float64
	CostConvertedClick   []float64
	ProductCost          []float64
	PositionPrice        int
	RoasBrand            float64
	RoasCategory         float64
	AvgPriceBrand        float64
	AvgPriceCategory     float64
	CpcBrand             float64
	CpcCategory          float64
	Price                float64
	MinimumCpc           float64
	Brand                string
	Category             string
	ProductType          string
	RoasWithoutProd      float64
	AllConversionValue   float64
	TotalCost            float64
	TotalClicks          float64
	TotalConversions     float64
	TotalImpressions     float64
	TotalAvgCPM          float64
	TotalAvgCPC          float64
	Distances            float64
	Rank                 int
	Pareto               bool
}

func GetPages(path string) map[string]*Metrics {
	path = os.Getenv("GOPATH") + path
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	var c map[string]*Metrics
	json.Unmarshal(raw, &c)
	return c
}

func TestSimplex(t *testing.T) {
	// Convey("Given the following matrix maximization problem 1 ...", t, func() {
	// 	a := []float64{40, 60}
	// 	b := []float64{70, 40, 90}
	// 	constdir := []string{"<=", "<=", "<="}
	// 	A := map[int]map[int]float64{
	// 		1: map[int]float64{1: 2, 2: 1},
	// 		2: map[int]float64{1: 1, 2: 1},
	// 		3: map[int]float64{1: 1, 2: 3},
	// 	}
	// 	Convey("The solutions to 1 should be ... ", func() {
	// 		solutions, _ := Simplex(A, b, a, constdir)
	// 		out := map[int]float64{0: 15, 1: 25}
	// 		So(solutions, ShouldResemble, out)
	// 	})
	// })
	//
	Convey("Given the following matrix maximization problem 2 ... ", t, func() {
		a := []float64{3, 2, 5}
		b := []float64{430, 460, 420}
		constdir := []string{"<=", "<=", "<="}
		A := map[int]map[int]float64{
			1: map[int]float64{1: 1, 2: 2, 3: 1},
			2: map[int]float64{1: 3, 2: 0, 3: 2},
			3: map[int]float64{1: 1, 2: 4, 3: 0},
		}
		Convey("The solutions to 2 should be ... ", func() {
			solutions, _ := Simplex(A, b, a, constdir)
			out := map[int]float64{0: 0, 1: 100, 2: 230}
			So(solutions, ShouldResemble, out)
		})
	})

	// Convey("Given the following matrix maximization problem 3 ...", t, func() {
	// 	a := []float64{3, 2}
	// 	b := []float64{18, 42, 24}
	// 	constdir := []string{"<=", "<=", "<="}
	// 	A := map[int]map[int]float64{
	// 		1: map[int]float64{1: 2, 2: 1},
	// 		2: map[int]float64{1: 2, 2: 3},
	// 		3: map[int]float64{1: 3, 2: 1},
	// 	}
	// 	Convey("The solutions to 3 should be ... ", func() {
	// 		solutions, _ := Simplex(A, b, a, constdir)
	// 		out := map[int]float64{0: 3, 1: 12}
	// 		So(solutions, ShouldResemble, out)
	// 	})
	// })
	//
	// Convey("Given the following matrix maximization problem 4 ...", t, func() {
	// 	a := []float64{2, 3, 10, 5, 3}
	// 	b := []float64{425, 25, 120, 6, 1}
	// 	constdir := []string{"<=", "<=", "<=", "<=", "<="}
	// 	A := map[int]map[int]float64{
	// 		1: map[int]float64{1: 1, 2: 1, 3: 1, 4: 3, 5: 2},
	// 		2: map[int]float64{1: 2, 2: -4, 3: 1, 4: 0, 5: 0},
	// 		3: map[int]float64{1: -1, 2: -1, 3: -3, 4: 0, 5: 0},
	// 		4: map[int]float64{1: 1, 2: 11, 3: 5, 4: 0, 5: 0},
	// 		5: map[int]float64{1: 1, 2: 1, 3: 7, 4: 0, 5: 0},
	// 	}
	// 	Convey("The solutions to 4 should be ... ", func() {
	// 		solutions, _ := Simplex(A, b, a, constdir)
	// 		out := map[int]float64{0: 0, 1: 0.51, 2: 0.07, 3: 141.47, 4: 0}
	// 		So(solutions, ShouldResemble, out)
	// 	})
	// })
	//
	// Convey("Given the following matrix maximization problem 5 ...", t, func() {
	// 	a := []float64{225000, 18000, 92000}
	// 	b := []float64{17000, 1.25, 1.25, 1.25, 1.0, 0.8, 0.1}
	// 	constdir := []string{"<=", "<=", "<=", "<=", ">=", ">=", ">="}
	// 	A := map[int]map[int]float64{
	// 		1: map[int]float64{1: 11000, 2: 988, 3: 5000},
	// 		2: map[int]float64{1: 1, 2: 0, 3: 0},
	// 		3: map[int]float64{1: 0, 2: 1, 3: 0},
	// 		4: map[int]float64{1: 0, 2: 0, 3: 1},
	// 		5: map[int]float64{1: 1, 2: 0, 3: 0},
	// 		6: map[int]float64{1: 0, 2: 1, 3: 0},
	// 		7: map[int]float64{1: 0, 2: 0, 3: 1},
	// 	}
	// 	Convey("The solutions to 5 should be ... ", func() {
	// 		solutions, _ := Simplex(A, b, a, constdir)
	// 		out := map[int]float64{0: 1.25, 1: 0.8, 2: 0.49}
	// 		So(solutions, ShouldResemble, out)
	// 	})
	// })

	// Convey("Given the following matrix maximization problem 6 ...", t, func() {
	// 	a := []float64{1, 1, 2}
	// 	b := []float64{50, 36, 10}
	// 	constdir := []string{"<=", ">=", ">="}
	// 	A := map[int]map[int]float64{
	// 		1: map[int]float64{1: 2, 2: 1, 3: 1},
	// 		2: map[int]float64{1: 2, 2: 1, 3: 0},
	// 		3: map[int]float64{1: 1, 2: 0, 3: 1},
	// 	}
	// 	Convey("The solutions to 3 should be ... ", func() {
	// 		solutions, _ := Simplex(A, b, a, constdir)
	// 		out := map[int]float64{0: 0, 1: 36, 2: 14}
	// 		So(solutions, ShouldResemble, out)
	// 	})
	// })
	//
	// Convey("Given the following matrix maximization problem 7 ...", t, func() {
	// 	a := []float64{1450, 1450, 1450, 3950}
	// 	b := []float64{40064, 1.25, 1.25, 1.25, 1.25, 0.96, 0.68, 0.39, 0.1}
	// 	constdir := []string{"<=", "<=", "<=", "<=", "<=", ">=", ">=", ">=", ">="}
	// 	A := map[int]map[int]float64{
	// 		1: map[int]float64{1: 1430.86, 2: 1430.86, 3: 1430.86, 4: 1430.86},
	// 		2: map[int]float64{1: 1, 2: 0, 3: 0, 4: 0},
	// 		3: map[int]float64{1: 0, 2: 1, 3: 0, 4: 0},
	// 		4: map[int]float64{1: 0, 2: 0, 3: 1, 4: 0},
	// 		5: map[int]float64{1: 0, 2: 0, 3: 0, 4: 1},
	// 		6: map[int]float64{1: 1, 2: 0, 3: 0, 4: 0},
	// 		7: map[int]float64{1: 0, 2: 1, 3: 0, 4: 0},
	// 		8: map[int]float64{1: 0, 2: 0, 3: 1, 4: 0},
	// 		9: map[int]float64{1: 0, 2: 0, 3: 0, 4: 1},
	// 	}
	// 	Convey("The solutions to 7 should be ... ", func() {
	// 		solutions, _ := Simplex(A, b, a, constdir)
	// 		out := map[int]float64{0: 1.25, 1: 1.25, 2: 1.25, 3: 1.25}
	// 		So(solutions, ShouldResemble, out)
	// 	})
	// })

	Convey("Given the following matrix maximization problem 8 ...", t, func() {
		a := []float64{1, 24, 10}
		b := []float64{84.75, 2.5, 3.5, 0.5}
		constdir := []string{"=", "<=", "<=", "<="}
		A := map[int]map[int]float64{
			1: map[int]float64{1: 5, 2: 25, 3: 0.5},
			2: map[int]float64{1: 1, 2: 0, 3: 0},
			3: map[int]float64{1: 0, 2: 1, 3: 0},
			4: map[int]float64{1: 0, 2: 0, 3: 1},
		}
		Convey("The solutions to 3 should be ... ", func() {
			solutions, _ := Simplex(A, b, a, constdir)
			out := map[int]float64{0: 0, 1: 3.38, 2: 0.5}
			So(solutions, ShouldResemble, out)
		})
	})

}
