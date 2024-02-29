package main

import "fmt"

type image [][]int32

func rotateImage(img image) image {
	// image is nxn, length of img will be the same for each sub-element in img
	imgLen := len(img)
	imgTmp := make(image, imgLen)
	for i := 0; i < imgLen; i++ {
		imgTmp[i] = make([]int32, imgLen)
	}

	col := 0
	row := imgLen - 1

	for i := 0; i < imgLen; i++ {
		for j := 0; j < imgLen; j++ {
			imgTmp[i][j] = img[row][col]
			row--
		}
		col++
		row = imgLen - 1
	}

	return imgTmp
}

func main() {
	img := image{
		{1, 0, 0, 1},
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
	}

	fmt.Printf("%v\n", img)
	rotated := rotateImage(img)
	fmt.Printf("%v\n", rotated)
}
