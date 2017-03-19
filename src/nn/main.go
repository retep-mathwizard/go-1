package main

import (
	"fmt"
	"math/rand"
	"time"
)

//great tutorial at https://stevenmiller888.github.io/mind-how-to-build-a-neural-network/
func randomWeights(width, depth int) [][]float64 {
	rand.Seed(time.Now().UnixNano()) //truly random
	output := blank(width, depth)
	for x := 0; x < depth; x++ {
		for y := 0; y < width; y++ {
			output[x][y] = rand.Float64()
		}
	}

	return output
}
func blank(width, height int) [][]float64 {
	output := make([][]float64, height)
	for x := 0; x < height; x++ {
		output[x] = make([]float64, width)
	}
	return output
}

func dot(Matrix1, Matrix2 [][]float64) [][]float64 {
	//Matrix1 is the weights, multiply by inputs to get output
	var output := blank(len(Matrix1[0]), len(Matrix1))

}
func main() {
	//inputs := [][]float64{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	startingWeights := randomWeights(3, 1)
	fmt.Println(startingWeights)

}
