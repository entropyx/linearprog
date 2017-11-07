package linearprog

import (
	"fmt"
	"math"
)

func Round(num float64, precision int) float64 {
	var out float64
	v1 := math.Pow(10, float64(precision))
	v2 := int(num*v1 + math.Copysign(0.5, num*v1))
	out = float64(v2) / v1
	return out
}

func Matrix(A map[int]map[int]float64) {
	for i := 0; i < len(A); i++ {
		var row []float64
		for j := 0; j < len(A[i]); j++ {
			row = append(row, A[i][j])
		}
		fmt.Println(row)
	}
}

func Simplex(A map[int]map[int]float64, b []float64, a []float64, constdir []string) (map[int]float64, float64) {
	var feasible bool
	Pivot := make(map[int]map[int]float64)
	solutions := make(map[int]float64)
	var rows, columns int
	rows = len(b) + 1
	columns = len(a) + len(b) + 2

	for i := 0; i < rows; i++ {
		Pivot[i] = map[int]float64{}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			Pivot[i][j] = 0
		}
	}

	Pivot[0][0] = 1.00
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if i == 0 && j >= 1 && j <= len(a) {
				Pivot[i][j] = -a[j-1]
			} else if i > 0 && j >= 1 && j <= len(a) {
				Pivot[i][j] = A[i][j]
			} else if i > 0 && j > len(a) && j < columns-1 && constdir[i-1] == "<=" {
				Pivot[i][i+len(a)] = 1.00
			} else if i > 0 && j > len(a) && j < columns-1 && constdir[i-1] == ">=" {
				Pivot[i][i+len(a)] = -1.00
			} else if i > 0 && j == columns-1 {
				Pivot[i][j] = b[i-1]
			}

		}
	}

	for {

		colpivot := 1
		min := Pivot[0][1]
		for i := 1; i < columns; i++ {
			if feasible {
				if Pivot[0][i] < 0 && min > Pivot[0][i] {
					colpivot = i
					min = Pivot[0][i]
				}
			} else {
				if Pivot[0][i] < 0 && min > Pivot[0][i] && i <= len(a) {
					colpivot = i
					min = Pivot[0][i]
				}
			}

		}
		min = math.Inf(1)
		rowspivot := 1
		for k := 1; k < rows; k++ {
			v := float64(Pivot[k][columns-1]) / float64(Pivot[k][colpivot])
			if v > 0 {
				if min > v {
					min = v
					rowspivot = k
				}
			}
		}
		elementpivot := Pivot[rowspivot][colpivot]
		for j := 0; j < columns; j++ {
			Pivot[rowspivot][j] = float64(Pivot[rowspivot][j]) / float64(elementpivot)
		}
		for i := 0; i < rows; i++ {
			lowervalue := Pivot[i][colpivot]
			for j := 0; j < columns; j++ {
				if i != rowspivot {
					Pivot[i][j] = (-lowervalue)*(Pivot[rowspivot][j]) + Pivot[i][j]
				}
			}
		}
		number := 0
		numberfeasible := 0
		for i := 0; i < columns; i++ {
			if Pivot[0][i] >= 0 {
				number = number + 1
			}
			if Pivot[0][i] >= 0 && i < len(a) {
				numberfeasible = numberfeasible + 1
			}
		}
		if numberfeasible == len(a) {
			feasible = true
		}
		if number == columns {
			break
		}
	}

	opt := Pivot[0][columns-1]

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if Pivot[0][i] > 0 {
				solutions[i-1] = 0
			} else if Pivot[0][i] == 0 && Pivot[j][i] == 1 {
				solutions[i-1] = Round(Pivot[j][columns-1], 2)
			}
		}
	}

	return solutions, opt
}
