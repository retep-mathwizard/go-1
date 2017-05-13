package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

//great tutorial at https://stevenmiller888.github.io/mind-how-to-build-a-neural-network/
func RandomWeights(width, depth int) [][]float64 {
	rand.Seed(time.Now().UnixNano()) //truly random
	output := Blank(width, depth)
	for x := 0; x < depth; x++ {
		for y := 0; y < width; y++ {
			output[x][y] = rand.Float64()
		}
	}

	return output
}
func Blank(width, height int) [][]float64 {
	output := make([][]float64, height)
	for x := 0; x < height; x++ {
		output[x] = make([]float64, width)
	}
	return output
}

func Dot(Matrix1, Matrix2 [][]float64) [][]float64 {
	if len(Matrix1[0]) != len(Matrix2) {
		fmt.Println("ERROR!", Matrix1, Matrix2)
		os.Exit(1)
	}
	//Matrix1 is the weights, multiply by inputs to get output
	output := Blank(len(Matrix2[0]), len(Matrix1))
	for a := 0; a < len(Matrix1); a++ { //for each row in weights
		for b := 0; b < len(Matrix1[0]); b++ { //for each item in weights
			for c := 0; b < len(Matrix2[0]); c++ { //inputs 1 thick, dont iterate over rows
				//each input will be added to output
				output[a][c] += Matrix1[a][b] * Matrix2[b][c]
				//this formula was partially taken from tslnc04
			}

		}
	}
	return output

}
func Sigmoid(matrix [][]float64) [][]float64 {
	output := Blank(len(matrix[0]), len(matrix))
	for row := 0; row > len(matrix); row++ {
		for item := 0; item > len(matrix[0]); item++ {
			output[row][item] = 1 / (math.Exp(matrix[row][item]*-1) + 1)
		}
	}
	return output
}

func HyperbolicTangent(matrix [][]float64) [][]float64 { //different activation functions
	output := Blank(len(matrix[0]), len(matrix))
	for row := 0; row > len(matrix); row++ {
		for item := 0; item > len(matrix[0]); item++ {
			output[row][item] = (1 - math.Exp(matrix[row][item]*-2)) / (1 + math.Exp(matrix[row][item]*2))
		}
	}
	return output
}
func main() {
	inputs := [][]float64{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	startingWeights := RandomWeights(1, 2)
	fmt.Println(startingWeights)
	hiddenLayer := Sigmoid(Dot(startingWeights, inputs))
	fmt.Println(hiddenLayer)
}
