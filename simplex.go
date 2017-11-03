package main

import (
	"fmt"
	"math"
)

func simplex(A map[int]map[int]float64, rows int, columns int) map[int]map[int]float64 {

	for {
		/////////////////////empieza el pivoteo/////////////////////////
		var colpivot int
		for i := 0; i < columns; i++ {
			if A[0][i] < 0 && A[0][i] < A[0][i+1] {
				colpivot = i
			}
		}
		/////////////////fila pivote////////////////////////////////
		w := make([]float64, rows-1)

		var rowspivot int
		for k := 1; k < rows; k++ {
			w[k-1] = float64(A[k][columns-1]) / float64(A[k][colpivot])
			if w[k-1] < 0 {
				w[k-1] = math.MaxFloat64
			}
			if A[k][colpivot] == 0 {
				w[k-1] = math.MaxFloat64
			}
		}

		minimum := math.MaxFloat64
		for i := 0; i < len(w); i++ {
			if w[i] < minimum {
				minimum = w[i]
				rowspivot = i + 1
			}
		}
		fmt.Println("minimo", minimum)
		fmt.Println("filapivote", rowspivot)

		elementpivot := A[rowspivot][colpivot]

		////////////////////normalizar y reduccion/////////////

		for j := 0; j < columns; j++ {
			A[rowspivot][j] = float64(A[rowspivot][j]) / float64(elementpivot)
		}
		////////////nueva matriz luego de las operaciones con el pivote/////////
		Aux := make(map[int]map[int]float64)

		for i := 0; i < rows; i++ {
			Aux[i] = map[int]float64{}
		}

		for i := 0; i < rows; i++ {
			for j := 0; j < columns; j++ {

				if i != rowspivot && A[i][colpivot] != 0 {
					Aux[i][j] = (-A[i][colpivot])*(A[rowspivot][j]) + A[i][j]
				} else {
					Aux[i][j] = A[i][j]
				}
			}
		}
		A = Aux
		////////////condicion de parada//////////////////
		number := 0
		for i := 1; i <= columns-1; i++ {
			if A[0][i] >= 0 {
				number = number + 1
			}
		}

		if number == columns-1 {
			break
		}
	}
	return A
}

func main() {
	// ejemplo que ya funciona
	// a := []float64{3, 2, 5}
	// b := []float64{430, 460, 420}
	// A1 := []float64{1, 2, 1, 3, 0, 2, 1, 4, 0} //colocar los valores de las restricciones de derecha a izquierda incluyendo los ceros
	//x1=0 x2=100 x3=230 Z=1350

	// ejemplo que ya funciona
	// a := []float64{3, 2}
	// b := []float64{18, 42, 24}
	// A1 := []float64{2, 1, 2, 3, 3, 1}
	//x1=3 x2=12 z=33

	// ejemplo que ya funciona
	a := []float64{40, 60}
	b := []float64{70, 40, 90}
	A1 := []float64{2, 1, 1, 1, 1, 3}
	//x1=15 x2=25 z=2100

	// ejemplo que ya funciona
	// a := []float64{2, 2, -3}
	// b := []float64{4, 2, 12}
	// A1 := []float64{-1, 1, 1, 2, -1, 1, 1, 1, 3} //colocar los valores de las restricciones de derecha a izquierda incluyendo los ceros
	//x1=4 x2=8 x3=0 z=24

	// ejemplo que ya funciona
	// a := []float64{2, 3, 10}
	// b := []float64{425, 25, 120}
	// A1 := []float64{1, 1, 1, 2, -4, 1, -1, -1, -3} //colocar los valores de las restricciones de derecha a izquierda incluyendo los ceros
	//x1=0 x2=80 x3=345 z=3690

	// ejemplo que ya funciona
	// a := []float64{2, 3, 10, 5, 3}
	// b := []float64{425, 25, 120, 6, 1}
	// A1 := []float64{1, 1, 1, 3, 4, 2, -4, 1, 0, 0, -1, -1, -3, 0, 0, 1, 11, 5, 0, 0, 1, 1, 7, 0, 0}
	// x1=0.00000000   x2=0.51388889   x3=0.06944444 x4=141.47222222   x5=0.00000000 z=709.5972

	fmt.Println("coeficientes de la funcion maxima:", a)
	fmt.Println("vector b:", b)

	C := make(map[int]map[int]float64)

	for i := 0; i < len(b); i++ {
		C[i] = map[int]float64{}
	}

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(a); j++ {
			C[i][j] = A1[i*(len(a))+j]
		}
	}

	fmt.Println("sub matriz de las restricciones C", C)

	A := make(map[int]map[int]float64)
	var rows, columns int
	rows = len(b) + 1
	columns = len(a) + len(b) + 2
	for i := 0; i < rows; i++ {
		A[i] = map[int]float64{}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			A[i][j] = 0
		}
	}

	A[0][0] = 1.00
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 && j >= 1 && j <= len(a) {
				A[i][j] = -a[j-1]
			} else if i > 0 && j >= 1 && j <= len(a) {
				A[i][j] = C[i-1][j-1]
			} else if j > len(a) && j < columns-1 {
				A[j-len(a)][j] = 1.00
			} else if i > 0 && j == columns-1 {
				A[i][j] = b[i-1]
			}
		}
	}
	A = simplex(A, rows, columns)
	/////////////////valores maximos del simplex///////////

	result := make(map[int]float64)

	result[0] = A[0][columns-1]
	fmt.Println("maximum of the objective function", result[0])

	for i := 1; i <= len(a); i++ {
		for j := 0; j <= len(b); j++ {
			if A[0][i] > 0 {
				result[i] = 0
			} else if A[0][i] == 0 && A[j][i] == 1 {
				result[i] = A[j][columns-1]
			}
		}
	}
	fmt.Println("resultado", result)
}
