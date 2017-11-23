package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	homeSlice := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		smallSlice := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			newInt := uint8(j * 550)
			smallSlice[j] = newInt
		}
		homeSlice[i] = smallSlice
	}
	return homeSlice
}

func main() {
	pic.Show(Pic)
}
